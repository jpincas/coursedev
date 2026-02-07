package main

import (
	"fmt"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/go-tea/attributes"
	h "github.com/jpincas/go-tea/html"
)

// PollInfo contains poll state for quiz rendering
type PollInfo struct {
	IsActive     bool
	IsThisQuiz   bool
	IsInstructor bool
	HasResponded bool
	IsClosed     bool
	ShowResults  bool
	Responses    map[int]int
	TotalVotes   int
}

// renderQuiz renders a quiz block with interactive options
func renderQuiz(q *QuizBlock, blockIndex int, quizStates map[int]*QuizState, pollInfo *PollInfo, isFollowing bool) h.Element {
	// Look up state for this specific quiz
	state := quizStates[blockIndex]
	answered := state != nil && state.Answered

	// Check if this quiz is being polled
	isPollActive := pollInfo != nil && pollInfo.IsActive && pollInfo.IsThisQuiz

	// Render poll mode for instructor
	if isPollActive && pollInfo.IsInstructor {
		return renderQuizPollInstructor(q, blockIndex, pollInfo)
	}

	// Render poll mode for learners
	if isPollActive && !pollInfo.IsInstructor {
		return renderQuizPollLearner(q, blockIndex, state, pollInfo)
	}

	// When following a live session and quiz is not answered and no active poll,
	// show a locked "waiting" state so follower can't click before poll is opened
	if isFollowing && !answered {
		options := make([]h.Element, len(q.Options))
		for i, opt := range q.Options {
			options[i] = h.Div(a.Attrs(
				a.Class("py-4 px-5 border-2 border-zinc-700 rounded-xl bg-zinc-800/30 text-zinc-500 text-[0.95rem]")),
				h.Text(opt),
			)
		}
		return h.Div(a.Attrs(a.Class("bg-zinc-900 border border-zinc-800 rounded-2xl p-8 my-8 not-prose"), a.Custom("data-quiz-id", q.ID)),
			h.P(a.Attrs(a.Class("text-lg font-medium mb-6 text-zinc-100")), h.Text(q.Question)),
			h.Div(a.Attrs(a.Class("flex flex-col gap-3")),
				options...,
			),
			h.Div(a.Attrs(a.Class("mt-6 p-4 rounded-xl bg-blue-500/10 border border-blue-500 text-blue-400 text-sm text-center")),
				h.Text("Waiting for instructor to open this quiz as a live poll..."),
			),
		)
	}

	// Regular quiz mode - also show poll start button for instructor
	options := make([]h.Element, len(q.Options))
	for i, opt := range q.Options {
		classes := "py-4 px-5 border-2 border-zinc-700 rounded-xl bg-zinc-800/50 text-left cursor-pointer transition-all duration-200 text-zinc-300 text-[0.95rem]"
		if answered {
			if i == q.Answer {
				classes = "py-4 px-5 border-2 border-accent rounded-xl bg-accent/10 text-left text-accent text-[0.95rem]"
			} else if i == state.ChosenIdx {
				classes = "py-4 px-5 border-2 border-red-500 rounded-xl bg-red-500/10 text-left text-red-400 text-[0.95rem]"
			} else {
				classes = "py-4 px-5 border-2 border-zinc-700 rounded-xl bg-zinc-800/50 text-left text-zinc-300 text-[0.95rem]"
			}
		} else {
			classes += " hover:border-accent hover:text-zinc-100 hover:bg-accent-dim"
		}

		var clickHandler string
		if !answered {
			payload := map[string]int{
				"blockIndex": blockIndex,
				"answer":     i,
			}
			clickHandler = gt.SendBasicMessage("QUIZ_ANSWER", payload)
		}

		options[i] = h.Button(a.Attrs(
			a.Class(classes),
			a.OnClick(clickHandler),
			a.Disabled(answered)),
			h.Text(opt),
		)
	}

	var feedback h.Element
	if answered {
		var feedbackClass, resultClass, resultText string
		if state.Correct {
			feedbackClass = "mt-6 p-5 rounded-xl bg-accent/10 border border-accent"
			resultClass = "font-semibold mb-2 text-accent"
			resultText = "Correct!"
		} else {
			feedbackClass = "mt-6 p-5 rounded-xl bg-red-500/10 border border-red-500"
			resultClass = "font-semibold mb-2 text-red-400"
			resultText = "Incorrect"
		}
		feedback = h.Div(a.Attrs(a.Class(feedbackClass)),
			h.P(a.Attrs(a.Class(resultClass)), h.Text(resultText)),
			h.P(a.Attrs(a.Class("text-zinc-400 text-sm")), h.Text(q.Explanation)),
			h.Button(a.Attrs(
				a.Class("mt-4 py-2.5 px-5 bg-accent text-zinc-950 rounded-lg cursor-pointer font-medium transition-all duration-150 hover:bg-accent-hover border-none"),
				a.OnClick(gt.SendBasicMessage("QUIZ_RETRY", blockIndex))),
				h.Text("Try Again"),
			).RenderIf(!state.Correct),
		)
	} else {
		feedback = h.Div(a.Attrs()) // Empty placeholder
	}

	// Add poll start button for instructor (when not already polling)
	var pollControls h.Element
	if pollInfo != nil && pollInfo.IsInstructor && !isPollActive {
		pollControls = h.Div(a.Attrs(a.Class("flex gap-3 mt-4 pt-4 border-t border-zinc-800")),
			h.Button(a.Attrs(
				a.Class("py-2 px-4 rounded-lg bg-accent text-zinc-950 text-sm font-medium cursor-pointer transition-all duration-200 hover:bg-accent-hover border-none"),
				a.OnClick(gt.SendBasicMessage("START_POLL", map[string]int{"blockIndex": blockIndex}))),
				h.Text("\U0001F4CA Start Live Poll"),
			),
		)
	} else {
		pollControls = h.Div(a.Attrs())
	}

	return h.Div(a.Attrs(a.Class("bg-zinc-900 border border-zinc-800 rounded-2xl p-8 my-8 not-prose"), a.Custom("data-quiz-id", q.ID)),
		h.P(a.Attrs(a.Class("text-lg font-medium mb-6 text-zinc-100")), h.Text(q.Question)),
		h.Div(a.Attrs(a.Class("flex flex-col gap-3")),
			options...,
		),
		feedback,
		pollControls,
	)
}

