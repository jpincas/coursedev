package parsing

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

var frontmatterDelimiter = []byte("---")

// ExtractFrontmatter extracts YAML frontmatter from markdown content
// Returns the parsed metadata, the remaining markdown content, and any error
func ExtractFrontmatter(content []byte) (*PageMeta, []byte, error) {
	content = bytes.TrimSpace(content)

	// Check if content starts with frontmatter delimiter
	if !bytes.HasPrefix(content, frontmatterDelimiter) {
		// No frontmatter, return empty meta and full content
		return &PageMeta{}, content, nil
	}

	// Find the end of frontmatter
	rest := content[len(frontmatterDelimiter):]
	endIdx := bytes.Index(rest, frontmatterDelimiter)
	if endIdx == -1 {
		// No closing delimiter, treat as no frontmatter
		return &PageMeta{}, content, nil
	}

	// Extract and parse frontmatter
	frontmatter := bytes.TrimSpace(rest[:endIdx])
	markdown := bytes.TrimSpace(rest[endIdx+len(frontmatterDelimiter):])

	var meta PageMeta
	if err := yaml.Unmarshal(frontmatter, &meta); err != nil {
		return nil, nil, err
	}

	return &meta, markdown, nil
}
