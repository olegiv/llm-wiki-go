package wikilint

// sourcesSection is the heading wiki pages must use to list their sources.
const sourcesSection = "Sources"

// trackedSections are H2 sections whose emptiness is reported.
var trackedSections = []string{"Sources", "Contradictions", "Open questions"}

// checkH1 returns a non-empty message when the page lacks a single H1 title.
func checkH1(p *Page) string {
	switch {
	case p.H1Count == 0:
		return "missing H1 title"
	case p.H1Count > 1:
		return "multiple H1 titles"
	}
	return ""
}

// requiresSources reports whether a page needs a `## Sources` section.
// index.md and log.md are exempt.
func requiresSources(relPath string) bool {
	return relPath != "index.md" && relPath != "log.md"
}

// checkRequiredSources returns a non-empty message when the Sources section
// is required but missing from the page.
func checkRequiredSources(p *Page) string {
	if !requiresSources(p.RelPath) {
		return ""
	}
	if _, ok := p.Sections[sourcesSection]; !ok {
		return "missing required `## Sources` section"
	}
	return ""
}

// checkEmptySections returns messages for any tracked section that is present
// but has no non-whitespace content.
func checkEmptySections(p *Page) []string {
	var msgs []string
	for _, name := range trackedSections {
		body, ok := p.Sections[name]
		if !ok {
			continue
		}
		if body == "" {
			msgs = append(msgs, "`## "+name+"` section is empty")
		}
	}
	return msgs
}
