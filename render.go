package main

import (
	"encoding/json"
	"fmt"
	stdhtml "html"
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
			h.Body(a.Attrs(a.Class("bg-stone-50 text-stone-900 font-sans antialiased min-h-screen flex items-center justify-center p-8")),
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

	// Check if agent sidebar should be rendered
	agentState := m.getAgentStateForRender()
	isFollowing := m.FollowingLive

	// Shared <head> element
	head := h.Head(a.Attrs(),
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
			a.Href("https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&family=JetBrains+Mono:wght@400;500;600&display=swap"))),
		// Stylesheets
		h.Link(a.Attrs(
			a.Rel("stylesheet"),
			a.Href("static/css/main.css"))),
	)

	// Full-screen agent workspace mode
	if agentState != nil && agentState.WorkspaceOpen {
		return h.Html(a.Attrs(a.Lang("en")),
			head,
			h.Body(a.Attrs(
				a.Class(fmt.Sprintf("bg-zinc-950 text-zinc-200 font-sans antialiased h-screen overflow-hidden %s", fontClass))),
				renderAgentWorkspace(m, agentState, isFollowing),
				h.Script(a.Attrs(a.Src("static/js/main.js"))),
			))
	}

	// Normal course layout
	return h.Html(a.Attrs(a.Lang("en")),
		head,
		h.Body(a.Attrs(
			a.Class(fmt.Sprintf("bg-stone-50 text-stone-900 font-sans antialiased min-h-screen %s", fontClass))),
			h.Div(a.Attrs(a.Class("flex flex-col min-h-screen md:flex-row")),
				renderSidebar(m),
				h.Main(a.Attrs(a.Class("flex-1 py-6 md:py-10 overflow-y-auto")),
					h.Div(a.Attrs(a.Id("view"), a.Class("relative w-[660px] shrink-0 box-content px-6 md:px-12")),
						h.UnsafeRaw(string(m.RenderRoute(m))),
						renderAnnotationLayer(m)),
				),
			),
			h.Script(a.Attrs(a.Src("static/js/main.js"))),
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
		a.Class("block py-3 px-4 bg-stone-100 border border-stone-200 rounded-lg text-stone-500 no-underline text-sm font-medium text-center transition-all duration-150 hover:bg-stone-200 hover:border-accent hover:text-accent")),
		h.Text(viewToggleLabel),
	)

	// Presenter/follower controls
	var sessionControl h.Element
	if m.IsPresenting {
		sessionControl = h.Div(a.Attrs(a.Class("flex items-center justify-between py-2 px-3 bg-accent-dim rounded-lg border border-accent")),
			h.Span(a.Attrs(a.Class("text-sm font-medium text-accent")), h.Text("Presenting")),
			h.Button(a.Attrs(
				a.Class("py-1.5 px-3 text-sm rounded-md bg-stone-200 text-stone-700 font-medium cursor-pointer transition-all duration-150 hover:bg-stone-300 border-none"),
				a.OnClick(gt.SendBasicMessageNoArgs("STOP_PRESENTING"))),
				h.Text("Stop"),
			),
		)
	} else if m.FollowingLive {
		sessionControl = h.Div(a.Attrs(a.Class("flex items-center justify-between py-2 px-3 bg-blue-50 rounded-lg border border-blue-300")),
			h.Span(a.Attrs(a.Class("text-sm font-medium text-blue-600")), h.Text("Following")),
			h.Button(a.Attrs(
				a.Class("py-1.5 px-3 text-sm rounded-md bg-stone-200 text-stone-700 font-medium cursor-pointer transition-all duration-150 hover:bg-stone-300 border-none"),
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
			sessionControl = h.Div(a.Attrs(a.Class("flex items-center justify-between py-2 px-3 bg-amber-50 rounded-lg border border-amber-300")),
				h.Span(a.Attrs(a.Class("text-sm font-medium text-amber-600")), h.Text("Live Session")),
				h.Button(a.Attrs(
					a.Class("py-1.5 px-3 text-sm rounded-md bg-accent text-white font-medium cursor-pointer transition-all duration-150 hover:bg-accent-hover border-none"),
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
			a.Class("block py-3 px-4 bg-accent-dim border border-accent rounded-lg text-accent no-underline text-sm font-medium text-center transition-all duration-150 hover:bg-accent/15")),
			h.Text("Admin"),
		)
	} else {
		adminLink = h.Span(a.Attrs())
	}

	// Logout button
	logoutBtn := h.Button(a.Attrs(
		a.Class("py-1.5 px-3 text-sm rounded-md bg-stone-200 text-stone-600 font-medium cursor-pointer transition-all duration-150 hover:bg-stone-300 border-none mt-2"),
		a.OnClick(gt.SendBasicMessageNoArgs("LOGOUT"))),
		h.Text("Logout"),
	)

	return h.Div(a.Attrs(a.Class("mt-auto pt-4 border-t border-stone-200 flex flex-row flex-wrap gap-3 md:flex-col")),
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
		return h.Div(a.Attrs(a.Class("py-2 px-3 bg-stone-100 rounded-lg border border-stone-200 text-stone-400 text-sm")),
			h.Text("No active cohorts"),
		)
	}

	if len(activeCohorts) == 1 {
		c := activeCohorts[0]
		return h.Div(a.Attrs(a.Class("flex items-center justify-between py-2 px-3 bg-stone-100 rounded-lg border border-stone-200")),
			h.Span(a.Attrs(a.Class("text-sm text-stone-500 truncate mr-2")), h.Text(c.Name)),
			h.Button(a.Attrs(
				a.Class("py-1.5 px-3 text-sm rounded-md bg-stone-200 text-stone-700 font-medium cursor-pointer transition-all duration-150 hover:bg-stone-300 border-none shrink-0"),
				a.OnClick(gt.SendBasicMessage("START_PRESENTING", c.ID.String()))),
				h.Text("Present"),
			),
		)
	}

	// Multiple cohorts — show a button per cohort
	buttons := make([]h.Element, len(activeCohorts))
	for i, c := range activeCohorts {
		buttons[i] = h.Button(a.Attrs(
			a.Class("w-full py-2 px-3 text-sm rounded-md bg-stone-200 text-stone-700 font-medium cursor-pointer transition-all duration-150 hover:bg-stone-300 border-none text-left"),
			a.OnClick(gt.SendBasicMessage("START_PRESENTING", c.ID.String()))),
			h.Text(fmt.Sprintf("Present to %s", c.Name)),
		)
	}
	return h.Div(a.Attrs(a.Class("flex flex-col gap-2 py-2 px-3 bg-stone-100 rounded-lg border border-stone-200")),
		append([]h.Element{
			h.Span(a.Attrs(a.Class("text-xs text-stone-400 font-medium uppercase tracking-wider")), h.Text("Present to")),
		}, buttons...)...,
	)
}

// renderSidebar renders the module/page navigation sidebar
func renderSidebar(m *Model) h.Element {
	mod := m.currentModule()
	currentPageIdx := m.CurrentPage

	if mod == nil {
		return h.Aside(a.Attrs(a.Class("w-full relative border-b border-stone-200 p-6 overflow-y-auto flex flex-col gap-6 max-h-[50vh] bg-white/90 backdrop-blur-sm md:w-72 md:h-screen md:sticky md:top-0 md:border-r md:border-b-0 md:max-h-none md:shrink-0")),
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

	// Module index for sub-numbering (e.g. 6.1, 6.2)
	modIdx := globalCourse.ModuleIndex(mod.ID)

	var pages []h.Element
	for i, page := range mod.Pages {
		isCurrent := i == currentPageIdx

		title := page.Meta.Title
		if title == "" {
			title = fmt.Sprintf("Page %d", i+1)
		}

		pageNum := fmt.Sprintf("%d.%d", modIdx, i+1)
		viewed := progress != nil && progress.PagesViewed[i]

		// If page has sections and is the current page, render as group with sub-items
		if page.Sections != nil && isCurrent {
			// Page title as group header (non-clickable)
			pageNumClass := "flex items-center justify-center min-w-[1.75rem] h-6 px-1 bg-accent text-white rounded-full text-xs font-semibold shrink-0"
			pages = append(pages, h.Div(a.Attrs(a.Class("flex items-center gap-3 py-2.5 px-3 rounded-lg text-accent text-sm font-medium border border-transparent")),
				h.Span(a.Attrs(a.Class(pageNumClass)), h.Text(pageNum)),
				h.Span(a.Attrs(a.Class("flex-1 truncate")), h.Text(title)),
			))

			// Section sub-items
			for si, section := range page.Sections {
				isCurrentSlide := si == m.CurrentSlide
				slideClasses := "flex items-center gap-2 py-1.5 px-3 pl-12 rounded-md text-stone-400 no-underline text-xs transition-all duration-150 border border-transparent external"
				if isCurrentSlide {
					slideClasses = "flex items-center gap-2 py-1.5 px-3 pl-12 rounded-md text-accent no-underline text-xs transition-all duration-150 border border-accent/20 bg-accent-dim external"
				}

				// Slide viewed indicator
				slideViewed := progress != nil && progress.SlidesViewed != nil && progress.SlidesViewed[i] != nil && progress.SlidesViewed[i][si]
				if slideViewed && !isCurrentSlide {
					slideClasses = "flex items-center gap-2 py-1.5 px-3 pl-12 rounded-md text-stone-700 no-underline text-xs transition-all duration-150 border border-transparent external"
				}

				if navDisabled {
					slideClasses += " opacity-60 cursor-not-allowed"
				} else {
					slideClasses += " hover:bg-stone-100 hover:text-stone-800 cursor-pointer"
				}

				var slideClickHandler string
				if !navDisabled {
					slideClickHandler = gt.SendBasicMessage("NAV_SLIDE", si) + "; return false;"
				} else {
					slideClickHandler = "return false;"
				}

				var slideStatus h.Element
				if !m.IsOwner && slideViewed {
					slideStatus = h.Span(a.Attrs(a.Class("text-xs w-4 text-center shrink-0 text-accent")), h.Text("\u2713"))
				} else {
					slideStatus = h.Span(a.Attrs(a.Class("text-xs w-4 text-center shrink-0")), h.Text(""))
				}

				pages = append(pages, h.A(a.Attrs(
					a.Href("#"),
					a.Class(slideClasses),
					a.OnClick(slideClickHandler)),
					h.Span(a.Attrs(a.Class("flex-1 truncate")), h.Text(section.Title)),
					slideStatus,
				))
			}
			continue
		}

		// Regular page item (no sections, or not current page)
		classes := "flex items-center gap-3 py-2.5 px-3 rounded-lg text-stone-500 no-underline text-sm transition-all duration-150 border border-transparent external"
		if isCurrent {
			classes = "flex items-center gap-3 py-2.5 px-3 rounded-lg text-accent no-underline text-sm transition-all duration-150 border border-accent/30 bg-accent-dim external"
		}
		if navDisabled {
			classes += " opacity-60 cursor-not-allowed"
		} else {
			classes += " hover:bg-stone-100 hover:text-stone-800"
		}
		if viewed && !isCurrent {
			classes += " text-stone-700"
		}

		pageNumClass := "flex items-center justify-center min-w-[1.75rem] h-6 px-1 bg-stone-200 text-stone-500 rounded-full text-xs font-semibold shrink-0"
		if isCurrent {
			pageNumClass = "flex items-center justify-center min-w-[1.75rem] h-6 px-1 bg-accent text-white rounded-full text-xs font-semibold shrink-0"
		} else if viewed {
			pageNumClass = "flex items-center justify-center min-w-[1.75rem] h-6 px-1 bg-stone-300 text-stone-600 rounded-full text-xs font-semibold shrink-0"
		}

		var statusIcon h.Element
		if m.IsOwner {
			statusIcon = h.Span(a.Attrs(a.Class("text-xs w-5 text-center shrink-0")), h.Text(""))
		} else if m.isPageComplete(mod.ID, i) {
			statusIcon = h.Span(a.Attrs(a.Class("text-xs w-5 text-center shrink-0 text-accent")), h.Text("\u2713"))
		} else if m.pageHasIncompleteQuiz(mod.ID, i) {
			statusIcon = h.Span(a.Attrs(a.Class("text-xs w-5 text-center shrink-0 text-amber-500")), h.Text("\u25CB"))
		} else if viewed {
			statusIcon = h.Span(a.Attrs(a.Class("text-xs w-5 text-center shrink-0 text-accent")), h.Text("\u2713"))
		} else {
			statusIcon = h.Span(a.Attrs(a.Class("text-xs w-5 text-center shrink-0")), h.Text(""))
		}

		var clickHandler string
		if !navDisabled {
			clickHandler = gt.SendBasicMessage("NAV_PAGE", i) + "; return false;"
		} else {
			clickHandler = "return false;"
		}

		pages = append(pages, h.A(a.Attrs(
			a.Href("#"),
			a.Class(classes),
			a.OnClick(clickHandler)),
			h.Span(a.Attrs(a.Class(pageNumClass)), h.Text(pageNum)),
			h.Span(a.Attrs(a.Class("flex-1 truncate")), h.Text(title)),
			statusIcon,
		))
	}

	// Module completion status (students only)
	var moduleStatus h.Element
	if m.IsOwner {
		// Owner just sees page count
		moduleStatus = h.Div(a.Attrs(a.Class("flex items-center gap-2 text-xs text-stone-400 py-2 px-3 bg-stone-100 rounded-md")),
			h.Span(a.Attrs(),
				h.Text(fmt.Sprintf("%d pages", len(mod.Pages)))),
		)
	} else if progress != nil && progress.CompletedAt != nil {
		moduleStatus = h.Div(a.Attrs(a.Class("flex items-center gap-2 text-xs text-accent py-2 px-3 bg-accent-dim rounded-md")),
			h.Span(a.Attrs(a.Class("text-sm")), h.Text("\u2713")),
			h.Span(a.Attrs(), h.Text("Module Complete")),
		)
	} else if m.moduleNeedsQuizzes(mod.ID) {
		moduleStatus = h.Div(a.Attrs(a.Class("flex items-center gap-2 text-xs text-amber-600 py-2 px-3 bg-amber-50 rounded-md")),
			h.Span(a.Attrs(a.Class("text-sm")), h.Text("\u25CB")),
			h.Span(a.Attrs(), h.Text("Quizzes needed")),
		)
	} else {
		pagesViewed := 0
		if progress != nil {
			pagesViewed = len(progress.PagesViewed)
		}
		moduleStatus = h.Div(a.Attrs(a.Class("flex items-center gap-2 text-xs text-stone-400 py-2 px-3 bg-stone-100 rounded-md")),
			h.Span(a.Attrs(),
				h.Text(fmt.Sprintf("%d of %d pages", pagesViewed, len(mod.Pages)))),
		)
	}

	// Add indicator when following presenter
	var followIndicator h.Element
	if navDisabled {
		followIndicator = h.Div(a.Attrs(a.Class("text-xs text-blue-600 p-2 mt-2 bg-blue-50 rounded-md text-center")),
			h.Text("Navigation controlled by presenter"),
		)
	} else {
		followIndicator = h.Span(a.Attrs())
	}

	return h.Aside(a.Attrs(a.Class("w-full relative border-b border-stone-200 p-6 overflow-y-auto flex flex-col gap-6 max-h-[50vh] bg-white/90 backdrop-blur-sm md:w-72 md:h-screen md:sticky md:top-0 md:border-r md:border-b-0 md:max-h-none md:shrink-0")),
		h.Div(a.Attrs(a.Class("flex flex-col gap-3")),
			h.H2(a.Attrs(a.Class("text-lg font-semibold text-stone-900 tracking-tight")), h.Text(mod.Meta.Title)),
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
		cardClasses := "bg-white border border-stone-200 rounded-2xl p-6 transition-all duration-200 relative cursor-pointer shadow-sm hover:border-accent hover:bg-accent-dim hover:shadow-md"
		if !available {
			cardClasses = "bg-white border border-stone-200 rounded-2xl p-6 transition-all duration-200 relative opacity-50 cursor-not-allowed shadow-sm"
		}
		if id == mdl.CurrentModule {
			cardClasses += " border-accent"
		}

		// Module number styling
		numClass := "flex items-center justify-center w-8 h-8 bg-accent text-white rounded-full font-bold text-sm"
		if !available {
			numClass = "flex items-center justify-center w-8 h-8 bg-stone-300 text-stone-500 rounded-full font-bold text-sm"
		}

		// Status badge and progress bar — only for students
		var statusEl h.Element
		var progressBar h.Element
		if mdl.IsOwner {
			// Owner sees page count only, no progress state
			statusEl = h.Span(a.Attrs(a.Class("text-xs py-1 px-2.5 rounded-full bg-stone-100 text-stone-400 font-medium")),
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
				statusEl = h.Span(a.Attrs(a.Class("text-xs py-1 px-2.5 rounded-full bg-accent-dim text-accent font-medium")), h.Text("Completed"))
			} else if mdl.moduleNeedsQuizzes(id) {
				statusEl = h.Span(a.Attrs(a.Class("text-xs py-1 px-2.5 rounded-full bg-amber-50 text-amber-600 font-medium")), h.Text("Quizzes needed"))
			} else if progress != nil && progress.Started {
				fillClass = "h-full bg-blue-500 rounded-full transition-all duration-300"
				statusEl = h.Span(a.Attrs(a.Class("text-xs py-1 px-2.5 rounded-full bg-blue-50 text-blue-600 font-medium")),
					h.Text(fmt.Sprintf("%d/%d pages", pagesViewed, len(mod.Pages))))
			} else if !available {
				statusEl = h.Span(a.Attrs(a.Class("text-xs py-1 px-2.5 rounded-full bg-stone-100 text-stone-400 font-medium")), h.Text("Locked"))
			} else {
				statusEl = h.Span(a.Attrs(a.Class("text-xs py-1 px-2.5 rounded-full bg-stone-100 text-stone-400 font-medium")), h.Text("Not started"))
			}

			progressBar = h.Div(a.Attrs(a.Class("h-1 bg-stone-200 rounded-full overflow-hidden mb-3")),
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
			h.H3(a.Attrs(a.Class("text-xl font-semibold mb-2 text-stone-900 tracking-tight")), h.Text(mod.Meta.Title)),
			h.P(a.Attrs(a.Class("text-stone-500 text-sm mb-4 leading-relaxed")), h.Text(mod.Meta.Description)),
			progressBar,
			h.Div(a.Attrs(a.Class("flex gap-4 text-xs text-stone-400")),
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
			h.H1(a.Attrs(a.Class("text-3xl md:text-4xl font-bold tracking-tight text-stone-900")), h.Text("Course Modules")),
		)
	} else {
		header = h.Div(a.Attrs(a.Class("mb-10 flex flex-col items-start gap-4 md:flex-row md:items-center md:justify-between")),
			h.Div(a.Attrs(),
				h.H1(a.Attrs(a.Class("text-3xl md:text-4xl font-bold tracking-tight text-stone-900")), h.Text("Course Modules")),
			),
			h.Div(a.Attrs(a.Class("flex gap-8")),
				h.Div(a.Attrs(a.Class("text-center")),
					h.Span(a.Attrs(a.Class("block text-3xl font-bold text-accent tracking-tight")), h.Text(fmt.Sprintf("%d%%", overallPct))),
					h.Span(a.Attrs(a.Class("text-xs text-stone-400 uppercase tracking-widest")), h.Text("Complete")),
				),
				h.Div(a.Attrs(a.Class("text-center")),
					h.Span(a.Attrs(a.Class("block text-3xl font-bold text-accent tracking-tight")), h.Text(fmt.Sprintf("%d/%d", completedModules, totalModules))),
					h.Span(a.Attrs(a.Class("text-xs text-stone-400 uppercase tracking-widest")), h.Text("Modules")),
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
	log.Printf("renderPageContentWithNav: sessionID=%s isPresenting=%v isFollowing=%v slide=%d", mdl.SessionID, isPresenting, isFollowing, mdl.CurrentSlide)

	// Determine which HTML and blocks to render
	var narrativeHTML []byte
	var blocks []Block
	var sectionTitle string

	if page.Sections != nil && mdl.CurrentSlide >= 0 && mdl.CurrentSlide < len(page.Sections) {
		section := page.Sections[mdl.CurrentSlide]
		narrativeHTML = section.NarrativeHTML
		blocks = section.Blocks
		sectionTitle = section.Title
	} else {
		narrativeHTML = page.NarrativeHTML
		blocks = page.Blocks
	}

	// Split the narrative HTML at block markers
	segments := splitHTMLAtBlockMarkers(narrativeHTML)
	children := make([]h.Element, 0, len(segments)+len(blocks))

	for i, segment := range segments {
		if len(segment) > 0 {
			children = append(children, h.UnsafeRaw(string(segment)))
		}
		if i < len(blocks) {
			children = append(children,
				renderBlock(blocks[i], i, quizStates, activeHotspot, isPresenting, isFollowing, mdl))
		}
	}

	// Build navigation controls
	// For sectioned pages, prev/next cycle through slides first
	sectionCount := 0
	if page.Sections != nil {
		sectionCount = len(page.Sections)
	}

	hasPrevSlide := mdl.CurrentSlide > 0
	hasNextSlide := sectionCount > 0 && mdl.CurrentSlide < sectionCount-1
	hasPrevPage := currentIdx > 0
	hasNextPage := mod != nil && currentIdx < len(mod.Pages)-1
	quizzesBlocking := !mdl.IsOwner && mdl.currentPageHasUnansweredRequiredQuizzes()

	// Navigation messages are always the same - the handlers check if presenting
	prevMsg := "PREV_PAGE"
	nextMsg := "NEXT_PAGE"

	var prevBtn, nextBtn h.Element

	// When following instructor, disable navigation
	if isFollowing {
		prevBtn = h.Div(a.Attrs(a.Class("flex items-center gap-4 py-4 px-5 bg-white border border-stone-200 rounded-xl flex-1 max-w-full md:max-w-[280px] text-left opacity-30 cursor-default")))
		nextBtn = h.Div(a.Attrs(a.Class("flex items-center gap-4 py-4 px-5 bg-white border border-stone-200 rounded-xl flex-1 max-w-full md:max-w-[280px] text-right md:ml-auto opacity-30 cursor-default")))
	} else {
		// --- Previous button ---
		if hasPrevSlide {
			// Previous slide within same page
			prevTitle := page.Sections[mdl.CurrentSlide-1].Title
			prevBtn = h.Button(a.Attrs(
				a.Class("flex items-center gap-4 py-4 px-5 bg-white border border-stone-200 rounded-xl cursor-pointer transition-all duration-200 flex-1 max-w-full md:max-w-[280px] text-left shadow-sm hover:border-accent hover:bg-accent-dim hover:shadow-md"),
				a.OnClick(gt.SendBasicMessageNoArgs(prevMsg))),
				h.Span(a.Attrs(a.Class("text-xl text-accent shrink-0")), h.Text("\u2190")),
				h.Span(a.Attrs(a.Class("flex flex-col gap-1 min-w-0")),
					h.Span(a.Attrs(a.Class("text-xs text-stone-400 uppercase tracking-widest")), h.Text("Previous")),
					h.Span(a.Attrs(a.Class("text-sm font-medium text-stone-800 truncate")), h.Text(prevTitle)),
				),
			)
		} else if hasPrevPage {
			prevPage := mod.Pages[currentIdx-1]
			prevTitle := prevPage.Meta.Title
			if prevTitle == "" {
				prevTitle = fmt.Sprintf("Page %d", currentIdx)
			}
			// If previous page has sections, show last section title
			if prevPage.Sections != nil {
				prevTitle = prevPage.Sections[len(prevPage.Sections)-1].Title
			}
			prevBtn = h.Button(a.Attrs(
				a.Class("flex items-center gap-4 py-4 px-5 bg-white border border-stone-200 rounded-xl cursor-pointer transition-all duration-200 flex-1 max-w-full md:max-w-[280px] text-left shadow-sm hover:border-accent hover:bg-accent-dim hover:shadow-md"),
				a.OnClick(gt.SendBasicMessageNoArgs(prevMsg))),
				h.Span(a.Attrs(a.Class("text-xl text-accent shrink-0")), h.Text("\u2190")),
				h.Span(a.Attrs(a.Class("flex flex-col gap-1 min-w-0")),
					h.Span(a.Attrs(a.Class("text-xs text-stone-400 uppercase tracking-widest")), h.Text("Previous")),
					h.Span(a.Attrs(a.Class("text-sm font-medium text-stone-800 truncate")), h.Text(prevTitle)),
				),
			)
		} else {
			prevBtn = h.Div(a.Attrs(a.Class("flex items-center gap-4 py-4 px-5 bg-white border border-stone-200 rounded-xl flex-1 max-w-full md:max-w-[280px] text-left opacity-30 cursor-default")))
		}

		// --- Next button ---
		if hasNextSlide {
			// Next slide within same page
			nextTitle := page.Sections[mdl.CurrentSlide+1].Title
			nextBtn = h.Button(a.Attrs(
				a.Class("flex items-center gap-4 py-4 px-5 bg-white border border-stone-200 rounded-xl cursor-pointer transition-all duration-200 flex-1 max-w-full md:max-w-[280px] text-right md:ml-auto shadow-sm hover:border-accent hover:bg-accent-dim hover:shadow-md"),
				a.OnClick(gt.SendBasicMessageNoArgs(nextMsg))),
				h.Span(a.Attrs(a.Class("flex flex-col gap-1 min-w-0")),
					h.Span(a.Attrs(a.Class("text-xs text-stone-400 uppercase tracking-widest")), h.Text("Next")),
					h.Span(a.Attrs(a.Class("text-sm font-medium text-stone-800 truncate")), h.Text(nextTitle)),
				),
				h.Span(a.Attrs(a.Class("text-xl text-accent shrink-0")), h.Text("\u2192")),
			)
		} else if quizzesBlocking {
			// Quizzes blocking — show on last slide or non-sectioned page
			nextLabel := "Complete quizzes to continue"
			nextTitle := ""
			if hasNextPage {
				nextTitle = mod.Pages[currentIdx+1].Meta.Title
				if nextTitle == "" {
					nextTitle = fmt.Sprintf("Page %d", currentIdx+2)
				}
			}
			nextBtn = h.Div(a.Attrs(a.Class("flex items-center gap-4 py-4 px-5 bg-white border border-amber-400 rounded-xl flex-1 max-w-full md:max-w-[280px] text-right md:ml-auto cursor-not-allowed")),
				h.Span(a.Attrs(a.Class("flex flex-col gap-1 min-w-0")),
					h.Span(a.Attrs(a.Class("text-xs text-amber-600 uppercase tracking-widest")), h.Text(nextLabel)),
					h.Span(a.Attrs(a.Class("text-sm font-medium text-stone-500 truncate")), h.Text(nextTitle)),
				),
			)
		} else if hasNextPage {
			nextPage := mod.Pages[currentIdx+1]
			nextTitle := nextPage.Meta.Title
			if nextTitle == "" {
				nextTitle = fmt.Sprintf("Page %d", currentIdx+2)
			}
			// If next page has sections, show first section title
			if nextPage.Sections != nil {
				nextTitle = nextPage.Sections[0].Title
			}
			nextBtn = h.Button(a.Attrs(
				a.Class("flex items-center gap-4 py-4 px-5 bg-white border border-stone-200 rounded-xl cursor-pointer transition-all duration-200 flex-1 max-w-full md:max-w-[280px] text-right md:ml-auto shadow-sm hover:border-accent hover:bg-accent-dim hover:shadow-md"),
				a.OnClick(gt.SendBasicMessageNoArgs(nextMsg))),
				h.Span(a.Attrs(a.Class("flex flex-col gap-1 min-w-0")),
					h.Span(a.Attrs(a.Class("text-xs text-stone-400 uppercase tracking-widest")), h.Text("Next")),
					h.Span(a.Attrs(a.Class("text-sm font-medium text-stone-800 truncate")), h.Text(nextTitle)),
				),
				h.Span(a.Attrs(a.Class("text-xl text-accent shrink-0")), h.Text("\u2192")),
			)
		} else if mod != nil {
			// Last slide of last page — check for next module
			nextModID := globalCourse.NextModule(mod.ID)
			if nextModID != "" {
				nextMod := globalModules[nextModID]
				if nextMod != nil {
					nextModTitle := nextMod.Meta.Title
					if mdl.IsOwner || globalCourse.PrerequisitesMet(nextModID, mdl.Progress) {
						nextBtn = h.Button(a.Attrs(
							a.Class("flex items-center gap-4 py-4 px-5 bg-white border border-accent rounded-xl cursor-pointer transition-all duration-200 flex-1 max-w-full md:max-w-[320px] text-right md:ml-auto shadow-sm hover:bg-accent-dim hover:shadow-md"),
							a.OnClick(gt.SendBasicMessage("NAV_MODULE", nextModID))),
							h.Span(a.Attrs(a.Class("flex flex-col gap-1 min-w-0")),
								h.Span(a.Attrs(a.Class("text-xs text-accent uppercase tracking-widest")), h.Text("Next Module")),
								h.Span(a.Attrs(a.Class("text-sm font-medium text-stone-800 truncate")), h.Text(nextModTitle)),
							),
							h.Span(a.Attrs(a.Class("text-xl text-accent shrink-0")), h.Text("\u2192")),
						)
					} else {
						nextBtn = h.Div(a.Attrs(a.Class("flex items-center gap-4 py-4 px-5 bg-white border border-stone-200 rounded-xl flex-1 max-w-full md:max-w-[320px] text-right md:ml-auto opacity-50 cursor-not-allowed")),
							h.Span(a.Attrs(a.Class("flex flex-col gap-1 min-w-0")),
								h.Span(a.Attrs(a.Class("text-xs text-stone-400 uppercase tracking-widest")), h.Text("Next Module (Locked)")),
								h.Span(a.Attrs(a.Class("text-sm font-medium text-stone-500 truncate")), h.Text(nextModTitle)),
							),
							h.Span(a.Attrs(a.Class("text-xl text-stone-400 shrink-0")), h.Text("\U0001F512")),
						)
					}
				} else {
					nextBtn = h.Div(a.Attrs(a.Class("flex items-center gap-4 py-4 px-5 bg-white border border-stone-200 rounded-xl flex-1 max-w-full md:max-w-[280px] text-right md:ml-auto opacity-30 cursor-default")))
				}
			} else {
				nextBtn = h.Div(a.Attrs(a.Class("flex items-center gap-4 py-4 px-5 bg-white border border-accent rounded-xl flex-1 max-w-full md:max-w-[280px] text-right md:ml-auto")),
					h.Span(a.Attrs(a.Class("flex flex-col gap-1 min-w-0")),
						h.Span(a.Attrs(a.Class("text-xs text-accent uppercase tracking-widest")), h.Text("Course Complete")),
						h.Span(a.Attrs(a.Class("text-sm font-medium text-accent")), h.Text("\u2713 Well done!")),
					),
				)
			}
		} else {
			nextBtn = h.Div(a.Attrs(a.Class("flex items-center gap-4 py-4 px-5 bg-white border border-stone-200 rounded-xl flex-1 max-w-full md:max-w-[280px] text-right md:ml-auto opacity-30 cursor-default")))
		}
	}

	// Page indicator
	pageIndicator := ""
	if mod != nil {
		modIdx := globalCourse.ModuleIndex(mod.ID)
		pageIndicator = fmt.Sprintf("Section %d.%d of %d", modIdx, currentIdx+1, len(mod.Pages))
	}

	// Add presenter badge if presenting
	var modeIndicator h.Element
	if isPresenting {
		modeIndicator = h.Span(a.Attrs(a.Class("text-sm py-1 px-3 rounded-full bg-accent text-white font-semibold")), h.Text("\U0001F4E1 Presenting"))
	} else if isFollowing {
		modeIndicator = h.Span(a.Attrs(a.Class("text-sm py-1 px-3 rounded-full bg-blue-500 text-white")), h.Text("\U0001F441 Following"))
	} else {
		modeIndicator = h.Span(a.Attrs())
	}

	pageID := ""
	if mod != nil {
		pageID = fmt.Sprintf("%s:%d:%d", mod.ID, currentIdx, mdl.CurrentSlide)
	}

	// Build header: for sectioned pages show page title as breadcrumb and section title as H1
	var headerEl h.Element
	if sectionTitle != "" {
		// Slide indicator
		slideIndicator := fmt.Sprintf("Slide %d of %d", mdl.CurrentSlide+1, sectionCount)

		headerEl = h.Div(a.Attrs(a.Class("mb-10 pb-6 border-b border-stone-200")),
			h.Div(a.Attrs(a.Class("flex justify-between items-center mb-2")),
				h.Span(a.Attrs(a.Class("block text-xs text-accent font-medium tracking-widest uppercase")), h.Text(pageIndicator)),
				modeIndicator,
			),
			h.Span(a.Attrs(a.Class("block text-sm text-stone-400 mb-1")), h.Text(page.Meta.Title)),
			h.H1(a.Attrs(a.Class("text-3xl md:text-4xl font-bold tracking-tight leading-tight text-stone-900")), h.Text(sectionTitle)),
			h.Span(a.Attrs(a.Class("block text-xs text-stone-400 mt-2")), h.Text(slideIndicator)),
		)
	} else {
		headerEl = h.Div(a.Attrs(a.Class("mb-10 pb-6 border-b border-stone-200")),
			h.Div(a.Attrs(a.Class("flex justify-between items-center mb-2")),
				h.Span(a.Attrs(a.Class("block text-xs text-accent font-medium tracking-widest uppercase")), h.Text(pageIndicator)),
				modeIndicator,
			),
			h.H1(a.Attrs(a.Class("text-3xl md:text-4xl font-bold tracking-tight leading-tight text-stone-900")), h.Text(page.Meta.Title)),
		)
	}

	return h.Article(a.Attrs(a.Class("max-w-full"), a.Custom("data-page-id", pageID)),
		headerEl,
		h.Div(a.Attrs(a.Class("prose prose-stone max-w-none content-prose")),
			children...,
		),
		h.Nav(a.Attrs(a.Class("flex flex-col gap-4 mt-12 pt-8 border-t border-stone-200 md:flex-row md:justify-between")),
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
func renderBlock(block Block, index int, quizStates map[int]*QuizState, activeHotspot string, isPresenting bool, isFollowing bool, mdl *Model) h.Element {
	switch b := block.(type) {
	case *QuizBlock:
		// Build poll info for this quiz
		pollInfo := getPollInfoForQuiz(b.ID, index, isPresenting, mdl)
		return renderQuiz(b, index, quizStates, pollInfo, isFollowing)
	case *CalloutBlock:
		return renderCallout(b)
	case *AnnotatedImageBlock:
		return renderAnnotatedImage(b, activeHotspot)
	case *ExerciseBlock:
		return renderExercise(b)
	case *AgentBlock:
		return renderAgentLaunchCard(b)
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
		h.H2(a.Attrs(a.Class("mb-3 text-2xl font-semibold text-stone-900")), h.Text("No Content Available")),
		h.P(a.Attrs(a.Class("text-stone-400 max-w-md mx-auto")), h.Text("No modules have been loaded. Please add content to the content/ directory.")),
	)
}

// renderError renders an error message
func renderError(err error) h.Element {
	return h.Div(a.Attrs(a.Class("bg-red-50 border border-red-300 rounded-xl p-8")),
		h.H2(a.Attrs(a.Class("text-red-600 mb-2 text-xl font-semibold")), h.Text("Error")),
		h.P(a.Attrs(a.Class("text-stone-600")), h.Text(err.Error())),
	)
}


// ============================================================================
// Annotation Layer
// ============================================================================

// renderAnnotationLayer renders the canvas overlay and toolbar for presenter annotations
func renderAnnotationLayer(m *Model) h.Element {
	// Only show in live sessions (presenting or following)
	session := m.getPresentingSession()
	if session == nil {
		session = m.getFollowingSession()
	}
	if session == nil {
		return h.Span(a.Attrs())
	}

	// Get strokes for current page
	strokes := session.GetAnnotations(m.CurrentModule, m.CurrentPage)
	strokesJSON := "[]"
	if len(strokes) > 0 {
		if data, err := json.Marshal(strokes); err == nil {
			strokesJSON = string(data)
		}
	}

	isPresenter := "false"
	if m.IsPresenting {
		isPresenter = "true"
	}

	// Annotation overlay div — JS creates canvas inside this
	overlay := h.Div(a.Attrs(
		a.Id("annotation-layer"),
		a.Custom("data-strokes", stdhtml.EscapeString(strokesJSON)),
		a.Custom("data-is-presenter", isPresenter),
		a.Class("absolute top-0 left-0 w-full h-full pointer-events-none z-40"),
	))

	if !m.IsPresenting {
		// Followers only get the overlay (no toolbar)
		return overlay
	}

	// Presenter toolbar — fixed at bottom center of viewport
	toolbar := h.Div(a.Attrs(
		a.Id("annotation-toolbar"),
		a.Class("fixed bottom-6 left-1/2 -translate-x-1/2 z-50 flex items-center gap-2 justify-center"),
	),
		h.Div(a.Attrs(a.Class("flex items-center gap-2 bg-white/95 backdrop-blur-sm border border-stone-200 rounded-full py-2 px-4 shadow-xl")),
			h.Button(a.Attrs(
				a.Id("annotation-tool-pen"),
				a.Class("py-2 px-4 text-sm rounded-full font-medium cursor-pointer transition-all duration-150 border-none bg-stone-100 text-stone-700 hover:bg-stone-200"),
				a.OnClick(`TrainingApp.selectAnnotationTool("pen")`)),
				h.Text("\u270F Pen"),
			),
			h.Button(a.Attrs(
				a.Id("annotation-tool-highlighter"),
				a.Class("py-2 px-4 text-sm rounded-full font-medium cursor-pointer transition-all duration-150 border-none bg-stone-100 text-stone-700 hover:bg-stone-200"),
				a.OnClick(`TrainingApp.selectAnnotationTool("highlighter")`)),
				h.Text("\u2588 Highlight"),
			),
			h.Button(a.Attrs(
				a.Class("py-2 px-4 text-sm rounded-full font-medium cursor-pointer transition-all duration-150 border-none bg-stone-100 text-red-500 hover:bg-red-50"),
				a.OnClick(gt.SendBasicMessageNoArgs("CLEAR_ANNOTATIONS"))),
				h.Text("\u2715 Clear"),
			),
		),
	)

	return h.Div(a.Attrs(),
		overlay,
		toolbar,
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
const inputClasses = "w-full py-3 px-4 border border-stone-300 rounded-lg bg-white text-stone-900 text-base transition-all duration-150 focus:outline-none focus:border-accent focus:ring-2 focus:ring-accent/20 placeholder:text-stone-400"

// renderCohortCodeEntry renders the cohort code entry screen
func renderCohortCodeEntry(m *Model) h.Element {
	return h.Div(a.Attrs(a.Class("w-full max-w-md")),
		h.Div(a.Attrs(a.Class("bg-white border border-stone-200 rounded-2xl p-10 shadow-xl")),
			h.Div(a.Attrs(a.Class("text-center mb-8")),
				h.H1(a.Attrs(a.Class("text-2xl font-bold mb-2 text-stone-900")), h.Text("Welcome")),
				h.P(a.Attrs(a.Class("text-stone-500")), h.Text("Enter your cohort code to begin")),
			),
			h.Form(a.Attrs(
				a.Id("cohort-code-form"),
				a.Class("flex flex-col gap-5"),
				a.OnSubmit(`event.preventDefault(); gotea.updateFormState({message:"SUBMIT_COHORT_CODE"}, "cohort-code-form")`)),
				h.Div(a.Attrs(a.Class("flex flex-col gap-2")),
					h.Label(a.Attrs(a.For("code"), a.Class("text-sm font-medium text-stone-700")), h.Text("Cohort Code")),
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
					a.Class("w-full py-3.5 px-6 bg-accent text-white font-semibold rounded-lg text-base transition-all duration-150 hover:bg-accent-hover cursor-pointer border-none")),
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
		h.Div(a.Attrs(a.Class("bg-white border border-stone-200 rounded-2xl p-10 shadow-xl")),
			h.Div(a.Attrs(a.Class("text-center mb-8")),
				h.H1(a.Attrs(a.Class("text-2xl font-bold mb-2 text-stone-900")), h.Text("Join "+cohortName)),
				h.P(a.Attrs(a.Class("text-stone-500")), h.Text("Enter your details to continue")),
			),
			h.Form(a.Attrs(
				a.Id("login-form"),
				a.Class("flex flex-col gap-5"),
				a.OnSubmit(`event.preventDefault(); gotea.updateFormState({message:"SUBMIT_LOGIN"}, "login-form")`)),
				h.Div(a.Attrs(a.Class("flex flex-col gap-2")),
					h.Label(a.Attrs(a.For("email"), a.Class("text-sm font-medium text-stone-700")), h.Text("Email")),
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
					h.Label(a.Attrs(a.For("name"), a.Class("text-sm font-medium text-stone-700")), h.Text("Your Name")),
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
					a.Class("w-full py-3.5 px-6 bg-accent text-white font-semibold rounded-lg text-base transition-all duration-150 hover:bg-accent-hover cursor-pointer border-none")),
					h.Text("Join Course"),
				),
			),
			h.Button(a.Attrs(
				a.Class("bg-transparent text-stone-500 py-2 mt-4 hover:text-accent cursor-pointer border-none text-sm"),
				a.OnClick(gt.SendBasicMessageNoArgs("LOGOUT"))),
				h.Text("\u2190 Back"),
			),
		),
	)
}

// renderAuthError renders an auth error screen with retry option
func renderAuthError(m *Model) h.Element {
	return h.Div(a.Attrs(a.Class("w-full max-w-md")),
		h.Div(a.Attrs(a.Class("bg-white border border-red-300 rounded-2xl p-10 shadow-xl")),
			h.Div(a.Attrs(a.Class("text-center mb-8")),
				h.H1(a.Attrs(a.Class("text-2xl font-bold mb-2 text-stone-900")), h.Text("Oops")),
			),
			h.P(a.Attrs(a.Class("text-red-600 text-base mb-6 leading-relaxed")), h.Text(m.AuthError)),
			h.Button(a.Attrs(
				a.Class("w-full py-3.5 px-6 bg-accent text-white font-semibold rounded-lg text-base transition-all duration-150 hover:bg-accent-hover cursor-pointer border-none"),
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
		return h.Div(a.Attrs(a.Class("bg-red-50 border border-red-300 rounded-xl p-8")),
			h.H2(a.Attrs(a.Class("text-red-600 mb-2 text-xl font-semibold")), h.Text("Access Denied")),
			h.P(a.Attrs(a.Class("text-stone-500")), h.Text("You do not have permission to access this area.")),
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

		cardClass := "bg-white border border-stone-200 rounded-xl p-5 shadow-sm"
		if expired {
			cardClass += " opacity-60"
		}

		statusBadge := h.Span(a.Attrs(a.Class("text-xs py-1 px-3 rounded-full font-medium bg-accent/15 text-accent")), h.Text("Active"))
		if expired {
			statusBadge = h.Span(a.Attrs(a.Class("text-xs py-1 px-3 rounded-full font-medium bg-red-500/15 text-red-600")), h.Text("Expired"))
		}

		cohortCards = append(cohortCards, h.Div(a.Attrs(a.Class(cardClass)),
			h.Div(a.Attrs(a.Class("flex justify-between items-center mb-3")),
				h.H3(a.Attrs(a.Class("text-lg font-semibold text-stone-900")), h.Text(cohort.Name)),
				statusBadge,
			),
			h.Div(a.Attrs(a.Class("mb-3")),
				h.Code(a.Attrs(a.Class("font-mono text-base py-2 px-3 bg-stone-50 border border-stone-200 rounded-md text-accent inline-block")), h.Text(cohort.Code)),
			),
			h.Div(a.Attrs(a.Class("flex gap-6 mb-4 text-sm text-stone-500")),
				h.Span(a.Attrs(), h.Text(fmt.Sprintf("Students: %d/%d", count, cohort.MaxStudents))),
				h.Span(a.Attrs(), h.Text(fmt.Sprintf("Expires: %s", cohort.ExpiresAt.Format("Jan 2, 2006")))),
			),
			h.Div(a.Attrs(a.Class("flex gap-2")),
				h.Button(a.Attrs(
					a.Class("py-1.5 px-3 text-sm rounded-md bg-stone-200 text-stone-700 font-medium cursor-pointer transition-all duration-150 hover:bg-stone-300 border-none"),
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
			h.P(a.Attrs(a.Class("text-stone-400")), h.Text("No cohorts yet. Create one to get started.")),
		))
	}

	return h.Div(a.Attrs(a.Class("max-w-[900px] mx-auto py-4")),
		h.Div(a.Attrs(a.Class("flex justify-between items-center mb-10 flex-wrap gap-4")),
			h.H1(a.Attrs(a.Class("text-3xl font-bold tracking-tight text-stone-900")), h.Text("Cohort Management")),
			h.A(a.Attrs(
				a.Href("/"),
				a.Class("py-2.5 px-5 bg-stone-200 text-stone-700 font-semibold rounded-lg transition-all duration-150 hover:bg-stone-300 no-underline")),
				h.Text("\u2190 Back to Course"),
			),
		),
		h.Div(a.Attrs(a.Class("mb-10")),
			h.H2(a.Attrs(a.Class("text-xl font-semibold mb-5 text-stone-900")), h.Text("Create New Cohort")),
			h.Form(a.Attrs(
				a.Id("create-cohort-form"),
				a.Class("bg-white border border-stone-200 rounded-xl p-6 shadow-sm"),
				a.OnSubmit(`event.preventDefault(); gotea.updateFormState({message:"CREATE_COHORT"}, "create-cohort-form")`)),
				h.Div(a.Attrs(a.Class("grid grid-cols-1 gap-4 mb-5 sm:grid-cols-[2fr_1fr_1fr]")),
					h.Div(a.Attrs(a.Class("flex flex-col gap-2")),
						h.Label(a.Attrs(a.For("name"), a.Class("text-sm font-medium text-stone-700")), h.Text("Cohort Name")),
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
						h.Label(a.Attrs(a.For("expiresAt"), a.Class("text-sm font-medium text-stone-700")), h.Text("Expires")),
						h.Input(a.Attrs(
							a.Type("date"),
							a.Id("expiresAt"),
							a.Name("expiresAt"),
							a.Class(inputClasses),
							a.Required(true),
						)),
					),
					h.Div(a.Attrs(a.Class("flex flex-col gap-2")),
						h.Label(a.Attrs(a.For("maxStudents"), a.Class("text-sm font-medium text-stone-700")), h.Text("Max Students")),
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
					a.Class("py-3 px-6 bg-accent text-white font-semibold rounded-lg transition-all duration-150 hover:bg-accent-hover cursor-pointer border-none")),
					h.Text("Create Cohort"),
				),
			),
		),
		h.Div(a.Attrs(a.Class("mb-10")),
			h.H2(a.Attrs(a.Class("text-xl font-semibold mb-5 text-stone-900")), h.Text("Existing Cohorts")),
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
