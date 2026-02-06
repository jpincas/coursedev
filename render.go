package main

import (
	"fmt"
	"log"
	"regexp"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/go-tea/attributes"
	h "github.com/jpincas/go-tea/html"
)

// Block marker pattern for splitting narrative HTML
var blockMarkerPattern = regexp.MustCompile(`<div data-block-index="(\d+)"></div>`)

// renderLayout renders the full page layout
func renderLayout(m *Model) h.Element {
	// If not authenticated, show auth flow
	if !m.IsAuthenticated() {
		return h.Html(a.Attrs(a.Lang("en")),
			h.Head(a.Attrs(),
				h.Meta(a.Attrs(a.Charset("UTF-8"))),
				h.Meta(a.Attrs(
					a.Name("viewport"),
					a.Content("width=device-width, initial-scale=1.0"))),
				h.Title(a.Attrs(), h.Text("Training - Login")),
				h.Link(a.Attrs(
					a.Rel("preconnect"),
					a.Href("https://fonts.googleapis.com"))),
				h.Link(a.Attrs(
					a.Rel("preconnect"),
					a.Href("https://fonts.gstatic.com"),
					a.Custom("crossorigin", ""))),
				h.Link(a.Attrs(
					a.Rel("stylesheet"),
					a.Href("https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap"))),
				h.Link(a.Attrs(
					a.Rel("stylesheet"),
					a.Href("static/css/main.css"))),
			),
			h.Body(a.Attrs(a.Class("bg-zinc-950 text-zinc-50 font-sans antialiased min-h-screen flex items-center justify-center p-8")),
				renderAuthFlow(m),
				h.Script(a.Attrs(a.Src("static/js/main.js"))),
			))
	}

	fontClass := "font-medium"
	switch m.Preferences.FontSize {
	case "small":
		fontClass = "font-small"
	case "large":
		fontClass = "font-large"
	}

	return h.Html(a.Attrs(a.Lang("en")),
		h.Head(a.Attrs(),
			h.Meta(a.Attrs(a.Charset("UTF-8"))),
			h.Meta(a.Attrs(
				a.Name("viewport"),
				a.Content("width=device-width, initial-scale=1.0"))),
			h.Title(a.Attrs(), h.Text("Training")),
			// Fonts
			h.Link(a.Attrs(
				a.Rel("preconnect"),
				a.Href("https://fonts.googleapis.com"))),
			h.Link(a.Attrs(
				a.Rel("preconnect"),
				a.Href("https://fonts.gstatic.com"),
				a.Custom("crossorigin", ""))),
			h.Link(a.Attrs(
				a.Rel("stylesheet"),
				a.Href("https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap"))),
			// Stylesheets
			h.Link(a.Attrs(
				a.Rel("stylesheet"),
				a.Href("static/css/main.css"))),
			h.Link(a.Attrs(
				a.Rel("stylesheet"),
				a.Href("static/vendor/katex.min.css"))),
			// Vendor scripts (loaded in head for availability)
			h.Script(a.Attrs(a.Src("static/vendor/mermaid.min.js"))),
			h.Script(a.Attrs(a.Src("static/vendor/katex.min.js"))),
		),
		h.Body(a.Attrs(
			a.Class(fmt.Sprintf("bg-zinc-950 text-zinc-50 font-sans antialiased min-h-screen %s", fontClass))),
			h.Div(a.Attrs(a.Class("flex flex-col min-h-screen md:flex-row")),
				renderSidebar(m),
				h.Main(a.Attrs(a.Class("flex-1 p-6 md:py-10 md:px-12 max-w-[900px] overflow-y-auto")),
					h.Div(a.Attrs(a.Id("view")),
						h.UnsafeRaw(string(m.RenderRoute(m))))),
			),
			h.Script(a.Attrs(a.Src("static/js/main.js"))),
			h.Script(a.Attrs(a.Src("static/js/training.js"))),
		))
}

// renderSidebarControls renders the control area at the bottom of the sidebar
func renderSidebarControls(m *Model) h.Element {
	// View toggle: switches between course material and progress view
	var viewToggleHref, viewToggleLabel string
	if m.Router.Route == "/modules" {
		viewToggleHref = "/"
		viewToggleLabel = "Back to Course"
	} else {
		viewToggleHref = "/modules"
		viewToggleLabel = "View Progress"
	}
	viewToggle := h.A(a.Attrs(
		a.Href(viewToggleHref),
		a.Class("block py-3 px-4 bg-zinc-800/50 border border-zinc-700 rounded-lg text-zinc-400 no-underline text-sm font-medium text-center transition-all duration-150 hover:bg-zinc-800 hover:border-accent hover:text-accent")),
		h.Text(viewToggleLabel),
	)

	// Presenter/follower controls
	var sessionControl h.Element
	if m.IsPresenting {
		sessionControl = h.Div(a.Attrs(a.Class("flex items-center justify-between py-2 px-3 bg-accent-dim rounded-lg border border-accent")),
			h.Span(a.Attrs(a.Class("text-sm font-medium text-accent")), h.Text("Presenting")),
			h.Button(a.Attrs(
				a.Class("py-1.5 px-3 text-sm rounded-md bg-zinc-800 text-zinc-300 font-medium cursor-pointer transition-all duration-150 hover:bg-zinc-700 border-none"),
				a.OnClick(gt.SendBasicMessageNoArgs("STOP_PRESENTING"))),
				h.Text("Stop"),
			),
		)
	} else if m.FollowingLive {
		sessionControl = h.Div(a.Attrs(a.Class("flex items-center justify-between py-2 px-3 bg-blue-500/10 rounded-lg border border-blue-500")),
			h.Span(a.Attrs(a.Class("text-sm font-medium text-blue-400")), h.Text("Following")),
			h.Button(a.Attrs(
				a.Class("py-1.5 px-3 text-sm rounded-md bg-zinc-800 text-zinc-300 font-medium cursor-pointer transition-all duration-150 hover:bg-zinc-700 border-none"),
				a.OnClick(gt.SendBasicMessageNoArgs("LEAVE_SESSION"))),
				h.Text("Self-paced"),
			),
		)
	} else {
		// Check for available live session for student's cohort
		var hasLiveSession bool
		if m.CohortID != nil {
			hasLiveSession = GetLiveSessionForCohort(*m.CohortID) != nil
		}
		if hasLiveSession {
			sessionControl = h.Div(a.Attrs(a.Class("flex items-center justify-between py-2 px-3 bg-amber-400/10 rounded-lg border border-amber-400")),
				h.Span(a.Attrs(a.Class("text-sm font-medium text-amber-400")), h.Text("Live Session")),
				h.Button(a.Attrs(
					a.Class("py-1.5 px-3 text-sm rounded-md bg-accent text-zinc-950 font-medium cursor-pointer transition-all duration-150 hover:bg-accent-hover border-none"),
					a.OnClick(gt.SendBasicMessageNoArgs("JOIN_SESSION"))),
					h.Text("Join"),
				),
			)
		} else if m.IsOwner {
			// Owner can start presenting — render cohort selection
			sessionControl = renderPresenterCohortSelector()
		} else {
			sessionControl = h.Span(a.Attrs())
		}
	}

	// Admin link for owner
	var adminLink h.Element
	if m.IsOwner {
		adminLink = h.A(a.Attrs(
			a.Href("/admin"),
			a.Class("block py-3 px-4 bg-accent-dim border border-accent rounded-lg text-accent no-underline text-sm font-medium text-center transition-all duration-150 hover:bg-accent/20")),
			h.Text("Admin"),
		)
	} else {
		adminLink = h.Span(a.Attrs())
	}

	// Logout button
	logoutBtn := h.Button(a.Attrs(
		a.Class("py-1.5 px-3 text-sm rounded-md bg-zinc-800 text-zinc-300 font-medium cursor-pointer transition-all duration-150 hover:bg-zinc-700 border-none mt-2"),
		a.OnClick(gt.SendBasicMessageNoArgs("LOGOUT"))),
		h.Text("Logout"),
	)

	return h.Div(a.Attrs(a.Class("mt-auto pt-4 border-t border-zinc-800 flex flex-row flex-wrap gap-3 md:flex-col")),
		viewToggle,
		sessionControl,
		adminLink,
		logoutBtn,
	)
}

