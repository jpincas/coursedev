package main

import (
	gt "github.com/jpincas/go-tea"
)

func (m *Model) handleAgentAdvance(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	agent := mdl.Agent
	if agent == nil {
		return gt.Respond()
	}

	script := agent.Config.Script
	if agent.ScriptIndex >= len(script) {
		return gt.Respond() // Script complete
	}

	event := script[agent.ScriptIndex]

	switch event.Type {
	case "note":
		agent.CurrentNote = event.Text
		agent.ScriptIndex++

	case "user":
		agent.CurrentNote = ""
		agent.ScriptIndex = processEventGroup(agent, agent.ScriptIndex)

	case "compaction":
		agent.CurrentNote = ""
		processCompaction(agent, event)
		agent.ScriptIndex++

	case "clear":
		agent.CurrentNote = ""
		processClear(agent, event)
		agent.ScriptIndex++
	}

	// If presenting, broadcast agent state to followers
	if mdl.IsPresenting {
		if session := mdl.getPresentingSession(); session != nil {
			session.SetAgent(mdl.Agent)
			app.Broadcast()
		}
	}

	return gt.Respond()
}

func (m *Model) handleAgentReset(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	if mdl.Agent == nil {
		return gt.Respond()
	}
	mdl.Agent = initAgentState(mdl.Agent.Config)

	if mdl.IsPresenting {
		if session := mdl.getPresentingSession(); session != nil {
			session.SetAgent(mdl.Agent)
			app.Broadcast()
		}
	}

	return gt.Respond()
}

func (m *Model) handleAgentToggleSystem(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	if mdl.Agent != nil {
		mdl.Agent.ShowSystem = !mdl.Agent.ShowSystem
	}
	return gt.Respond()
}

func (m *Model) handleAgentToggleTools(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	if mdl.Agent != nil {
		mdl.Agent.ShowToolCalls = !mdl.Agent.ShowToolCalls
	}
	return gt.Respond()
}

func (m *Model) handleAgentToggleContext(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	if mdl.Agent != nil {
		mdl.Agent.ShowFullContext = !mdl.Agent.ShowFullContext
	}
	return gt.Respond()
}

func (m *Model) handleAgentToggleSidebar(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	if mdl.Agent != nil {
		mdl.Agent.SidebarOpen = !mdl.Agent.SidebarOpen
	}
	return gt.Respond()
}

func (m *Model) handleAgentToggleScratchpad(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	if mdl.Agent != nil {
		mdl.Agent.ScratchpadOpen = !mdl.Agent.ScratchpadOpen
	}
	return gt.Respond()
}

func (m *Model) handleAgentViewFile(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	if mdl.Agent == nil {
		return gt.Respond()
	}
	filename := msg.ArgsToString()
	if mdl.Agent.ViewingFile == filename {
		mdl.Agent.ViewingFile = "" // Toggle off
	} else {
		mdl.Agent.ViewingFile = filename
	}
	return gt.Respond()
}

func (m *Model) handleAgentOpenWorkspace(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	if mdl.Agent != nil {
		mdl.Agent.WorkspaceOpen = true
	}
	return gt.Respond()
}

func (m *Model) handleAgentCloseWorkspace(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	if mdl.Agent != nil {
		mdl.Agent.WorkspaceOpen = false
	}
	return gt.Respond()
}
