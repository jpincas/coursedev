package parsing

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

// Regex to match fenced code blocks with our custom languages
var fencedBlockRegex = regexp.MustCompile("(?s)```(quiz|math|callout|annotated-image|terminal-replay|exercise)\n(.*?)```")

// ParseMarkdown parses markdown content with custom extensions
func ParseMarkdown(content []byte) (*ParseResult, error) {
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

		block, err := parseBlock(blockType, blockContent)
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
func parseBlock(blockType string, content []byte) (Block, error) {
	switch blockType {
	case "quiz":
		return parseQuizBlock(content)
	case "callout":
		return parseCalloutBlock(content)
	case "annotated-image":
		return parseAnnotatedImageBlock(content)
	case "exercise":
		return parseExerciseBlock(content)
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