// renderPresenterCohortSelector renders a list of active cohorts for the admin to present to
func renderPresenterCohortSelector() h.Element {
	cohorts, err := GetAllCohorts()
	if err != nil {
		LogDBError("GetAllCohorts (presenter selector)", err)
		return h.Span(a.Attrs())
	}

	// Filter to non-expired cohorts
	var activeCohorts []*Cohort
	for _, c := range cohorts {
		if !c.IsExpired() {
			activeCohorts = append(activeCohorts, c)
		}
	}

	if len(activeCohorts) == 0 {
		return h.Div(a.Attrs(a.Class("py-2 px-3 bg-zinc-800/50 rounded-lg border border-zinc-700 text-zinc-500 text-sm")),
			h.Text("No active cohorts"),
		)
	}

	if len(activeCohorts) == 1 {
		c := activeCohorts[0]
		return h.Div(a.Attrs(a.Class("flex items-center justify-between py-2 px-3 bg-zinc-800/50 rounded-lg border border-zinc-700")),
			h.Span(a.Attrs(a.Class("text-sm text-zinc-400 truncate mr-2")), h.Text(c.Name)),
			h.Button(a.Attrs(
				a.Class("py-1.5 px-3 text-sm rounded-md bg-zinc-800 text-zinc-300 font-medium cursor-pointer transition-all duration-150 hover:bg-zinc-700 border-none shrink-0"),
				a.OnClick(gt.SendBasicMessage("START_PRESENTING", c.ID.String()))),
				h.Text("Present"),
			),
		)
	}

	// Multiple cohorts — show a button per cohort
	buttons := make([]h.Element, len(activeCohorts))
	for i, c := range activeCohorts {
		buttons[i] = h.Button(a.Attrs(
			a.Class("w-full py-2 px-3 text-sm rounded-md bg-zinc-800 text-zinc-300 font-medium cursor-pointer transition-all duration-150 hover:bg-zinc-700 border-none text-left"),
			a.OnClick(gt.SendBasicMessage("START_PRESENTING", c.ID.String()))),
			h.Text(fmt.Sprintf("Present to %s", c.Name)),
		)
	}
	return h.Div(a.Attrs(a.Class("flex flex-col gap-2 py-2 px-3 bg-zinc-800/50 rounded-lg border border-zinc-700")),
		append([]h.Element{
			h.Span(a.Attrs(a.Class("text-xs text-zinc-500 font-medium uppercase tracking-wider")), h.Text("Present to")),
		}, buttons...)...,
	)
}

