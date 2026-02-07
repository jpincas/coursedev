package main

// CourseGraph represents the course structure and prerequisites
type CourseGraph struct {
	ModuleOrder   []string            // Canonical ordering of modules
	Prerequisites map[string][]string // module ID -> required module IDs
}

// ModuleIndex returns the 1-based index of a module in ModuleOrder, or 0 if not found
func (cg *CourseGraph) ModuleIndex(moduleID string) int {
	for i, id := range cg.ModuleOrder {
		if id == moduleID {
			return i + 1
		}
	}
	return 0
}

// NextModule returns the ID of the next module after the given one, or "" if it's the last
func (cg *CourseGraph) NextModule(moduleID string) string {
	for i, id := range cg.ModuleOrder {
		if id == moduleID && i+1 < len(cg.ModuleOrder) {
			return cg.ModuleOrder[i+1]
		}
	}
	return ""
}

// PrerequisitesMet checks if all prerequisites for a module are satisfied
func (cg *CourseGraph) PrerequisitesMet(moduleID string, progress map[string]*ModuleProgress) bool {
	prereqs, ok := cg.Prerequisites[moduleID]
	if !ok {
		return true // No prerequisites defined
	}

	for _, req := range prereqs {
		p, exists := progress[req]
		if !exists || p.CompletedAt == nil {
			return false
		}
	}
	return true
}

// AvailableModules returns the list of modules the learner can access
func (cg *CourseGraph) AvailableModules(progress map[string]*ModuleProgress) []string {
	var available []string
	for _, id := range cg.ModuleOrder {
		if cg.PrerequisitesMet(id, progress) {
			available = append(available, id)
		}
	}
	return available
}

// ValidateDAG checks that the prerequisite graph is acyclic
func (cg *CourseGraph) ValidateDAG() error {
	// Track visited and in-progress nodes for cycle detection
	visited := make(map[string]bool)
	inProgress := make(map[string]bool)

	var visit func(id string) bool
	visit = func(id string) bool {
		if inProgress[id] {
			return false // Cycle detected
		}
		if visited[id] {
			return true // Already processed
		}

		inProgress[id] = true
		for _, prereq := range cg.Prerequisites[id] {
			if !visit(prereq) {
				return false
			}
		}
		inProgress[id] = false
		visited[id] = true
		return true
	}

	for _, id := range cg.ModuleOrder {
		if !visit(id) {
			return &CycleError{ModuleID: id}
		}
	}

	return nil
}

// CycleError indicates a cycle was detected in the prerequisite graph
type CycleError struct {
	ModuleID string
}

func (e *CycleError) Error() string {
	return "cycle detected in prerequisites involving module: " + e.ModuleID
}
