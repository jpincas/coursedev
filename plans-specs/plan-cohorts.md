# Cohorts, Authentication & Persistence Plan

## Overview

Transform the training app from session-based anonymous access to cohort-based authenticated access with persistent progress.

**Core identity model**: A student is uniquely identified by `(cohort_id, email)`. The same email in different cohorts = different students with separate progress. Browser sessions map to students, allowing progress restoration after cookie expiry.

---

## Database Schema

### New Tables

```sql
-- Cohorts define groups of students
CREATE TABLE cohorts (
    id TEXT PRIMARY KEY,              -- UUID
    code TEXT UNIQUE NOT NULL,        -- e.g., "abc-defg-hij"
    name TEXT NOT NULL,               -- Display name
    expires_at DATETIME NOT NULL,
    max_students INTEGER NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Students are uniquely identified by cohort + email
CREATE TABLE students (
    id TEXT PRIMARY KEY,              -- UUID
    cohort_id TEXT NOT NULL,
    email TEXT NOT NULL,
    name TEXT NOT NULL,
    current_module TEXT NOT NULL DEFAULT '',
    current_page INTEGER NOT NULL DEFAULT 0,
    font_size TEXT NOT NULL DEFAULT 'medium',
    theme TEXT NOT NULL DEFAULT 'light',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(cohort_id, email),
    FOREIGN KEY (cohort_id) REFERENCES cohorts(id) ON DELETE CASCADE
);

-- Maps browser sessions to students
CREATE TABLE sessions (
    session_id TEXT PRIMARY KEY,      -- Gotea session UUID
    student_id TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
);

-- Indexes
CREATE INDEX idx_cohorts_code ON cohorts(code);
CREATE INDEX idx_students_cohort ON students(cohort_id);
CREATE INDEX idx_sessions_student ON sessions(student_id);
```

### Modified Tables

Rename `learner_id` → `student_id` in existing progress tables:
- `module_progress`
- `page_views`
- `quiz_scores`

Drop the old `learners` table (replaced by `students`).

---

## Authentication Flow

### Entry Points

1. **Base URL `/`** - Shows cohort code entry if not authenticated
2. **Join link `/join/:code`** - Pre-fills the cohort code

### Flow States

```
┌─────────────────┐
│   New Session   │
│  (no mapping)   │
└────────┬────────┘
         │
         ▼
┌─────────────────┐     invalid/expired
│ Enter Cohort    │─────────────────────► Error screen
│    Code         │
└────────┬────────┘
         │ valid
         ▼
┌─────────────────┐     cohort full
│ Enter Email     │─────────────────────► "Contact training manager"
│   (+ Name)      │
└────────┬────────┘
         │ success
         ▼
┌─────────────────┐
│ Course Content  │◄──── Returning session (has mapping)
└─────────────────┘
```

### Returning User

1. Session exists → Look up session → student mapping
2. Load student record (includes cohort_id)
3. Validate cohort is not expired
4. If expired: clear session, show "cohort expired" error
5. If valid: load progress, show content

### New User in Existing Cohort

1. Enter cohort code → validated
2. Enter email → found existing student for (cohort, email)
3. Create session → student mapping
4. Load their existing progress

### Brand New User

1. Enter cohort code → validated, check max_students
2. Enter email + name → no existing student
3. Create student record
4. Create session → student mapping
5. Initialize empty progress

---

## Course Owner (Admin)

### Authentication

- Environment variable: `COURSEDEV_ADMIN_KEY`
- Format: Same as cohort codes (e.g., `admin-master-key`)
- Entered in the same cohort code field

### Owner Session

