# Pion WebRTC Live Instructor Video — Implementation Spec

## Context

This is a Go-based instructor-led training application built on **Gotea**, a custom Go TEA (The Elm Architecture) framework. Gotea provides server-side rendered SPAs with WebSocket communication and DOM patching via **morphdom**. All application state lives on the server; the client is thin.

The app already has:
- A working WebSocket connection per client (managed by Gotea)
- Training sessions where an instructor navigates content and students follow along in sync
- A session model that tracks which users are instructors vs students

We want to embed live video from the instructor directly into the training UI, so students see the instructor's camera feed without needing a separate Zoom/Teams call.

## Approach

Use **Pion WebRTC** (`github.com/pion/webrtc/v4`) as an embedded SFU (Selective Forwarding Unit) within the Go binary. The instructor's browser captures their camera/mic via `getUserMedia`, sends it to the server over WebRTC, and the server fans the media out to all student peer connections. Signaling (SDP offer/answer and ICE candidate exchange) happens over the existing Gotea WebSocket — no additional signaling server is needed.

This is a **one-to-many broadcast** pattern: one publisher (instructor), many receive-only viewers (students).

## Architecture Overview

```
┌─────────────────────────────────────────────────────┐
│                   Go Binary                         │
│                                                     │
│  ┌──────────────────┐    ┌───────────────────────┐  │
│  │     Gotea         │    │   Broadcast Manager   │  │
│  │  (training app    │    │                       │  │
│  │   logic, content  │    │  - Instructor PC      │  │
│  │   navigation,     │◄──►│  - Video/Audio tracks │  │
│  │   session mgmt)   │    │  - Viewer PCs[]       │  │
│  │                   │    │  - ICE handling        │  │
│  └──────┬───────────┘    └──────────┬────────────┘  │
│         │                           │                │
│         │  Shared WebSocket conn    │                │
│         └──────────┬────────────────┘                │
└────────────────────┼─────────────────────────────────┘
                     │
            ┌────────┴────────┐
            │   Browser(s)    │
            │                 │
            │  Gotea client   │  ← DOM patching as normal
            │  + webrtc.js    │  ← getUserMedia (instructor)
            │                 │  ← <video> playback (all)
            └─────────────────┘
```

## Dependencies

```
go get github.com/pion/webrtc/v4
```

No other new dependencies should be needed. The existing Gotea WebSocket infrastructure handles signaling.

## Implementation Plan

### 1. WebSocket Message Multiplexing

The existing Gotea WebSocket carries messages between client and server. We need to add WebRTC signaling messages alongside the existing Gotea message types.

**New message types to add (JSON over the existing WebSocket):**

```json
// Client → Server
{ "type": "webrtc:offer", "sdp": "..." }
{ "type": "webrtc:ice-candidate", "candidate": { ... } }
{ "type": "webrtc:stop-broadcast" }

// Server → Client
{ "type": "webrtc:answer", "sdp": "..." }
{ "type": "webrtc:ice-candidate", "candidate": { ... } }
{ "type": "webrtc:broadcast-started" }
{ "type": "webrtc:broadcast-stopped" }
{ "type": "webrtc:viewer-offer", "sdp": "..." }
```

**Implementation approach:** In the WebSocket message handler, add a router/switch that checks for the `webrtc:` prefix. Messages with this prefix are dispatched to the broadcast manager; all others continue to Gotea as normal. This should be a minimal, non-invasive change to the existing WebSocket handling code.

Look at the existing WebSocket message handling code to understand the current message format and routing, then add the `webrtc:` message type alongside it using the same patterns.

### 2. Broadcast Manager (Server-Side Go)

Create a new package/file for managing WebRTC broadcast state. This should be a struct that can be associated with a training session.