// renderQuizPollInstructor renders the poll view for the instructor
func renderQuizPollInstructor(q *QuizBlock, blockIndex int, pollInfo *PollInfo) h.Element {
	// Build results bars
	results := make([]h.Element, len(q.Options))
	for i, opt := range q.Options {
		count := pollInfo.Responses[i]
		pct := 0
		if pollInfo.TotalVotes > 0 {
			pct = (count * 100) / pollInfo.TotalVotes
		}

		rowClass := "grid grid-cols-[1fr_2fr_auto] gap-4 items-center p-2 rounded-lg bg-zinc-950"
		if pollInfo.IsClosed && i == q.Answer {
			rowClass = "grid grid-cols-[1fr_2fr_auto] gap-4 items-center p-2 rounded-lg bg-green-500/15 border border-green-500"
		}

		results[i] = h.Div(a.Attrs(a.Class(rowClass)),
			h.Div(a.Attrs(a.Class("text-sm text-zinc-200")), h.Text(opt)),
			h.Div(a.Attrs(a.Class("h-6 bg-zinc-800 rounded-md overflow-hidden")),
				h.Div(a.Attrs(
					a.Class("h-full bg-accent rounded-md transition-all duration-300"),
					a.Custom("style", fmt.Sprintf("width: %d%%", pct))),
				),
			),
			h.Div(a.Attrs(a.Class("text-sm font-semibold text-zinc-400 min-w-[70px] text-right")),
				h.Text(fmt.Sprintf("%d (%d%%)", count, pct)),
			),
		)
	}

	// Control buttons
	var controls h.Element
	if pollInfo.IsClosed {
		controls = h.Div(a.Attrs(a.Class("flex gap-3 mt-4 pt-4 border-t border-zinc-800")),
			h.Button(a.Attrs(
				a.Class("py-2 px-4 rounded-lg bg-zinc-600 text-white text-sm font-medium cursor-pointer transition-all duration-200 hover:bg-zinc-500 border-none"),
				a.OnClick(gt.SendBasicMessageNoArgs("CLOSE_POLL"))),
				h.Text("Clear Poll"),
			),
		)
	} else {
		controls = h.Div(a.Attrs(a.Class("flex gap-3 mt-4 pt-4 border-t border-zinc-800")),
			h.Button(a.Attrs(
				a.Class("py-2 px-4 rounded-lg bg-zinc-800 text-zinc-300 text-sm font-medium cursor-pointer transition-all duration-200 hover:bg-zinc-700 border-none"),
				a.OnClick(gt.SendBasicMessageNoArgs("TOGGLE_POLL_RESULTS"))),
				h.Text(func() string {
					if pollInfo.ShowResults {
						return "Hide Results"
					}
					return "Show Results"
				}()),
			),
			h.Button(a.Attrs(
				a.Class("py-2 px-4 rounded-lg bg-green-500 text-white text-sm font-medium cursor-pointer transition-all duration-200 hover:bg-green-600 border-none"),
				a.OnClick(gt.SendBasicMessageNoArgs("END_POLL"))),
				h.Text("End Poll & Reveal"),
			),
		)
	}

	statusText := fmt.Sprintf("\U0001F4CA Live Poll \u2022 %d responses", pollInfo.TotalVotes)
	if pollInfo.IsClosed {
		statusText = fmt.Sprintf("\U0001F4CA Poll Ended \u2022 %d responses", pollInfo.TotalVotes)
	}

	return h.Div(a.Attrs(a.Class("bg-zinc-900 border-2 border-accent rounded-2xl p-6 my-8 not-prose")),
		h.Div(a.Attrs(a.Class("flex justify-between items-center mb-4")),
			h.Span(a.Attrs(a.Class("text-sm font-semibold text-accent")), h.Text(statusText)),
		),
		h.P(a.Attrs(a.Class("text-lg font-medium mb-6 text-zinc-100")), h.Text(q.Question)),
		h.Div(a.Attrs(a.Class("flex flex-col gap-3 my-6")),
			results...,
		),
		controls,
	)
}