When master key is entered:
- No email/name required (or just name for display)
- `IsOwner = true` flag set on Model
- No cohort/student record created (transient session)
- Progress not persisted (owner isn't taking the course)

### Owner Capabilities

- See "Admin" link in sidebar
- Access `/admin` route (cohort management)
- Can start presenting (students cannot)
- Full navigation without progress tracking

---

## Model Changes

```go
type Model struct {
    gt.Router

    // Shared (read-only at runtime)
    Course  *CourseGraph
    Modules map[string]*Module

    // Session identity
    SessionID uuid.UUID       // Gotea session ID (always set)
    StudentID *uuid.UUID      // Set after login (nil for owner or unauthenticated)
    CohortID  *uuid.UUID      // Set after login (nil for owner)
    IsOwner   bool            // True if logged in via admin key

    // Authentication flow state
    AuthStage      AuthStage  // Current step in auth flow
    PendingCohort  *Cohort    // Validated cohort during auth
    AuthError      string     // Error message to display

    // Learner state (loaded from DB after auth)
    CurrentModule string
    CurrentPage   int
    Progress      map[string]*ModuleProgress
    ActiveQuiz    *QuizState
    ActiveHotspot string
    Preferences   LearnerPreferences

    // Live session mode (existing)
    PresentingSessionID *uuid.UUID
    FollowingSessionID  *uuid.UUID
}

type AuthStage string
const (
    AuthStageNone       AuthStage = ""           // Check session
    AuthStageEnterCode  AuthStage = "enter_code" // Waiting for cohort code
    AuthStageEnterEmail AuthStage = "enter_email"// Waiting for email/name
    AuthStageError      AuthStage = "error"      // Show error, retry
)
```

---

## New Messages

### Authentication

| Message | Args | Handler |
|---------|------|---------|
| `SUBMIT_COHORT_CODE` | `{code: string}` | Validate code, advance to email or set error |
| `SUBMIT_LOGIN` | `{email: string, name: string}` | Create/find student, create session mapping |
| `LOGOUT` | none | Clear session mapping, return to code entry |

### Admin (Cohort Management)

| Message | Args | Handler |
|---------|------|---------|
| `CREATE_COHORT` | `{name, expiresAt, maxStudents}` | Generate code, insert cohort |
| `UPDATE_COHORT` | `{id, name?, expiresAt?, maxStudents?}` | Update cohort |
| `DELETE_COHORT` | `{id}` | Delete cohort (cascades to students) |

---

## Routes

| Route | Description |
|-------|-------------|
| `/` | Main content or auth flow |
| `/join/:code` | Pre-fill cohort code, then auth flow |
| `/admin` | Cohort management (owner only) |
| `/admin/cohort/:id` | Edit specific cohort (owner only) |

---

## UI Components

### 1. Cohort Code Entry Screen

```
┌─────────────────────────────────────────┐
│                                         │
│         [Logo/Course Title]             │
│                                         │
│    Enter your cohort code to begin      │
│                                         │
│    ┌─────────────────────────────┐      │
│    │  xxx-xxxx-xxx               │      │
│    └─────────────────────────────┘      │
│                                         │
│           [ Continue ]                  │
│                                         │
└─────────────────────────────────────────┘
```

### 2. Email/Name Entry Screen

```
┌─────────────────────────────────────────┐
│                                         │
│         Joining: [Cohort Name]          │
│                                         │
│    Email                                │
│    ┌─────────────────────────────┐      │
│    │  alice@example.com          │      │
│    └─────────────────────────────┘      │
│                                         │
│    Your name                            │
│    ┌─────────────────────────────┐      │
│    │  Alice Smith                │      │
│    └─────────────────────────────┘      │
│                                         │
│           [ Join Course ]               │
│                                         │
└─────────────────────────────────────────┘
```

### 3. Error States

- **Invalid code**: "That code doesn't match any active cohort"
- **Expired cohort**: "This cohort has expired. Contact your training manager."
- **Cohort full**: "This cohort has reached its maximum capacity. Contact your training manager."

### 4. Admin Dashboard

```
┌─────────────────────────────────────────┐
│  Admin: Cohorts                         │
├─────────────────────────────────────────┤
│                                         │
│  [ + Create New Cohort ]                │
│                                         │
│  ┌───────────────────────────────────┐  │
│  │ Acme Corp Training                │  │
│  │ Code: abc-defg-hij                │  │
│  │ Students: 12/25 │ Expires: Mar 15 │  │
│  │ [Edit] [Copy Link]                │  │
│  └───────────────────────────────────┘  │
│                                         │
│  ┌───────────────────────────────────┐  │
│  │ Spring 2024 Cohort                │  │
│  │ Code: xyz-mnop-qrs                │  │
│  │ Students: 8/30  │ Expires: Apr 30 │  │
│  │ [Edit] [Copy Link]                │  │
│  └───────────────────────────────────┘  │
│                                         │
└─────────────────────────────────────────┘
```

---

## Cohort Code Generation

Format: `xxx-xxxx-xxx` (lowercase letters only, like Google Meet)

```go
func GenerateCohortCode() string {
    // 3-4-3 pattern = 10 letters = 26^10 ≈ 141 trillion combinations
    const letters = "abcdefghijklmnopqrstuvwxyz"
    // Use crypto/rand for generation
    // Check uniqueness before insert
}
```

---

## Implementation Phases

### Phase 1: Database Migration
1. Create new schema (cohorts, students, sessions)
2. Migrate existing learners table data (optional - could start fresh)
3. Update all DB functions to use new schema

### Phase 2: Authentication Flow
1. Add auth state to Model
2. Implement auth flow rendering (code entry, email entry, errors)
3. Implement auth message handlers
4. Update `Init()` to check session mapping

### Phase 3: Session Management
1. Create session → student mapping on successful auth
2. Load student data on returning session
3. Validate cohort expiry on each session init
4. Implement logout

### Phase 4: Owner Mode
1. Add `COURSEDEV_ADMIN_KEY` env var handling
2. Detect admin key in code entry
3. Set `IsOwner` flag, skip email entry
4. Gate presenter controls behind `IsOwner`

### Phase 5: Admin UI
1. Add `/admin` route
2. Implement cohort list view
3. Implement create cohort form
4. Implement edit cohort form
5. Add admin link to sidebar (owner only)

### Phase 6: Polish
1. Join link handling (`/join/:code`)
2. Student count display in admin
3. Copy join link button
4. Expired cohort cleanup (optional cron/manual)

---

## Security Considerations

- Cohort codes are not passwords - they're like meeting links (shareable within a cohort)
- Email is trusted without verification (acceptable for training context)
- Admin key should be strong and kept secret
- Session hijacking: standard Gotea session security applies
- SQL injection: use parameterized queries (already done)

---

## Testing Strategy

1. **Unit tests**: Cohort code generation, validation, expiry checks
2. **Integration tests**: Full auth flow via `tester.NewSession`
3. **Edge cases**:
   - Expired cohort mid-session
   - Max students exactly at limit
   - Same email, different cohorts
   - Cookie expiry → re-auth → progress restored

---

## Open Questions

None - all clarified in discussion.

---

## Estimated Scope

- Database: ~100 lines (schema + new functions)
- Auth flow: ~300 lines (handlers + rendering)
- Admin UI: ~200 lines (handlers + rendering)
- Model changes: ~50 lines
- Tests: ~200 lines

Total: ~850 lines of new/modified Go code