```go
// broadcast/manager.go

package broadcast

import (
    "sync"
    "github.com/pion/webrtc/v4"
)

type Manager struct {
    mu sync.RWMutex

    // WebRTC configuration (STUN servers etc)
    config webrtc.Configuration

    // Instructor's peer connection
    instructorPC *webrtc.PeerConnection

    // Local tracks that receive the instructor's media and fan out to viewers
    videoTrack *webrtc.TrackLocalStaticRTP
    audioTrack *webrtc.TrackLocalStaticRTP

    // Viewer peer connections, keyed by some client/student identifier
    viewers map[string]*webrtc.PeerConnection

    // Whether a broadcast is currently active
    active bool

    // Callback for sending signaling messages back to clients via the
    // existing WebSocket. The broadcast manager should NOT import or
    // depend on Gotea directly — use a callback/interface instead.
    sendMessage func(clientID string, msg interface{})
}
```

**Key design principle:** The broadcast manager should be decoupled from Gotea. It accepts a `sendMessage` callback (or interface) that Gotea provides, which it uses to send signaling messages (answers, ICE candidates) back to specific clients via their WebSocket connections. This keeps the broadcast package reusable and testable.

#### WebRTC Configuration

```go
func NewManager(sendMessage func(clientID string, msg interface{})) *Manager {
    return &Manager{
        config: webrtc.Configuration{
            ICEServers: []webrtc.ICEServer{
                {URLs: []string{"stun:stun.l.google.com:19302"}},
            },
        },
        viewers:     make(map[string]*webrtc.PeerConnection),
        sendMessage: sendMessage,
    }
}
```

For initial development, a public STUN server is fine. TURN is not needed unless users are behind strict symmetric NATs or corporate firewalls — note this as a future consideration but don't implement it now.

#### Handling the Instructor's Offer

When the instructor clicks "Start Broadcast" and the browser sends a `webrtc:offer`:

1. Create a new `PeerConnection` for the instructor
2. Set the remote description (the offer)
3. Register an `OnTrack` handler that:
   - Creates a `TrackLocalStaticRTP` for each incoming track (video and audio)
   - Stores them on the manager
   - Spawns a goroutine that reads RTP packets from the remote track and writes them to the local track (this is the fan-out mechanism)
4. Register an `OnICECandidate` handler that sends candidates back to the instructor's browser via the WebSocket
5. Create an answer and set it as the local description
6. Send the answer back to the instructor's browser via the WebSocket
7. Mark the broadcast as active and notify all connected students that a broadcast has started (so their UI can update)

#### Handling a Viewer's Connection

When a student's browser receives a `webrtc:broadcast-started` message, it should automatically initiate a viewer connection:

1. Create a new `PeerConnection` for this viewer
2. Add the existing `videoTrack` and `audioTrack` to the peer connection (these are the `TrackLocalStaticRTP` tracks being fed by the instructor)
3. Register `OnICECandidate` to send candidates back to this viewer
4. Create an offer from the server side, set it as local description
5. Send the offer to the viewer's browser
6. When the viewer's browser sends back an answer, set it as the remote description

**Important:** The server initiates the offer to viewers (not the other way around), because the server is the one adding tracks. This simplifies the client-side viewer code.

#### Late Joiners

If a student joins a session where the broadcast is already active, the manager should immediately set up a viewer peer connection for them using the same flow as above.

#### Stopping the Broadcast

When the instructor stops broadcasting:
1. Close the instructor's peer connection
2. Close all viewer peer connections
3. Clear the tracks
4. Notify all students via WebSocket so their UI updates
5. Clean up all state

#### Cleanup

When a student disconnects (WebSocket closes), close and remove their viewer peer connection. When the instructor disconnects, stop the broadcast entirely.

### 3. Integration with Gotea Session

The `Manager` should be instantiated and stored as part of the training session state. When the session is created, create a broadcast manager. When the session is destroyed, clean up the broadcast manager.

Wire the `sendMessage` callback to use Gotea's existing mechanism for sending WebSocket messages to specific clients.

In the Gotea view/render function for the training session, conditionally render the video container when a broadcast is active (or always render it as hidden and show/hide it).