// renderSidebar renders the module/page navigation sidebar
func renderSidebar(m *Model) h.Element {
	mod := m.currentModule()
	currentPageIdx := m.CurrentPage

	if mod == nil {
		return h.Aside(a.Attrs(a.Class("w-full relative border-b border-zinc-800 p-6 overflow-y-auto flex flex-col gap-6 max-h-[50vh] bg-zinc-900/80 backdrop-blur-sm md:w-72 md:h-screen md:sticky md:top-0 md:border-r md:border-b-0 md:max-h-none md:shrink-0")),
			h.Div(a.Attrs(a.Class("flex flex-col gap-3")),
				h.P(a.Attrs(), h.Text("No modules loaded")),
			),
			renderSidebarControls(m),
		)
	}

	// Progress is always from learner's own progress, not presenter's
	progress := m.Progress[mod.ID]

	// Determine if navigation is disabled (when following a live session)
	navDisabled := m.FollowingLive

	pages := make([]h.Element, len(mod.Pages))
	for i, page := range mod.Pages {
		classes := "flex items-center gap-3 py-2.5 px-3 rounded-lg text-zinc-400 no-underline text-sm transition-all duration-150 border border-transparent external"
		if i == currentPageIdx {
			classes = "flex items-center gap-3 py-2.5 px-3 rounded-lg text-accent no-underline text-sm transition-all duration-150 border border-accent/30 bg-accent-dim external"
		}
		if navDisabled {
			classes += " opacity-60 cursor-not-allowed"
		} else {
			classes += " hover:bg-zinc-800/50 hover:text-zinc-200"
		}
		// Check if page has been viewed (from learner's own progress)
		viewed := progress != nil && progress.PagesViewed[i]
		if viewed && i != currentPageIdx {
			classes += " text-zinc-200"
		}

		title := page.Meta.Title
		if title == "" {
			title = fmt.Sprintf("Page %d", i+1)
		}

		// Build page item with number and completion indicator
		pageNum := fmt.Sprintf("%d", i+1)

		// Page number styling
		pageNumClass := "flex items-center justify-center w-6 h-6 bg-zinc-800 rounded-full text-xs font-semibold shrink-0"
		if i == currentPageIdx {
			pageNumClass = "flex items-center justify-center w-6 h-6 bg-accent text-zinc-950 rounded-full text-xs font-semibold shrink-0"
		} else if viewed {
			pageNumClass = "flex items-center justify-center w-6 h-6 bg-zinc-700 rounded-full text-xs font-semibold shrink-0"
		}

		var statusIcon h.Element
		if m.IsOwner {
			// Owner sees no completion indicators
			statusIcon = h.Span(a.Attrs(a.Class("text-xs w-5 text-center shrink-0")), h.Text(""))
		} else if m.isPageComplete(mod.ID, i) {
			statusIcon = h.Span(a.Attrs(a.Class("text-xs w-5 text-center shrink-0 text-accent")), h.Text("\u2713"))
		} else if m.pageHasIncompleteQuiz(mod.ID, i) {
			statusIcon = h.Span(a.Attrs(a.Class("text-xs w-5 text-center shrink-0 text-amber-400")), h.Text("\u25CB"))
		} else if viewed {
			statusIcon = h.Span(a.Attrs(a.Class("text-xs w-5 text-center shrink-0 text-accent")), h.Text("\u2713"))
		} else {
			statusIcon = h.Span(a.Attrs(a.Class("text-xs w-5 text-center shrink-0")), h.Text(""))
		}

		// Disable click when following instructor
		var clickHandler string
		if !navDisabled {
			clickHandler = gt.SendBasicMessage("NAV_PAGE", i) + "; return false;"
		} else {
			clickHandler = "return false;"
		}

		pages[i] = h.A(a.Attrs(
			a.Href("#"),
			a.Class(classes),
			a.OnClick(clickHandler)),
			h.Span(a.Attrs(a.Class(pageNumClass)), h.Text(pageNum)),
			h.Span(a.Attrs(a.Class("flex-1 truncate")), h.Text(title)),
			statusIcon,
		)
	}

	// Module completion status (students only)
	var moduleStatus h.Element
	if m.IsOwner {
		// Owner just sees page count
		moduleStatus = h.Div(a.Attrs(a.Class("flex items-center gap-2 text-xs text-zinc-500 py-2 px-3 bg-zinc-800/50 rounded-md")),
			h.Span(a.Attrs(),
				h.Text(fmt.Sprintf("%d pages", len(mod.Pages)))),
		)
	} else if progress != nil && progress.CompletedAt != nil {
		moduleStatus = h.Div(a.Attrs(a.Class("flex items-center gap-2 text-xs text-accent py-2 px-3 bg-accent-dim rounded-md")),
			h.Span(a.Attrs(a.Class("text-sm")), h.Text("\u2713")),
			h.Span(a.Attrs(), h.Text("Module Complete")),
		)
	} else if m.moduleNeedsQuizzes(mod.ID) {
		moduleStatus = h.Div(a.Attrs(a.Class("flex items-center gap-2 text-xs text-amber-400 py-2 px-3 bg-amber-400/10 rounded-md")),
			h.Span(a.Attrs(a.Class("text-sm")), h.Text("\u25CB")),
			h.Span(a.Attrs(), h.Text("Quizzes needed")),
		)
	} else {
		pagesViewed := 0
		if progress != nil {
			pagesViewed = len(progress.PagesViewed)
		}
		moduleStatus = h.Div(a.Attrs(a.Class("flex items-center gap-2 text-xs text-zinc-500 py-2 px-3 bg-zinc-800/50 rounded-md")),
			h.Span(a.Attrs(),
				h.Text(fmt.Sprintf("%d of %d pages", pagesViewed, len(mod.Pages)))),
		)
	}

	// Add indicator when following presenter
	var followIndicator h.Element
	if navDisabled {
		followIndicator = h.Div(a.Attrs(a.Class("text-xs text-blue-400 p-2 mt-2 bg-blue-500/10 rounded-md text-center")),
			h.Text("Navigation controlled by presenter"),
		)
	} else {
		followIndicator = h.Span(a.Attrs())
	}

	return h.Aside(a.Attrs(a.Class("w-full relative border-b border-zinc-800 p-6 overflow-y-auto flex flex-col gap-6 max-h-[50vh] bg-zinc-900/80 backdrop-blur-sm md:w-72 md:h-screen md:sticky md:top-0 md:border-r md:border-b-0 md:max-h-none md:shrink-0")),
		h.Div(a.Attrs(a.Class("flex flex-col gap-3")),
			h.H2(a.Attrs(a.Class("text-lg font-semibold text-zinc-50 tracking-tight")), h.Text(mod.Meta.Title)),
			moduleStatus,
			followIndicator,
		),
		h.Nav(a.Attrs(a.Class("flex flex-col gap-1")),
			pages...,
		),
		renderSidebarControls(m),
	)
}

// renderCurrentPage renders the current page content
func (m *Model) renderCurrentPage(s gt.State) []byte {
	mdl := s.(*Model)

	page := mdl.currentPage()
	if page == nil {
		return renderEmptyState().Bytes()
	}

	mod := mdl.currentModule()

	log.Printf("renderCurrentPage: SessionID=%s IsPresenting=%v FollowingLive=%v module=%s page=%d",
		mdl.SessionID, mdl.IsPresenting, mdl.FollowingLive, mdl.CurrentModule, mdl.CurrentPage)

	return renderPageContentWithNav(page, mod, mdl.CurrentPage, mdl.ActiveQuizzes, mdl.ActiveHotspot, mdl.IsPresenting, mdl.FollowingLive, mdl).Bytes()
}

