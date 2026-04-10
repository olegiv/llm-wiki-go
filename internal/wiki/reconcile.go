package wiki

import "strings"

// ContradictionEntry is a single source's claim about a disputed fact.
type ContradictionEntry struct {
	Source string
	Claim  string
}

// FormatContradiction renders a deterministic Markdown block for a
// `## Contradictions` section describing a disagreement between sources.
// The block is suitable for appending beneath the H2 heading.
func FormatContradiction(topic string, entries []ContradictionEntry) string {
	if topic == "" {
		topic = "Unresolved contradiction"
	}
	var b strings.Builder
	b.WriteString("### ")
	b.WriteString(topic)
	b.WriteString("\n\n")
	for _, e := range entries {
		b.WriteString("- ")
		b.WriteString(e.Source)
		b.WriteString(": ")
		b.WriteString(e.Claim)
		b.WriteByte('\n')
	}
	return b.String()
}