### 4. Morphdom Protection

Gotea uses morphdom for DOM patching. The video container must be protected from morphdom updates since its contents are managed by client-side JavaScript.

**Add an `onBeforeElUpdated` callback** to the morphdom call (wherever Gotea invokes morphdom) that skips elements with a `data-gotea-ignore` attribute:

```javascript
onBeforeElUpdated: function(fromEl, toEl) {
    if (fromEl.hasAttribute('data-gotea-ignore')) return false;
    return true;
}
```

If Gotea already has a morphdom options configuration, add this there. If not, this is where to introduce one.

The video container in the rendered HTML should then use this attribute:

```html
<div id="instructor-video-container" data-gotea-ignore>
    <!-- Client JS owns this subtree -->
    <video id="instructor-video" autoplay playsinline muted></video>
</div>
```

**Note:** The `muted` attribute is only for the instructor's local preview (to avoid echo). Viewer video elements should NOT be muted.

### 5. Client-Side JavaScript

Create a small JavaScript module (`webrtc.js` or similar) that handles the browser-side WebRTC logic. This should be loaded as a standalone script, not managed by Gotea's rendering.

#### Message Routing

Hook into the existing WebSocket `onmessage` handler. Before the message reaches Gotea's handler, check if it's a `webrtc:` message and route it to the WebRTC module instead. Something like:

```javascript
// Wrap or extend the existing WebSocket onmessage
const originalOnMessage = ws.onmessage;
ws.onmessage = function(event) {
    const data = JSON.parse(event.data);
    if (data.type && data.type.startsWith('webrtc:')) {
        handleWebRTCMessage(data);
    } else {
        originalOnMessage.call(ws, event);
    }
};
```

Adapt this pattern to however Gotea currently sets up its WebSocket listener.

#### Instructor Flow

```javascript
// When instructor clicks "Start Broadcast"
async function startBroadcast() {
    const stream = await navigator.mediaDevices.getUserMedia({
        video: {
            width: { ideal: 640 },
            height: { ideal: 480 },
            frameRate: { ideal: 24 }
        },
        audio: true
    });

    const pc = new RTCPeerConnection({
        iceServers: [{ urls: 'stun:stun.l.google.com:19302' }]
    });

    // Add tracks from camera/mic to the peer connection
    stream.getTracks().forEach(track => pc.addTrack(track, stream));

    // Send ICE candidates to server as they are gathered
    pc.onicecandidate = (event) => {
        if (event.candidate) {
            sendWebRTCMessage({
                type: 'webrtc:ice-candidate',
                candidate: event.candidate
            });
        }
    };

    // Create and send offer
    const offer = await pc.createOffer();
    await pc.setLocalDescription(offer);
    sendWebRTCMessage({
        type: 'webrtc:offer',
        sdp: offer.sdp
    });

    // Show local preview (muted to avoid echo)
    const videoEl = document.getElementById('instructor-video');
    videoEl.srcObject = stream;
    videoEl.muted = true;

    // Store for later cleanup
    window._broadcastPC = pc;
    window._broadcastStream = stream;
}

// When server sends back an answer
function handleAnswer(sdp) {
    const answer = new RTCSessionDescription({
        type: 'answer',
        sdp: sdp
    });
    window._broadcastPC.setRemoteDescription(answer);
}

// When instructor clicks "Stop Broadcast"
function stopBroadcast() {
    if (window._broadcastStream) {
        window._broadcastStream.getTracks().forEach(t => t.stop());
    }
    if (window._broadcastPC) {
        window._broadcastPC.close();
    }
    sendWebRTCMessage({ type: 'webrtc:stop-broadcast' });
}
```

#### Viewer Flow

