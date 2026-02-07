package main

// Block is the interface for all interactive block types
type Block interface {
	BlockType() string
	BlockID() string
}

// QuizBlock represents a quiz question
type QuizBlock struct {
	ID          string   `yaml:"id"`
	Type        string   `yaml:"type"` // multiple-choice, multi-select, true-false, ordering, free-text
	Question    string   `yaml:"question"`
	Options     []string `yaml:"options"`
	Answer      int      `yaml:"answer"`       // For single answer (index)
	Answers     []int    `yaml:"answers"`      // For multi-select
	Explanation string   `yaml:"explanation"`
}

func (b *QuizBlock) BlockType() string { return "quiz" }
func (b *QuizBlock) BlockID() string   { return b.ID }

// MermaidBlock represents a Mermaid diagram
type MermaidBlock struct {
	Source string // The raw Mermaid source code
}

func (b *MermaidBlock) BlockType() string { return "mermaid" }
func (b *MermaidBlock) BlockID() string   { return "" }

// CalloutBlock represents an admonition/callout
type CalloutBlock struct {
	Type    string `yaml:"type"`    // info, warning, tip, danger, note
	Title   string `yaml:"title"`
	Content string `yaml:"content"`
}

func (b *CalloutBlock) BlockType() string { return "callout" }
func (b *CalloutBlock) BlockID() string   { return "" }

// AnnotatedImageBlock represents an image with interactive hotspots
type AnnotatedImageBlock struct {
	ID       string    `yaml:"id"`
	Src      string    `yaml:"src"`
	Alt      string    `yaml:"alt"`
	Hotspots []Hotspot `yaml:"hotspots"`
}

type Hotspot struct {
	X      string `yaml:"x"`      // e.g., "30%"
	Y      string `yaml:"y"`      // e.g., "45%"
	Label  string `yaml:"label"`
	Detail string `yaml:"detail"`
}

func (b *AnnotatedImageBlock) BlockType() string { return "annotated-image" }
func (b *AnnotatedImageBlock) BlockID() string   { return b.ID }

// ExerciseBlock represents a code exercise
type ExerciseBlock struct {
	ID         string             `yaml:"id"`
	Language   string             `yaml:"language"`
	Prompt     string             `yaml:"prompt"`
	Starter    string             `yaml:"starter"`
	Validation ExerciseValidation `yaml:"validation"`
}

type ExerciseValidation struct {
	Type     string `yaml:"type"`     // output-match, contains-keywords, manual
	Expected string `yaml:"expected"` // For output-match
	Keywords []string `yaml:"keywords"` // For contains-keywords
}

func (b *ExerciseBlock) BlockType() string { return "exercise" }
func (b *ExerciseBlock) BlockID() string   { return b.ID }