// renderQuizPollLearner renders the poll view for learners
func renderQuizPollLearner(q *QuizBlock, blockIndex int, state *QuizState, pollInfo *PollInfo) h.Element {
	hasAnswered := pollInfo.HasResponded || (state != nil && state.BlockIndex == blockIndex && state.Answered)

	// If poll is closed and showing results, show the results
	if pollInfo.IsClosed && pollInfo.ShowResults {
		results := make([]h.Element, len(q.Options))
		for i, opt := range q.Options {
			count := pollInfo.Responses[i]
			pct := 0
			if pollInfo.TotalVotes > 0 {
				pct = (count * 100) / pollInfo.TotalVotes
			}

			isCorrect := i == q.Answer
			isUserChoice := state != nil && state.ChosenIdx == i

			rowClass := "grid grid-cols-[1fr_2fr_auto] gap-4 items-center p-2 rounded-lg bg-zinc-950"
			if isCorrect {
				rowClass = "grid grid-cols-[1fr_2fr_auto] gap-4 items-center p-2 rounded-lg bg-green-500/15 border border-green-500"
			} else if isUserChoice {
				rowClass = "grid grid-cols-[1fr_2fr_auto] gap-4 items-center p-2 rounded-lg bg-amber-400/15 border border-amber-400"
			}

			// Build label for this option
			var label h.Element
			if isCorrect && isUserChoice {
				label = h.Span(a.Attrs(a.Class("text-xs font-semibold text-green-400 mt-1")), h.Text("\u2713 Your answer \u2014 Correct!"))
			} else if isCorrect {
				label = h.Span(a.Attrs(a.Class("text-xs font-semibold text-green-400 mt-1")), h.Text("\u2713 Correct answer"))
			} else if isUserChoice {
				label = h.Span(a.Attrs(a.Class("text-xs font-semibold text-amber-400 mt-1")), h.Text("Your answer"))
			} else {
				label = h.Span(a.Attrs())
			}

			results[i] = h.Div(a.Attrs(a.Class(rowClass)),
				h.Div(a.Attrs(a.Class("text-sm text-zinc-200")),
					h.Div(a.Attrs(), h.Text(opt)),
					label,
				),
				h.Div(a.Attrs(a.Class("h-6 bg-zinc-800 rounded-md overflow-hidden")),
					h.Div(a.Attrs(
						a.Class("h-full bg-accent rounded-md transition-all duration-300"),
						a.Custom("style", fmt.Sprintf("width: %d%%", pct))),
					),
				),
				h.Div(a.Attrs(a.Class("text-sm font-semibold text-zinc-400 min-w-[70px] text-right")),
					h.Text(fmt.Sprintf("%d%%", pct)),
				),
			)
		}

		return h.Div(a.Attrs(a.Class("bg-zinc-900 border-2 border-accent rounded-2xl p-6 my-8 not-prose")),
			h.Div(a.Attrs(a.Class("flex justify-between items-center mb-4")),
				h.Span(a.Attrs(a.Class("text-sm font-semibold text-accent")), h.Text("\U0001F4CA Poll Results")),
			),
			h.P(a.Attrs(a.Class("text-lg font-medium mb-6 text-zinc-100")), h.Text(q.Question)),
			h.Div(a.Attrs(a.Class("flex flex-col gap-3 my-6")),
				results...,
			),
			h.P(a.Attrs(a.Class("mt-4 p-4 bg-zinc-950 rounded-lg italic text-zinc-400")), h.Text(q.Explanation)),
		)
	}

	// If learner has answered, show waiting state
	if hasAnswered {
		return h.Div(a.Attrs(a.Class("bg-zinc-900 border-2 border-accent rounded-2xl p-6 my-8 not-prose")),
			h.Div(a.Attrs(a.Class("flex justify-between items-center mb-4")),
				h.Span(a.Attrs(a.Class("text-sm font-semibold text-accent animate-pulse-soft")), h.Text("\U0001F4CA Live Poll")),
			),
			h.P(a.Attrs(a.Class("text-lg font-medium mb-6 text-zinc-100")), h.Text(q.Question)),
			h.Div(a.Attrs(a.Class("text-center py-8 text-accent text-lg")),
				h.Text("\u2713 Response submitted. Waiting for others..."),
			),
		)
	}

	// Show voting options
	options := make([]h.Element, len(q.Options))
	for i, opt := range q.Options {
		payload := map[string]int{
			"blockIndex": blockIndex,
			"answer":     i,
		}
		options[i] = h.Button(a.Attrs(
			a.Class("w-full py-4 px-5 border-2 border-zinc-700 rounded-xl bg-zinc-800/50 text-left cursor-pointer transition-all duration-200 text-zinc-300 text-[0.95rem] mb-2 hover:border-accent hover:text-zinc-100 hover:bg-accent-dim"),
			a.OnClick(gt.SendBasicMessage("QUIZ_ANSWER", payload))),
			h.Text(opt),
		)
	}

	return h.Div(a.Attrs(a.Class("bg-zinc-900 border-2 border-accent rounded-2xl p-6 my-8 not-prose")),
		h.Div(a.Attrs(a.Class("flex justify-between items-center mb-4")),
			h.Span(a.Attrs(a.Class("text-sm font-semibold text-accent animate-pulse-soft")), h.Text("\U0001F4CA Live Poll")),
		),
		h.P(a.Attrs(a.Class("text-lg font-medium mb-6 text-zinc-100")), h.Text(q.Question)),
		h.Div(a.Attrs(a.Class("flex flex-col gap-3")),
			options...,
		),
	)
}