// renderModuleSelector renders the combined module selection and progress view
func (m *Model) renderModuleSelector(s gt.State) []byte {
	mdl := s.(*Model)

	modules := make([]h.Element, 0, len(globalCourse.ModuleOrder))

	// Calculate overall progress for the header (students only)
	totalModules := len(globalCourse.ModuleOrder)
	completedModules := 0
	overallPct := 0
	if !mdl.IsOwner {
		for _, id := range globalCourse.ModuleOrder {
			if p := mdl.Progress[id]; p != nil && p.CompletedAt != nil {
				completedModules++
			}
		}
		if totalModules > 0 {
			overallPct = (completedModules * 100) / totalModules
		}
	}

	for i, id := range globalCourse.ModuleOrder {
		mod := globalModules[id]
		if mod == nil {
			continue
		}

		// Owner always has access; students must meet prerequisites
		available := mdl.IsOwner || globalCourse.PrerequisitesMet(id, mdl.Progress)
		progress := mdl.Progress[id]

		// Build card classes based on state
		cardClasses := "bg-zinc-900 border border-zinc-800 rounded-2xl p-6 transition-all duration-200 relative cursor-pointer hover:border-accent hover:bg-accent-dim"
		if !available {
			cardClasses = "bg-zinc-900 border border-zinc-800 rounded-2xl p-6 transition-all duration-200 relative opacity-50 cursor-not-allowed"
		}
		if id == mdl.CurrentModule {
			cardClasses += " border-accent"
		}

		// Module number styling
		numClass := "flex items-center justify-center w-8 h-8 bg-accent text-zinc-950 rounded-full font-bold text-sm"
		if !available {
			numClass = "flex items-center justify-center w-8 h-8 bg-zinc-700 text-zinc-400 rounded-full font-bold text-sm"
		}

		// Status badge and progress bar — only for students
		var statusEl h.Element
		var progressBar h.Element
		if mdl.IsOwner {
			// Owner sees page count only, no progress state
			statusEl = h.Span(a.Attrs(a.Class("text-xs py-1 px-2.5 rounded-full bg-zinc-800 text-zinc-500 font-medium")),
				h.Text(fmt.Sprintf("%d pages", len(mod.Pages))))
			progressBar = h.Span(a.Attrs()) // empty
		} else {
			// Student progress
			pagesViewed := 0
			pct := 0
			if progress != nil {
				pagesViewed = len(progress.PagesViewed)
				if len(mod.Pages) > 0 {
					pct = (pagesViewed * 100) / len(mod.Pages)
				}
			}

			fillClass := "h-full bg-accent rounded-full transition-all duration-300"
			if progress != nil && progress.CompletedAt != nil {
				statusEl = h.Span(a.Attrs(a.Class("text-xs py-1 px-2.5 rounded-full bg-accent/15 text-accent font-medium")), h.Text("Completed"))
			} else if mdl.moduleNeedsQuizzes(id) {
				statusEl = h.Span(a.Attrs(a.Class("text-xs py-1 px-2.5 rounded-full bg-amber-400/15 text-amber-400 font-medium")), h.Text("Quizzes needed"))
			} else if progress != nil && progress.Started {
				fillClass = "h-full bg-blue-500 rounded-full transition-all duration-300"
				statusEl = h.Span(a.Attrs(a.Class("text-xs py-1 px-2.5 rounded-full bg-blue-500/15 text-blue-400 font-medium")),
					h.Text(fmt.Sprintf("%d/%d pages", pagesViewed, len(mod.Pages))))
			} else if !available {
				statusEl = h.Span(a.Attrs(a.Class("text-xs py-1 px-2.5 rounded-full bg-zinc-800 text-zinc-500 font-medium")), h.Text("Locked"))
			} else {
				statusEl = h.Span(a.Attrs(a.Class("text-xs py-1 px-2.5 rounded-full bg-zinc-800 text-zinc-500 font-medium")), h.Text("Not started"))
			}

			progressBar = h.Div(a.Attrs(a.Class("h-1 bg-zinc-800 rounded-full overflow-hidden mb-3")),
				h.Div(a.Attrs(
					a.Class(fillClass),
					a.Custom("style", fmt.Sprintf("width: %d%%", pct))),
				),
			)
		}

		var clickHandler string
		if available {
			clickHandler = gt.SendBasicMessage("NAV_MODULE", id)
		}

		modules = append(modules, h.Div(a.Attrs(
			a.Class(cardClasses),
			a.OnClick(clickHandler)),
			h.Div(a.Attrs(a.Class("flex items-center gap-3 mb-3")),
				h.Span(a.Attrs(a.Class(numClass)), h.Text(fmt.Sprintf("%d", i+1))),
				statusEl,
			),
			h.H3(a.Attrs(a.Class("text-xl font-semibold mb-2 text-zinc-50 tracking-tight")), h.Text(mod.Meta.Title)),
			h.P(a.Attrs(a.Class("text-zinc-400 text-sm mb-4 leading-relaxed")), h.Text(mod.Meta.Description)),
			progressBar,
			h.Div(a.Attrs(a.Class("flex gap-4 text-xs text-zinc-500")),
				h.Span(a.Attrs(),
					h.Text(mod.Meta.EstimatedDuration)),
				h.Span(a.Attrs(),
					h.Text(fmt.Sprintf("%d pages", len(mod.Pages)))),
			),
		))
	}

	// Header — students see progress stats, owner just sees title
	var header h.Element
	if mdl.IsOwner {
		header = h.Div(a.Attrs(a.Class("mb-10")),
			h.H1(a.Attrs(a.Class("text-3xl md:text-4xl font-bold tracking-tight text-zinc-50")), h.Text("Course Modules")),
		)
	} else {
		header = h.Div(a.Attrs(a.Class("mb-10 flex flex-col items-start gap-4 md:flex-row md:items-center md:justify-between")),
			h.Div(a.Attrs(),
				h.H1(a.Attrs(a.Class("text-3xl md:text-4xl font-bold tracking-tight text-zinc-50")), h.Text("Course Modules")),
			),
			h.Div(a.Attrs(a.Class("flex gap-8")),
				h.Div(a.Attrs(a.Class("text-center")),
					h.Span(a.Attrs(a.Class("block text-3xl font-bold text-accent tracking-tight")), h.Text(fmt.Sprintf("%d%%", overallPct))),
					h.Span(a.Attrs(a.Class("text-xs text-zinc-500 uppercase tracking-widest")), h.Text("Complete")),
				),
				h.Div(a.Attrs(a.Class("text-center")),
					h.Span(a.Attrs(a.Class("block text-3xl font-bold text-accent tracking-tight")), h.Text(fmt.Sprintf("%d/%d", completedModules, totalModules))),
					h.Span(a.Attrs(a.Class("text-xs text-zinc-500 uppercase tracking-widest")), h.Text("Modules")),
				),
			),
		)
	}

	return h.Div(a.Attrs(a.Class("py-4")),
		header,
		h.Div(a.Attrs(a.Class("flex flex-col gap-4")),
			modules...,
		),
	).Bytes()
}