```javascript
// When server sends a viewer offer (broadcast started or late join)
async function handleViewerOffer(sdp) {
    const pc = new RTCPeerConnection({
        iceServers: [{ urls: 'stun:stun.l.google.com:19302' }]
    });

    // When we receive the instructor's tracks, display them
    pc.ontrack = (event) => {
        const videoEl = document.getElementById('instructor-video');
        // ontrack may fire multiple times (video + audio).
        // Set srcObject once, the browser handles combining tracks.
        if (!videoEl.srcObject) {
            videoEl.srcObject = new MediaStream();
        }
        videoEl.srcObject.addTrack(event.track);
        videoEl.play().catch(() => {
            // Autoplay may be blocked — show a "click to unmute" button
            showUnmutePrompt();
        });
    };

    pc.onicecandidate = (event) => {
        if (event.candidate) {
            sendWebRTCMessage({
                type: 'webrtc:ice-candidate',
                candidate: event.candidate
            });
        }
    };

    await pc.setRemoteDescription(new RTCSessionDescription({
        type: 'offer',
        sdp: sdp
    }));

    const answer = await pc.createAnswer();
    await pc.setLocalDescription(answer);

    sendWebRTCMessage({
        type: 'webrtc:answer',
        sdp: answer.sdp
    });

    window._viewerPC = pc;
}

// When broadcast stops
function handleBroadcastStopped() {
    if (window._viewerPC) {
        window._viewerPC.close();
        window._viewerPC = null;
    }
    const videoEl = document.getElementById('instructor-video');
    videoEl.srcObject = null;
    // Hide the video container or show a "broadcast ended" message
}
```

#### Autoplay Policy Handling

Browsers block autoplay of unmuted video. Viewers need to handle this:

```javascript
function showUnmutePrompt() {
    // Show a button/overlay on the video container that says
    // "Click to enable audio" — on click, call videoEl.play()
    // This satisfies the browser's user-gesture requirement
}
```

### 6. UI Components

#### Instructor Controls

Render a "Start Broadcast" / "Stop Broadcast" button in the instructor's training session UI. This should be part of Gotea's normal rendered HTML (an instructor toolbar or similar). The button sends a Gotea event, which triggers `startBroadcast()` on the client side via a small JS bridge.

Alternatively, the button can directly call the JS function if the instructor toolbar has client-side event handling.

#### Video Display

Render the video container as part of the training session view. Suggested placement: a picture-in-picture style overlay in the corner, or a sidebar panel. The container should be shown/hidden based on broadcast state.

```html
<!-- Rendered by Gotea as part of the session view -->
<div id="video-panel" class="video-panel" data-gotea-ignore>
    <video id="instructor-video" autoplay playsinline></video>
    <div id="video-controls">
        <!-- For instructor: stop button, camera/mic toggles -->
        <!-- For viewer: unmute button if needed -->
    </div>
</div>
```

Style the video panel with CSS to be a floating overlay or sidebar. Start with a simple fixed-position overlay in the bottom-right corner:

```css
.video-panel {
    position: fixed;
    bottom: 20px;
    right: 20px;
    width: 320px;
    z-index: 1000;
    border-radius: 8px;
    overflow: hidden;
    box-shadow: 0 4px 12px rgba(0,0,0,0.3);
    background: #000;
}

.video-panel video {
    width: 100%;
    display: block;
}

.video-panel.hidden {
    display: none;
}
```

### 7. Server-Side Viewer Offer Flow (Important Detail)

The flow where the **server** creates and sends the offer to viewers (rather than viewers sending offers to the server) requires careful implementation:

1. When a broadcast starts (or a late joiner connects), the server creates a PeerConnection for the viewer
2. The server adds the instructor's local tracks to this PC
3. The server creates an offer and sets it as local description
4. The server sends this offer to the viewer's browser via WebSocket
5. The viewer's browser creates an answer and sends it back
6. The server sets the answer as remote description

This is the correct flow because the server is the side adding tracks. If we did it the other way (viewer sends offer), the viewer wouldn't know to include the right transceiver directions and things get more complex.

### 8. ICE Candidate Handling

ICE candidates can arrive before the remote description is set (a race condition). Buffer them:

**Client side:** If an ICE candidate message arrives before `setRemoteDescription` has been called, queue it and apply after the remote description is set.

**Server side:** Same pattern — if ICE candidates arrive from the client before the PeerConnection's remote description is set, buffer them and apply after.

This is a common WebRTC gotcha. Pion may handle some of this internally, but be defensive.

### 9. RTP Forwarding Goroutine

The core of the SFU is the goroutine that reads RTP packets from the instructor's remote track and writes them to the local track:

```go
pc.OnTrack(func(remoteTrack *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
    // Create a local track with the same codec
    localTrack, err := webrtc.NewTrackLocalStaticRTP(
        remoteTrack.Codec().RTPCodecCapability,
        remoteTrack.Kind().String(),
        "instructor",
    )
    if err != nil {
        // handle error
        return
    }

    // Store the local track on the manager
    m.mu.Lock()
    if remoteTrack.Kind() == webrtc.RTPCodecTypeVideo {
        m.videoTrack = localTrack
    } else {
        m.audioTrack = localTrack
    }
    m.mu.Unlock()

    // Forward RTP packets
    buf := make([]byte, 1500)
    for {
        n, _, readErr := remoteTrack.Read(buf)
        if readErr != nil {
            return // Track ended or connection closed
        }
        if _, writeErr := localTrack.Write(buf[:n]); writeErr != nil {
            return
        }
    }
})
```

The `TrackLocalStaticRTP` automatically fans out written packets to all PeerConnections that have added this track. This is Pion's built-in SFU mechanism — you don't need to manually iterate over viewers.

## File Structure

Suggested new files (adapt to existing project structure):

```
broadcast/
    manager.go          # Broadcast manager struct and methods
    signaling.go        # Message types for WebRTC signaling

static/js/
    webrtc.js           # Client-side WebRTC module

static/css/
    video-panel.css     # Video overlay styles (or add to existing CSS)
```

Plus modifications to:
- WebSocket message handler (add webrtc: routing)
- Training session model (add broadcast manager)
- Training session view (add video container HTML)
- Morphdom configuration (add data-gotea-ignore support)
- Client-side Gotea JS (add webrtc: message routing)

## Testing

1. **Local testing:** Open two browser tabs, one as instructor, one as student. The instructor starts a broadcast; the student should see the video. Both are on localhost so NAT traversal isn't an issue.
2. **Same network:** Test with two different machines on the same LAN. Should work with just STUN.
3. **Different networks:** Test across different networks. If it fails, a TURN server is needed (future enhancement).

## Future Enhancements (Not for initial implementation)

- **TURN server support** for restrictive network environments (add Coturn config)
- **Mic mute/camera off toggles** for the instructor
- **Screen sharing** as an alternative to camera (use `getDisplayMedia` instead of `getUserMedia`)
- **Audio-only mode** for low bandwidth situations
- **Connection quality indicators** using Pion's stats
- **Recording** — Pion can write media to disk for session replay
- **Simulcast** — instructor sends multiple quality levels, server selects based on viewer bandwidth
- **Draggable/resizable** video overlay

## Key Gotchas

1. **Autoplay policy**: Browsers won't autoplay unmuted video without a user gesture. Handle this gracefully with an unmute prompt.
2. **ICE candidate buffering**: Candidates can arrive before remote description is set. Buffer them.
3. **Goroutine cleanup**: The RTP forwarding goroutines must exit cleanly when the broadcast stops or a viewer disconnects. Make sure PeerConnection.Close() terminates the Read() calls.
4. **Thread safety**: The broadcast manager will be accessed from multiple goroutines (one per WebSocket connection). Use the mutex.
5. **Codec negotiation**: Let Pion handle codec negotiation. Don't hardcode VP8 or H264 — use whatever the instructor's browser offers. The `TrackLocalStaticRTP` with `RTPCodecCapability` from the remote track handles this.
6. **HTTPS requirement**: `getUserMedia` only works on HTTPS (or localhost). The training app must be served over HTTPS in production.