// renderMermaid renders a Mermaid diagram block
func renderMermaid(b *MermaidBlock) h.Element {
	return h.Div(a.Attrs(a.Class("my-8 p-6 bg-zinc-900 rounded-xl border border-zinc-800 overflow-x-auto not-prose")),
		h.Div(a.Attrs(a.Class("mermaid")),
			h.Text(b.Source),
		),
	)
}

// renderMath renders a KaTeX math block
func renderMath(b *MathBlock) h.Element {
	return h.Div(a.Attrs(a.Class("my-8 p-6 bg-zinc-900 rounded-xl border border-zinc-800 text-center not-prose")),
		h.Div(a.Attrs(a.Class("katex-block")),
			h.Text(b.Source),
		),
	)
}

// renderCallout renders a callout/admonition block
func renderCallout(b *CalloutBlock) h.Element {
	icon := getCalloutIcon(b.Type)

	borderColor := "border-zinc-600"
	switch b.Type {
	case "info":
		borderColor = "border-blue-500"
	case "warning":
		borderColor = "border-amber-400"
	case "tip":
		borderColor = "border-accent"
	case "danger":
		borderColor = "border-red-500"
	}

	return h.Div(a.Attrs(a.Class(fmt.Sprintf("py-5 px-6 rounded-xl my-8 border-l-4 bg-zinc-900 not-prose %s", borderColor))),
		h.Div(a.Attrs(a.Class("flex items-center gap-2.5 font-semibold mb-2 text-zinc-100")),
			h.Span(a.Attrs(a.Class("text-lg")), h.Text(icon)),
			h.Span(a.Attrs(), h.Text(b.Title)),
		),
		h.Div(a.Attrs(a.Class("text-zinc-400 text-[0.95rem]")),
			h.Text(b.Content),
		),
	)
}