// renderPageContentWithNav renders a page with navigation controls
func renderPageContentWithNav(page *Page, mod *Module, currentIdx int, quizStates map[int]*QuizState, activeHotspot string, isPresenting bool, isFollowing bool, mdl *Model) h.Element {
	// Debug: log what message type the nav buttons will use
	log.Printf("renderPageContentWithNav: sessionID=%s isPresenting=%v isFollowing=%v", mdl.SessionID, isPresenting, isFollowing)

	// Split the narrative HTML at block markers
	segments := splitHTMLAtBlockMarkers(page.NarrativeHTML)
	children := make([]h.Element, 0, len(segments)+len(page.Blocks))

	for i, segment := range segments {
		if len(segment) > 0 {
			children = append(children, h.UnsafeRaw(string(segment)))
		}
		if i < len(page.Blocks) {
			children = append(children,
				renderBlock(page.Blocks[i], i, quizStates, activeHotspot, isPresenting, mdl))
		}
	}

	// Build navigation controls
	hasPrev := currentIdx > 0
	hasNext := mod != nil && currentIdx < len(mod.Pages)-1

	// Navigation messages are always the same - the handlers check if presenting
	prevMsg := "PREV_PAGE"
	nextMsg := "NEXT_PAGE"

	var prevBtn, nextBtn h.Element

	// When following instructor, disable navigation
	if isFollowing {
		prevBtn = h.Div(a.Attrs(a.Class("flex items-center gap-4 py-4 px-5 bg-zinc-900 border border-zinc-800 rounded-xl flex-1 max-w-full md:max-w-[280px] text-left opacity-30 cursor-default")))
		nextBtn = h.Div(a.Attrs(a.Class("flex items-center gap-4 py-4 px-5 bg-zinc-900 border border-zinc-800 rounded-xl flex-1 max-w-full md:max-w-[280px] text-right md:ml-auto opacity-30 cursor-default")))
	} else {
		if hasPrev {
			prevTitle := mod.Pages[currentIdx-1].Meta.Title
			if prevTitle == "" {
				prevTitle = fmt.Sprintf("Page %d", currentIdx)
			}
			prevBtn = h.Button(a.Attrs(
				a.Class("flex items-center gap-4 py-4 px-5 bg-zinc-900 border border-zinc-800 rounded-xl cursor-pointer transition-all duration-200 flex-1 max-w-full md:max-w-[280px] text-left hover:border-accent hover:bg-accent-dim"),
				a.OnClick(gt.SendBasicMessageNoArgs(prevMsg))),
				h.Span(a.Attrs(a.Class("text-xl text-accent shrink-0")), h.Text("\u2190")),
				h.Span(a.Attrs(a.Class("flex flex-col gap-1 min-w-0")),
					h.Span(a.Attrs(a.Class("text-xs text-zinc-500 uppercase tracking-widest")), h.Text("Previous")),
					h.Span(a.Attrs(a.Class("text-sm font-medium text-zinc-100 truncate")), h.Text(prevTitle)),
				),
			)
		} else {
			prevBtn = h.Div(a.Attrs(a.Class("flex items-center gap-4 py-4 px-5 bg-zinc-900 border border-zinc-800 rounded-xl flex-1 max-w-full md:max-w-[280px] text-left opacity-30 cursor-default")))
		}

		if hasNext {
			nextTitle := mod.Pages[currentIdx+1].Meta.Title
			if nextTitle == "" {
				nextTitle = fmt.Sprintf("Page %d", currentIdx+2)
			}
			nextBtn = h.Button(a.Attrs(
				a.Class("flex items-center gap-4 py-4 px-5 bg-zinc-900 border border-zinc-800 rounded-xl cursor-pointer transition-all duration-200 flex-1 max-w-full md:max-w-[280px] text-right md:ml-auto hover:border-accent hover:bg-accent-dim"),
				a.OnClick(gt.SendBasicMessageNoArgs(nextMsg))),
				h.Span(a.Attrs(a.Class("flex flex-col gap-1 min-w-0")),
					h.Span(a.Attrs(a.Class("text-xs text-zinc-500 uppercase tracking-widest")), h.Text("Next")),
					h.Span(a.Attrs(a.Class("text-sm font-medium text-zinc-100 truncate")), h.Text(nextTitle)),
				),
				h.Span(a.Attrs(a.Class("text-xl text-accent shrink-0")), h.Text("\u2192")),
			)
		} else {
			nextBtn = h.Div(a.Attrs(a.Class("flex items-center gap-4 py-4 px-5 bg-zinc-900 border border-zinc-800 rounded-xl flex-1 max-w-full md:max-w-[280px] text-right md:ml-auto opacity-30 cursor-default")))
		}
	}

	// Page indicator
	pageIndicator := ""
	if mod != nil {
		pageIndicator = fmt.Sprintf("Page %d of %d", currentIdx+1, len(mod.Pages))
	}

	// Add presenter badge if presenting
	var modeIndicator h.Element
	if isPresenting {
		modeIndicator = h.Span(a.Attrs(a.Class("text-sm py-1 px-3 rounded-full bg-accent text-zinc-950 font-semibold")), h.Text("\U0001F4E1 Presenting"))
	} else if isFollowing {
		modeIndicator = h.Span(a.Attrs(a.Class("text-sm py-1 px-3 rounded-full bg-blue-500 text-white")), h.Text("\U0001F441 Following"))
	} else {
		modeIndicator = h.Span(a.Attrs())
	}

	return h.Article(a.Attrs(a.Class("max-w-full")),
		h.Div(a.Attrs(a.Class("mb-10 pb-6 border-b border-zinc-800")),
			h.Div(a.Attrs(a.Class("flex justify-between items-center mb-2")),
				h.Span(a.Attrs(a.Class("block text-xs text-accent font-medium tracking-widest uppercase")), h.Text(pageIndicator)),
				modeIndicator,
			),
			h.H1(a.Attrs(a.Class("text-3xl md:text-4xl font-bold tracking-tight leading-tight text-zinc-50")), h.Text(page.Meta.Title)),
		),
		h.Div(a.Attrs(a.Class("prose prose-invert prose-zinc max-w-none content-prose")),
			children...,
		),
		h.Nav(a.Attrs(a.Class("flex flex-col gap-4 mt-12 pt-8 border-t border-zinc-800 md:flex-row md:justify-between")),
			prevBtn,
			nextBtn,
		),
	)
}

