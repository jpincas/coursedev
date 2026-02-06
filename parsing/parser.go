package parsing

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// ParseResult contains the output of parsing a markdown file
type ParseResult struct {
	HTML   []byte
	Blocks []Block
}

// Block is the interface for interactive block types
type Block interface {
	BlockType() string
	BlockID() string
}

// CourseConfig is the top-level course.yaml structure
type CourseConfig struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Modules     []string `yaml:"modules"` // Module directory names in order
}

// ModuleMeta contains module-level metadata from _module.yaml
type ModuleMeta struct {
	Title             string             `yaml:"title"`
	Description       string             `yaml:"description"`
	Prerequisites     []string           `yaml:"prerequisites"`
	Difficulty        string             `yaml:"difficulty"`
	Roles             []string           `yaml:"roles"`
	EstimatedDuration string             `yaml:"estimated_duration"`
	Completion        CompletionCriteria `yaml:"completion"`
}

// CompletionCriteria defines when a module is considered complete
type CompletionCriteria struct {
	RequireAllPages bool    `yaml:"require_all_pages"`
	RequireQuizzes  bool    `yaml:"require_quizzes"`
	MinQuizScore    float64 `yaml:"min_quiz_score"`
}

// PageMeta contains page-level metadata from frontmatter
type PageMeta struct {
	Title    string   `yaml:"title"`
	Duration string   `yaml:"duration"`
	Tags     []string `yaml:"tags"`
	Notes    string   `yaml:"notes"`
}

// Module represents a parsed training module
type Module struct {
	ID    string
	Meta  ModuleMeta
	Pages []Page
}

// Page represents a parsed page within a module
type Page struct {
	Filename      string
	Meta          PageMeta
	NarrativeHTML []byte
	Blocks        []Block
}

// CourseGraph represents the course structure
type CourseGraph struct {
	ModuleOrder   []string
	Prerequisites map[string][]string
}

// LoadCourse loads and parses all course content from a directory
func LoadCourse(contentDir string) (*CourseGraph, map[string]*Module, error) {
	// Load course.yaml if it exists
	courseConfig, err := loadCourseConfig(contentDir)
	if err != nil && !os.IsNotExist(err) {
		return nil, nil, fmt.Errorf("failed to load course.yaml: %w", err)
	}

	// Discover and parse modules
	modules := make(map[string]*Module)
	var moduleOrder []string

	// If course.yaml specifies module order, use that
	if courseConfig != nil && len(courseConfig.Modules) > 0 {
		moduleOrder = courseConfig.Modules
		for _, modID := range moduleOrder {
			modPath := filepath.Join(contentDir, modID)
			mod, err := parseModule(modID, modPath)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse module %s: %w", modID, err)
			}
			modules[modID] = mod
		}
	} else {
		// Auto-discover modules from directories
		entries, err := os.ReadDir(contentDir)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read content directory: %w", err)
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			if strings.HasPrefix(entry.Name(), ".") || entry.Name() == "shared-assets" {
				continue
			}

			modPath := filepath.Join(contentDir, entry.Name())
			// Check if this directory has a _module.yaml
			if _, err := os.Stat(filepath.Join(modPath, "_module.yaml")); os.IsNotExist(err) {
				continue
			}

			mod, err := parseModule(entry.Name(), modPath)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse module %s: %w", entry.Name(), err)
			}
			modules[entry.Name()] = mod
			moduleOrder = append(moduleOrder, entry.Name())
		}

		// Sort modules alphabetically if not specified in course.yaml
		sort.Strings(moduleOrder)
	}

	// Build prerequisites map
	prerequisites := make(map[string][]string)
	for id, mod := range modules {
		prerequisites[id] = mod.Meta.Prerequisites
	}

	course := &CourseGraph{
		ModuleOrder:   moduleOrder,
		Prerequisites: prerequisites,
	}

	return course, modules, nil
}

// loadCourseConfig loads the course.yaml file
func loadCourseConfig(contentDir string) (*CourseConfig, error) {
	path := filepath.Join(contentDir, "course.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config CourseConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse course.yaml: %w", err)
	}

	return &config, nil
}

// parseModule parses a single module directory
func parseModule(id, modPath string) (*Module, error) {
	// Load module metadata
	meta, err := loadModuleMeta(modPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load module metadata: %w", err)
	}

	// Find and parse all markdown files
	entries, err := os.ReadDir(modPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read module directory: %w", err)
	}

	var pages []Page
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		if strings.HasPrefix(entry.Name(), "_") {
			continue
		}

		pagePath := filepath.Join(modPath, entry.Name())
		page, err := ParsePage(pagePath)
		if err != nil {
			return nil, fmt.Errorf("failed to parse page %s: %w", entry.Name(), err)
		}
		page.Filename = entry.Name()
		pages = append(pages, *page)
	}

	// Sort pages by filename (assumes numeric prefixes like 01-, 02-)
	sort.Slice(pages, func(i, j int) bool {
		return pages[i].Filename < pages[j].Filename
	})

	return &Module{
		ID:    id,
		Meta:  *meta,
		Pages: pages,
	}, nil
}

// loadModuleMeta loads the _module.yaml file
func loadModuleMeta(modPath string) (*ModuleMeta, error) {
	path := filepath.Join(modPath, "_module.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var meta ModuleMeta
	if err := yaml.Unmarshal(data, &meta); err != nil {
		return nil, fmt.Errorf("failed to parse _module.yaml: %w", err)
	}

	return &meta, nil
}

// ParsePage parses a single markdown file
func ParsePage(path string) (*Page, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Extract frontmatter
	meta, markdown, err := ExtractFrontmatter(content)
	if err != nil {
		return nil, fmt.Errorf("failed to extract frontmatter: %w", err)
	}

	// Parse markdown with custom extensions
	result, err := ParseMarkdown(markdown)
	if err != nil {
		return nil, fmt.Errorf("failed to parse markdown: %w", err)
	}

	return &Page{
		Meta:          *meta,
		NarrativeHTML: result.HTML,
		Blocks:        result.Blocks,
	}, nil
}
