package parsing

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

// Regex to match fenced code blocks with our custom languages
var fencedBlockRegex = regexp.MustCompile("(?s)```(quiz|math|callout|annotated-image|terminal-replay|exercise|agent-demo|agent)\n(.*?)```")

// ParseMarkdown parses markdown content with custom extensions.
// baseDir is the directory containing the markdown file, used for resolving relative paths.
func ParseMarkdown(content []byte, baseDir string) (*ParseResult, error) {
	// First pass: extract custom blocks and replace with placeholders
	var blocks []Block
	blockIndex := 0

	processedContent := fencedBlockRegex.ReplaceAllFunc(content, func(match []byte) []byte {
		// Parse the block type and content
		submatches := fencedBlockRegex.FindSubmatch(match)
		if len(submatches) < 3 {
			return match // Return unchanged if parse fails
		}

		blockType := string(submatches[1])
		blockContent := submatches[2]

		block, err := parseBlock(blockType, blockContent, baseDir)
		if err != nil {
			// On parse error, leave the original content (will render as code block)
			return match
		}

		blocks = append(blocks, block)
		placeholder := fmt.Sprintf(`<div data-block-index="%d"></div>`, blockIndex)
		blockIndex++

		return []byte(placeholder)
	})

	// Second pass: render remaining markdown with Goldmark
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,            // GitHub Flavored Markdown
			extension.Typographer,    // Smart quotes, etc.
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(), // Allow raw HTML in markdown
		),
	)

	var buf bytes.Buffer
	if err := md.Convert(processedContent, &buf); err != nil {
		return nil, fmt.Errorf("failed to convert markdown: %w", err)
	}

	return &ParseResult{
		HTML:   buf.Bytes(),
		Blocks: blocks,
	}, nil
}

// parseBlock parses a custom block based on its type
func parseBlock(blockType string, content []byte, baseDir string) (Block, error) {
	switch blockType {
	case "quiz":
		return parseQuizBlock(content)
	case "callout":
		return parseCalloutBlock(content)
	case "annotated-image":
		return parseAnnotatedImageBlock(content)
	case "exercise":
		return parseExerciseBlock(content)
	case "agent", "agent-demo":
		return parseAgentBlock(content, baseDir)
	default:
		return nil, fmt.Errorf("unknown block type: %s", blockType)
	}
}

// Quiz block parsing
type QuizBlock struct {
	ID          string   `yaml:"id"`
	Type        string   `yaml:"type"`
	Question    string   `yaml:"question"`
	Options     []string `yaml:"options"`
	Answer      int      `yaml:"answer"`
	Answers     []int    `yaml:"answers"`
	Explanation string   `yaml:"explanation"`
}

func (b *QuizBlock) BlockType() string { return "quiz" }
func (b *QuizBlock) BlockID() string   { return b.ID }

func parseQuizBlock(content []byte) (*QuizBlock, error) {
	var block QuizBlock
	if err := yaml.Unmarshal(content, &block); err != nil {
		return nil, fmt.Errorf("failed to parse quiz block: %w", err)
	}
	if block.Type == "" {
		block.Type = "multiple-choice"
	}
	return &block, nil
}

// Callout block parsing
type CalloutBlock struct {
	Type    string `yaml:"type"`
	Title   string `yaml:"title"`
	Content string `yaml:"content"`
}

func (b *CalloutBlock) BlockType() string { return "callout" }
func (b *CalloutBlock) BlockID() string   { return "" }

func parseCalloutBlock(content []byte) (*CalloutBlock, error) {
	var block CalloutBlock
	if err := yaml.Unmarshal(content, &block); err != nil {
		return nil, fmt.Errorf("failed to parse callout block: %w", err)
	}
	if block.Type == "" {
		block.Type = "info"
	}
	return &block, nil
}

// Annotated image block parsing
type AnnotatedImageBlock struct {
	ID       string    `yaml:"id"`
	Src      string    `yaml:"src"`
	Alt      string    `yaml:"alt"`
	Hotspots []Hotspot `yaml:"hotspots"`
}

type Hotspot struct {
	X      string `yaml:"x"`
	Y      string `yaml:"y"`
	Label  string `yaml:"label"`
	Detail string `yaml:"detail"`
}

func (b *AnnotatedImageBlock) BlockType() string { return "annotated-image" }
func (b *AnnotatedImageBlock) BlockID() string   { return b.ID }

func parseAnnotatedImageBlock(content []byte) (*AnnotatedImageBlock, error) {
	var block AnnotatedImageBlock
	if err := yaml.Unmarshal(content, &block); err != nil {
		return nil, fmt.Errorf("failed to parse annotated-image block: %w", err)
	}
	return &block, nil
}

// Exercise block parsing
type ExerciseBlock struct {
	ID         string             `yaml:"id"`
	Language   string             `yaml:"language"`
	Prompt     string             `yaml:"prompt"`
	Starter    string             `yaml:"starter"`
	Validation ExerciseValidation `yaml:"validation"`
}

type ExerciseValidation struct {
	Type     string   `yaml:"type"`
	Expected string   `yaml:"expected"`
	Keywords []string `yaml:"keywords"`
}

func (b *ExerciseBlock) BlockType() string { return "exercise" }
func (b *ExerciseBlock) BlockID() string   { return b.ID }

func parseExerciseBlock(content []byte) (*ExerciseBlock, error) {
	var block ExerciseBlock
	if err := yaml.Unmarshal(content, &block); err != nil {
		return nil, fmt.Errorf("failed to parse exercise block: %w", err)
	}
	return &block, nil
}