// splitHTMLAtBlockMarkers splits HTML at block placeholder markers
func splitHTMLAtBlockMarkers(html []byte) [][]byte {
	indices := blockMarkerPattern.FindAllIndex(html, -1)
	if len(indices) == 0 {
		return [][]byte{html}
	}

	segments := make([][]byte, 0, len(indices)+1)
	start := 0
	for _, idx := range indices {
		segments = append(segments, html[start:idx[0]])
		start = idx[1]
	}
	segments = append(segments, html[start:])
	return segments
}

// renderBlock dispatches to the appropriate block renderer
func renderBlock(block Block, index int, quizStates map[int]*QuizState, activeHotspot string, isPresenting bool, mdl *Model) h.Element {
	switch b := block.(type) {
	case *QuizBlock:
		// Build poll info for this quiz
		pollInfo := getPollInfoForQuiz(b.ID, index, isPresenting, mdl)
		return renderQuiz(b, index, quizStates, pollInfo)
	case *MermaidBlock:
		return renderMermaid(b)
	case *MathBlock:
		return renderMath(b)
	case *CalloutBlock:
		return renderCallout(b)
	case *AnnotatedImageBlock:
		return renderAnnotatedImage(b, activeHotspot)
	case *TerminalReplayBlock:
		return renderTerminalReplay(b)
	case *ExerciseBlock:
		return renderExercise(b)
	default:
		return h.Div(a.Attrs()) // Empty placeholder for unknown blocks
	}
}

// getPollInfoForQuiz returns poll state for a specific quiz
func getPollInfoForQuiz(quizID string, blockIndex int, isPresenting bool, mdl *Model) *PollInfo {
	// Handle nil model (legacy path)
	if mdl == nil {
		return &PollInfo{
			IsActive:     false,
			IsInstructor: false,
		}
	}

	// Get the relevant live session (presenting or following)
	var session *LiveSession
	if isPresenting {
		session = mdl.getPresentingSession()
	} else if mdl.FollowingLive {
		session = mdl.getFollowingSession()
	}

	// No active session or no poll
	if session == nil {
		return &PollInfo{
			IsActive:     false,
			IsInstructor: isPresenting,
		}
	}

	session.RLock()
	defer session.RUnlock()

	if session.Poll == nil || !session.Poll.Active {
		return &PollInfo{
			IsActive:     false,
			IsInstructor: isPresenting,
		}
	}

	poll := session.Poll
	isThisQuiz := poll.QuizID == quizID && poll.BlockIndex == blockIndex

	// Calculate total votes
	totalVotes := 0
	for _, count := range poll.Responses {
		totalVotes += count
	}

	// Copy responses map
	responses := make(map[int]int)
	for k, v := range poll.Responses {
		responses[k] = v
	}

	return &PollInfo{
		IsActive:     poll.Active,
		IsThisQuiz:   isThisQuiz,
		IsInstructor: isPresenting,
		HasResponded: poll.Responded[mdl.SessionID],
		IsClosed:     poll.Closed,
		ShowResults:  poll.ShowResults,
		Responses:    responses,
		TotalVotes:   totalVotes,
	}
}

// renderEmptyState renders when no content is available
func renderEmptyState() h.Element {
	return h.Div(a.Attrs(a.Class("text-center py-16 px-8")),
		h.H2(a.Attrs(a.Class("mb-3 text-2xl font-semibold text-zinc-50")), h.Text("No Content Available")),
		h.P(a.Attrs(a.Class("text-zinc-500 max-w-md mx-auto")), h.Text("No modules have been loaded. Please add content to the content/ directory.")),
	)
}

// renderError renders an error message
func renderError(err error) h.Element {
	return h.Div(a.Attrs(a.Class("bg-red-500/10 border border-red-500 rounded-xl p-8")),
		h.H2(a.Attrs(a.Class("text-red-400 mb-2 text-xl font-semibold")), h.Text("Error")),
		h.P(a.Attrs(a.Class("text-zinc-400")), h.Text(err.Error())),
	)
}


// ============================================================================
// Auth Flow Rendering
// ============================================================================

// renderAuthFlow renders the appropriate auth screen based on AuthStage
func renderAuthFlow(m *Model) h.Element {
	switch m.AuthStage {
	case AuthStageEnterCode:
		return renderCohortCodeEntry(m)
	case AuthStageEnterEmail:
		return renderEmailEntry(m)
	case AuthStageError:
		return renderAuthError(m)
	default:
		return renderCohortCodeEntry(m)
	}
}

// Input styling shared across auth forms
const inputClasses = "w-full py-3 px-4 border border-zinc-700 rounded-lg bg-zinc-950 text-zinc-50 text-base transition-all duration-150 focus:outline-none focus:border-accent focus:ring-2 focus:ring-accent/20 placeholder:text-zinc-600"

