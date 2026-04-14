package wikilint

import (
	"bufio"
	"strings"
)

// Page is a parsed Markdown wiki page.
type Page struct {
	// RepoPath is the path used in reports (slash-separated, relative to cwd).
	RepoPath string
	// RelPath is the wiki-root-relative slash-separated path, e.g. "entities/foo.md".
	RelPath string
	// Title is the text of the first H1, if any.
	Title string
	// H1Count is the number of level-1 headings found.
	H1Count int
	// Sections maps H2 heading text to the trimmed body beneath it.
	Sections map[string]string
	// SectionOrder preserves the order H2 sections were encountered.
	SectionOrder []string
	// Links are Markdown links from non-code content.
	Links []Link
	// Wikilinks are Obsidian-style [[target]] links from non-code content.
	Wikilinks []Wikilink
	// Content is the raw page content.
	Content string
}

// Link is a Markdown link [text](target).
type Link struct {
	Target string
}

// Wikilink is an Obsidian-style [[target]] link.
type Wikilink struct {
	Target string
}

// ParsePage parses page content into a Page value.
func ParsePage(repoPath, relPath, content string) *Page {
	p := &Page{
		RepoPath: repoPath,
		RelPath:  relPath,
		Content:  content,
		Sections: map[string]string{},
	}
	parseHeadings(p)
	p.Links, p.Wikilinks = extractLinks(content)
	return p
}

// parseHeadings fills Title, H1Count, and Sections by scanning content line-by-line.
// It skips fenced code blocks so headings inside them don't count.
func parseHeadings(p *Page) {
	scanner := bufio.NewScanner(strings.NewReader(p.Content))
	scanner.Buffer(make([]byte, 64*1024), 256*1024)

	var (
		inFence        bool
		currentSection string
		sectionBody    strings.Builder
	)

	flush := func() {
		if currentSection == "" {
			return
		}
		if _, exists := p.Sections[currentSection]; !exists {
			p.SectionOrder = append(p.SectionOrder, currentSection)
		}
		p.Sections[currentSection] = strings.TrimSpace(sectionBody.String())
		sectionBody.Reset()
	}

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "```") || strings.HasPrefix(trimmed, "~~~") {
			inFence = !inFence
			if currentSection != "" {
				sectionBody.WriteString(line)
				sectionBody.WriteByte('\n')
			}
			continue
		}
		if inFence {
			if currentSection != "" {
				sectionBody.WriteString(line)
				sectionBody.WriteByte('\n')
			}
			continue
		}
		switch {
		case strings.HasPrefix(trimmed, "# ") || trimmed == "#":
			p.H1Count++
			title := strings.TrimSpace(strings.TrimPrefix(trimmed, "#"))
			if p.Title == "" {
				p.Title = title
			}
			flush()
			currentSection = ""
		case strings.HasPrefix(trimmed, "## ") || trimmed == "##":
			flush()
			currentSection = strings.TrimSpace(strings.TrimPrefix(trimmed, "##"))
		default:
			if currentSection != "" {
				sectionBody.WriteString(line)
				sectionBody.WriteByte('\n')
			}
		}
	}
	// scanner.Err() is intentionally not checked: if a line exceeds
	// 256 KiB the scan stops and the page is partially parsed. The
	// linter will report structural issues (missing H1, sections, etc.)
	// which is adequate for this edge case.
	flush()
}