// getCalloutIcon returns an emoji icon for the callout type
func getCalloutIcon(calloutType string) string {
	switch calloutType {
	case "info":
		return "\u2139\uFE0F"
	case "warning":
		return "\u26A0\uFE0F"
	case "tip":
		return "\U0001F4A1"
	case "danger":
		return "\U0001F6A8"
	case "note":
		return "\U0001F4DD"
	default:
		return "\U0001F4CC"
	}
}

// renderAnnotatedImage renders an image with interactive hotspots
func renderAnnotatedImage(b *AnnotatedImageBlock, activeHotspot string) h.Element {
	hotspots := make([]h.Element, len(b.Hotspots))
	for i, spot := range b.Hotspots {
		spotID := fmt.Sprintf("%s-hotspot-%d", b.ID, i)
		isActive := activeHotspot == spotID

		classes := "absolute cursor-pointer group"

		hotspots[i] = h.Div(a.Attrs(
			a.Class(classes),
			a.Custom("style", fmt.Sprintf("left: %s; top: %s", spot.X, spot.Y)),
			a.OnClick(gt.SendBasicMessage("HOTSPOT_TOGGLE", spotID))),
			h.Span(a.Attrs(a.Class("flex items-center justify-center w-7 h-7 bg-accent text-zinc-950 rounded-full font-bold -translate-x-1/2 -translate-y-1/2 transition-all duration-150 text-base group-hover:scale-110 group-hover:shadow-[0_0_0_4px_rgba(0,217,192,0.15)]")), h.Text("+")),
			h.Div(a.Attrs(a.Class("absolute top-9 left-1/2 -translate-x-1/2 bg-zinc-900 border border-zinc-800 rounded-xl p-4 min-w-[220px] shadow-xl z-10")),
				h.Strong(a.Attrs(a.Class("block mb-1 text-zinc-100")), h.Text(spot.Label)),
				h.P(a.Attrs(a.Class("text-zinc-400 text-sm m-0")), h.Text(spot.Detail)),
			).RenderIf(isActive),
		)
	}

	return h.Div(a.Attrs(a.Class("my-8 not-prose")),
		h.Div(a.Attrs(a.Class("relative inline-block")),
			h.Img(a.Attrs(
				a.Src(b.Src),
				a.Alt(b.Alt),
				a.Class("max-w-full h-auto rounded-xl border border-zinc-800"))),
			h.Div(a.Attrs(a.Class("absolute inset-0")),
				hotspots...,
			),
		),
	)
}

// renderTerminalReplay renders an asciinema terminal replay block
func renderTerminalReplay(b *TerminalReplayBlock) h.Element {
	autoplayAttr := "false"
	if b.Autoplay {
		autoplayAttr = "true"
	}

	return h.Div(a.Attrs(a.Class("my-8 not-prose")),
		h.Div(a.Attrs(a.Class("font-semibold mb-3 text-zinc-100")),
			h.Text(b.Title),
		).RenderIf(b.Title != ""),
		h.Div(a.Attrs(
			a.Class("asciinema-player"),
			a.Custom("data-src", b.Src),
			a.Custom("data-autoplay", autoplayAttr),
			a.Custom("data-speed", fmt.Sprintf("%.1f", b.Speed))),
		),
	)
}

// renderExercise renders a code exercise block
func renderExercise(b *ExerciseBlock) h.Element {
	return h.Div(a.Attrs(a.Class("bg-zinc-900 border border-zinc-800 rounded-2xl p-8 my-8 not-prose"), a.Custom("data-exercise-id", b.ID)),
		h.Div(a.Attrs(a.Class("mb-5 text-zinc-100 font-medium")),
			h.Text(b.Prompt),
		),
		h.Div(a.Attrs(),
			h.TextArea(a.Attrs(
				a.Class("w-full font-mono text-sm p-5 border border-zinc-700 rounded-xl bg-zinc-950 text-zinc-50 resize-y min-h-[200px] focus:outline-none focus:border-accent"),
				a.Custom("data-language", b.Language),
				a.Rows(10)),
				h.Text(b.Starter),
			),
		),
		h.Div(a.Attrs(a.Class("mt-5")),
			h.Button(a.Attrs(a.Class("py-3 px-6 bg-accent text-zinc-950 rounded-lg cursor-pointer font-semibold transition-all duration-150 hover:bg-accent-hover border-none")),
				h.Text("Submit"),
			),
		),
	)
}