// renderCohortCodeEntry renders the cohort code entry screen
func renderCohortCodeEntry(m *Model) h.Element {
	return h.Div(a.Attrs(a.Class("w-full max-w-md")),
		h.Div(a.Attrs(a.Class("bg-zinc-900 border border-zinc-800 rounded-2xl p-10 shadow-2xl")),
			h.Div(a.Attrs(a.Class("text-center mb-8")),
				h.H1(a.Attrs(a.Class("text-2xl font-bold mb-2 text-zinc-50")), h.Text("Welcome")),
				h.P(a.Attrs(a.Class("text-zinc-400")), h.Text("Enter your cohort code to begin")),
			),
			h.Form(a.Attrs(
				a.Id("cohort-code-form"),
				a.Class("flex flex-col gap-5"),
				a.OnSubmit(`event.preventDefault(); gotea.updateFormState({message:"SUBMIT_COHORT_CODE"}, "cohort-code-form")`)),
				h.Div(a.Attrs(a.Class("flex flex-col gap-2")),
					h.Label(a.Attrs(a.For("code"), a.Class("text-sm font-medium text-zinc-200")), h.Text("Cohort Code")),
					h.Input(a.Attrs(
						a.Type("text"),
						a.Id("code"),
						a.Name("code"),
						a.Placeholder("xxx-xxxx-xxx"),
						a.Class(inputClasses+" font-mono text-lg tracking-widest text-center"),
						a.Autofocus(true),
					)),
				),
				h.Button(a.Attrs(
					a.Type("submit"),
					a.Class("w-full py-3.5 px-6 bg-accent text-zinc-950 font-semibold rounded-lg text-base transition-all duration-150 hover:bg-accent-hover cursor-pointer border-none")),
					h.Text("Continue"),
				),
			),
		),
	)
}

// renderEmailEntry renders the email/name entry screen
func renderEmailEntry(m *Model) h.Element {
	cohortName := ""
	if m.PendingCohort != nil {
		cohortName = m.PendingCohort.Name
	}

	return h.Div(a.Attrs(a.Class("w-full max-w-md")),
		h.Div(a.Attrs(a.Class("bg-zinc-900 border border-zinc-800 rounded-2xl p-10 shadow-2xl")),
			h.Div(a.Attrs(a.Class("text-center mb-8")),
				h.H1(a.Attrs(a.Class("text-2xl font-bold mb-2 text-zinc-50")), h.Text("Join "+cohortName)),
				h.P(a.Attrs(a.Class("text-zinc-400")), h.Text("Enter your details to continue")),
			),
			h.Form(a.Attrs(
				a.Id("login-form"),
				a.Class("flex flex-col gap-5"),
				a.OnSubmit(`event.preventDefault(); gotea.updateFormState({message:"SUBMIT_LOGIN"}, "login-form")`)),
				h.Div(a.Attrs(a.Class("flex flex-col gap-2")),
					h.Label(a.Attrs(a.For("email"), a.Class("text-sm font-medium text-zinc-200")), h.Text("Email")),
					h.Input(a.Attrs(
						a.Type("email"),
						a.Id("email"),
						a.Name("email"),
						a.Placeholder("you@example.com"),
						a.Class(inputClasses),
						a.Required(true),
						a.Autofocus(true),
					)),
				),
				h.Div(a.Attrs(a.Class("flex flex-col gap-2")),
					h.Label(a.Attrs(a.For("name"), a.Class("text-sm font-medium text-zinc-200")), h.Text("Your Name")),
					h.Input(a.Attrs(
						a.Type("text"),
						a.Id("name"),
						a.Name("name"),
						a.Placeholder("Jane Smith"),
						a.Class(inputClasses),
						a.Required(true),
					)),
				),
				h.Button(a.Attrs(
					a.Type("submit"),
					a.Class("w-full py-3.5 px-6 bg-accent text-zinc-950 font-semibold rounded-lg text-base transition-all duration-150 hover:bg-accent-hover cursor-pointer border-none")),
					h.Text("Join Course"),
				),
			),
			h.Button(a.Attrs(
				a.Class("bg-transparent text-zinc-400 py-2 mt-4 hover:text-accent cursor-pointer border-none text-sm"),
				a.OnClick(gt.SendBasicMessageNoArgs("LOGOUT"))),
				h.Text("\u2190 Back"),
			),
		),
	)
}

// renderAuthError renders an auth error screen with retry option
func renderAuthError(m *Model) h.Element {
	return h.Div(a.Attrs(a.Class("w-full max-w-md")),
		h.Div(a.Attrs(a.Class("bg-zinc-900 border border-red-500 rounded-2xl p-10 shadow-2xl")),
			h.Div(a.Attrs(a.Class("text-center mb-8")),
				h.H1(a.Attrs(a.Class("text-2xl font-bold mb-2 text-zinc-50")), h.Text("Oops")),
			),
			h.P(a.Attrs(a.Class("text-red-400 text-base mb-6 leading-relaxed")), h.Text(m.AuthError)),
			h.Button(a.Attrs(
				a.Class("w-full py-3.5 px-6 bg-accent text-zinc-950 font-semibold rounded-lg text-base transition-all duration-150 hover:bg-accent-hover cursor-pointer border-none"),
				a.OnClick(gt.SendBasicMessageNoArgs("LOGOUT"))),
				h.Text("Try Again"),
			),
		),
	)
}

// ============================================================================
// Admin Rendering
// ============================================================================