// Agent block parsing
type AgentBlock struct {
	ID         string            `yaml:"id"`
	Title      string            `yaml:"title"`
	ModelLabel string            `yaml:"model_label"`
	System     string            `yaml:"system"`
	Scratchpad map[string]string `yaml:"scratchpad"`
	Tools      []string          `yaml:"tools"`
	Visibility AgentVisibility   `yaml:"visibility"`
	Sidebar    AgentSidebarConfig `yaml:"sidebar"`
	Script     []ScriptEvent     `yaml:"script"`
}

type ScriptEvent struct {
	Type            string            `yaml:"type"`
	Text            string            `yaml:"text"`
	Content         string            `yaml:"content"`
	Tokens          int               `yaml:"tokens"`
	Tool            string            `yaml:"tool"`
	Args            map[string]string `yaml:"args"`
	Summary         string            `yaml:"summary"`
	ResetScratchpad bool              `yaml:"reset_scratchpad"`
	Note            string            `yaml:"note"`
}

type AgentVisibility struct {
	SystemPrompt string `yaml:"system_prompt"`
	ToolCalls    string `yaml:"tool_calls"`
	FullContext  string `yaml:"full_context"`
	TokenCount   string `yaml:"token_count"`
	ModelName    string `yaml:"model_name"`
}

type AgentSidebarConfig struct {
	Width     string `yaml:"width"`
	StartOpen *bool  `yaml:"start_open"`
}

func (b *AgentBlock) BlockType() string { return "agent" }
func (b *AgentBlock) BlockID() string   { return b.ID }

func parseAgentBlock(content []byte, baseDir string) (*AgentBlock, error) {
	// Check if the content is a path reference to an external YAML file
	var pathRef struct {
		Path string `yaml:"path"`
	}
	if err := yaml.Unmarshal(content, &pathRef); err == nil && pathRef.Path != "" {
		// Resolve the path: if it starts with /content/, treat as relative to project root
		// by going up from baseDir to find the content/ ancestor
		resolvedPath := pathRef.Path
		if filepath.IsAbs(resolvedPath) {
			// Absolute paths like /content/module-prompting/foo.yaml:
			// find the content dir by walking up from baseDir
			dir := baseDir
			for dir != "/" && dir != "." {
				if filepath.Base(dir) == "content" {
					resolvedPath = filepath.Join(filepath.Dir(dir), resolvedPath)
					break
				}
				dir = filepath.Dir(dir)
			}
		} else {
			resolvedPath = filepath.Join(baseDir, resolvedPath)
		}

		data, err := os.ReadFile(resolvedPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load agent block from path %q: %w", pathRef.Path, err)
		}
		// Strip markdown fences if the file is wrapped in them (e.g. ```agent\n...\n```)
		trimmed := bytes.TrimSpace(data)
		if bytes.HasPrefix(trimmed, []byte("```")) {
			if idx := bytes.IndexByte(trimmed, '\n'); idx >= 0 {
				trimmed = trimmed[idx+1:]
			}
			trimmed = bytes.TrimSuffix(bytes.TrimSpace(trimmed), []byte("```"))
		}
		content = trimmed
	}

	var block AgentBlock
	if err := yaml.Unmarshal(content, &block); err != nil {
		return nil, fmt.Errorf("failed to parse agent block: %w", err)
	}

	// Defaults
	if block.Sidebar.Width == "" {
		block.Sidebar.Width = "40%"
	}
	if block.Sidebar.StartOpen == nil {
		t := true
		block.Sidebar.StartOpen = &t
	}

	// Visibility defaults
	if block.Visibility.SystemPrompt == "" {
		block.Visibility.SystemPrompt = "hidden"
	}
	if block.Visibility.ToolCalls == "" {
		block.Visibility.ToolCalls = "hidden"
	}
	if block.Visibility.FullContext == "" {
		block.Visibility.FullContext = "hidden"
	}
	if block.Visibility.TokenCount == "" {
		block.Visibility.TokenCount = "hidden"
	}
	if block.Visibility.ModelName == "" {
		block.Visibility.ModelName = "hidden"
	}

	// Validation
	if len(block.Script) == 0 {
		return nil, fmt.Errorf("agent block %q: script must not be empty", block.ID)
	}
	first := block.Script[0].Type
	if first != "note" && first != "user" {
		return nil, fmt.Errorf("agent block %q: script must start with note or user, got %q", block.ID, first)
	}

	for _, v := range []string{block.Visibility.SystemPrompt, block.Visibility.ToolCalls, block.Visibility.FullContext} {
		if v != "visible" && v != "hidden" && v != "toggleable" {
			return nil, fmt.Errorf("agent block %q: invalid visibility value %q", block.ID, v)
		}
	}
	for _, v := range []string{block.Visibility.TokenCount, block.Visibility.ModelName} {
		if v != "visible" && v != "hidden" {
			return nil, fmt.Errorf("agent block %q: invalid visibility value %q (must be visible or hidden)", block.ID, v)
		}
	}

	for i, ev := range block.Script {
		if ev.Type == "compaction" && ev.Summary == "" {
			return nil, fmt.Errorf("agent block %q: compaction event at index %d must have a summary", block.ID, i)
		}
		if ev.Type == "tool_call" && ev.Tool == "scratchpad_write" {
			if ev.Args["filename"] == "" || ev.Args["content"] == "" {
				return nil, fmt.Errorf("agent block %q: scratchpad_write tool_call at index %d must have filename and content args", block.ID, i)
			}
		}
	}

	return &block, nil
}