// renderAdmin renders the admin dashboard
func (m *Model) renderAdmin(s gt.State) []byte {
	mdl := s.(*Model)

	// Must be owner to access admin
	if !mdl.IsOwner {
		return h.Div(a.Attrs(a.Class("bg-red-500/10 border border-red-500 rounded-xl p-8")),
			h.H2(a.Attrs(a.Class("text-red-400 mb-2 text-xl font-semibold")), h.Text("Access Denied")),
			h.P(a.Attrs(a.Class("text-zinc-400")), h.Text("You do not have permission to access this area.")),
		).Bytes()
	}

	// Load all cohorts
	cohorts, err := GetAllCohorts()
	if err != nil {
		LogDBError("GetAllCohorts", err)
		cohorts = []*Cohort{}
	}

	// Render cohort list
	cohortCards := make([]h.Element, 0, len(cohorts))
	for _, cohort := range cohorts {
		count, _ := GetCohortStudentCount(cohort.ID)
		expired := cohort.IsExpired()

		cardClass := "bg-zinc-900 border border-zinc-800 rounded-xl p-5"
		if expired {
			cardClass += " opacity-60"
		}

		statusBadge := h.Span(a.Attrs(a.Class("text-xs py-1 px-3 rounded-full font-medium bg-accent/15 text-accent")), h.Text("Active"))
		if expired {
			statusBadge = h.Span(a.Attrs(a.Class("text-xs py-1 px-3 rounded-full font-medium bg-red-500/15 text-red-400")), h.Text("Expired"))
		}

		cohortCards = append(cohortCards, h.Div(a.Attrs(a.Class(cardClass)),
			h.Div(a.Attrs(a.Class("flex justify-between items-center mb-3")),
				h.H3(a.Attrs(a.Class("text-lg font-semibold text-zinc-50")), h.Text(cohort.Name)),
				statusBadge,
			),
			h.Div(a.Attrs(a.Class("mb-3")),
				h.Code(a.Attrs(a.Class("font-mono text-base py-2 px-3 bg-zinc-950 border border-zinc-800 rounded-md text-accent inline-block")), h.Text(cohort.Code)),
			),
			h.Div(a.Attrs(a.Class("flex gap-6 mb-4 text-sm text-zinc-400")),
				h.Span(a.Attrs(), h.Text(fmt.Sprintf("Students: %d/%d", count, cohort.MaxStudents))),
				h.Span(a.Attrs(), h.Text(fmt.Sprintf("Expires: %s", cohort.ExpiresAt.Format("Jan 2, 2006")))),
			),
			h.Div(a.Attrs(a.Class("flex gap-2")),
				h.Button(a.Attrs(
					a.Class("py-1.5 px-3 text-sm rounded-md bg-zinc-800 text-zinc-300 font-medium cursor-pointer transition-all duration-150 hover:bg-zinc-700 border-none"),
					a.OnClick(fmt.Sprintf("navigator.clipboard.writeText('%s/join/%s'); alert('Link copied!');", getBaseURL(), cohort.Code))),
					h.Text("Copy Link"),
				),
				h.Button(a.Attrs(
					a.Class("py-1.5 px-3 text-sm rounded-md bg-red-500 text-white font-medium cursor-pointer transition-all duration-150 hover:bg-red-600 border-none"),
					a.OnClick(gt.SendBasicMessage("DELETE_COHORT", cohort.ID.String()))),
					h.Text("Delete"),
				),
			),
		))
	}

	// Empty state if no cohorts
	if len(cohortCards) == 0 {
		cohortCards = append(cohortCards, h.Div(a.Attrs(a.Class("text-center py-16 px-8")),
			h.P(a.Attrs(a.Class("text-zinc-500")), h.Text("No cohorts yet. Create one to get started.")),
		))
	}

	return h.Div(a.Attrs(a.Class("max-w-[900px] mx-auto py-4")),
		h.Div(a.Attrs(a.Class("flex justify-between items-center mb-10 flex-wrap gap-4")),
			h.H1(a.Attrs(a.Class("text-3xl font-bold tracking-tight text-zinc-50")), h.Text("Cohort Management")),
			h.A(a.Attrs(
				a.Href("/"),
				a.Class("py-2.5 px-5 bg-zinc-800 text-zinc-300 font-semibold rounded-lg transition-all duration-150 hover:bg-zinc-700 no-underline")),
				h.Text("\u2190 Back to Course"),
			),
		),
		h.Div(a.Attrs(a.Class("mb-10")),
			h.H2(a.Attrs(a.Class("text-xl font-semibold mb-5 text-zinc-50")), h.Text("Create New Cohort")),
			h.Form(a.Attrs(
				a.Id("create-cohort-form"),
				a.Class("bg-zinc-900 border border-zinc-800 rounded-xl p-6"),
				a.OnSubmit(`event.preventDefault(); gotea.updateFormState({message:"CREATE_COHORT"}, "create-cohort-form")`)),
				h.Div(a.Attrs(a.Class("grid grid-cols-1 gap-4 mb-5 sm:grid-cols-[2fr_1fr_1fr]")),
					h.Div(a.Attrs(a.Class("flex flex-col gap-2")),
						h.Label(a.Attrs(a.For("name"), a.Class("text-sm font-medium text-zinc-200")), h.Text("Cohort Name")),
						h.Input(a.Attrs(
							a.Type("text"),
							a.Id("name"),
							a.Name("name"),
							a.Placeholder("e.g., Acme Corp March 2024"),
							a.Class(inputClasses),
							a.Required(true),
						)),
					),
					h.Div(a.Attrs(a.Class("flex flex-col gap-2")),
						h.Label(a.Attrs(a.For("expiresAt"), a.Class("text-sm font-medium text-zinc-200")), h.Text("Expires")),
						h.Input(a.Attrs(
							a.Type("date"),
							a.Id("expiresAt"),
							a.Name("expiresAt"),
							a.Class(inputClasses),
							a.Required(true),
						)),
					),
					h.Div(a.Attrs(a.Class("flex flex-col gap-2")),
						h.Label(a.Attrs(a.For("maxStudents"), a.Class("text-sm font-medium text-zinc-200")), h.Text("Max Students")),
						h.Input(a.Attrs(
							a.Type("number"),
							a.Id("maxStudents"),
							a.Name("maxStudents"),
							a.Value("25"),
							a.Min("1"),
							a.Max("1000"),
							a.Class(inputClasses),
						)),
					),
				),
				h.Button(a.Attrs(
					a.Type("submit"),
					a.Class("py-3 px-6 bg-accent text-zinc-950 font-semibold rounded-lg transition-all duration-150 hover:bg-accent-hover cursor-pointer border-none")),
					h.Text("Create Cohort"),
				),
			),
		),
		h.Div(a.Attrs(a.Class("mb-10")),
			h.H2(a.Attrs(a.Class("text-xl font-semibold mb-5 text-zinc-50")), h.Text("Existing Cohorts")),
			h.Div(a.Attrs(a.Class("flex flex-col gap-4")),
				cohortCards...,
			),
		),
	).Bytes()
}

// getBaseURL returns the base URL for the app (placeholder)
func getBaseURL() string {
	// In production, this would come from config
	return "http://localhost:8080"
}
