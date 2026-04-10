package wikilint

import (
	"path"
	"regexp"
	"strings"
)

var (
	codeFenceRE  = regexp.MustCompile("(?s)```.*?```")
	codeFenceRE2 = regexp.MustCompile("(?s)~~~.*?~~~")
	inlineCodeRE = regexp.MustCompile("`[^`\n]*`")
	imageRE      = regexp.MustCompile(`!\[[^\]]*\]\([^)]*\)`)
	mdLinkRE     = regexp.MustCompile(`\[([^\]]*)\]\(([^)\s]+)(?:\s+"[^"]*")?\)`)
	wikilinkRE   = regexp.MustCompile(`\[\[([^\]]+)\]\]`)
)

// linkKind classifies a Markdown link target.
type linkKind int

const (
	linkExternal linkKind = iota
	linkAnchor
	linkNonMarkdown
	linkMarkdown
)

// classifyLink inspects a raw Markdown link target and returns its kind.
func classifyLink(target string) linkKind {
	t := strings.TrimSpace(target)
	if t == "" || strings.HasPrefix(t, "#") {
		return linkAnchor
	}
	lower := strings.ToLower(t)
	for _, prefix := range []string{"http://", "https://", "mailto:", "tel:", "ftp://", "ftps://"} {
		if strings.HasPrefix(lower, prefix) {
			return linkExternal
		}
	}
	noAnchor := stripAnchor(t)
	if noAnchor == "" {
		return linkAnchor
	}
	ext := strings.ToLower(path.Ext(noAnchor))
	if ext == ".md" || ext == ".markdown" {
		return linkMarkdown
	}
	return linkNonMarkdown
}

// stripAnchor removes a trailing "#anchor" from a target.
func stripAnchor(target string) string {
	if i := strings.Index(target, "#"); i >= 0 {
		return target[:i]
	}
	return target
}

// extractLinks returns Markdown links and Obsidian wikilinks from content,
// ignoring anything inside fenced code blocks, inline code, or image embeds.
func extractLinks(content string) ([]Link, []Wikilink) {
	stripped := codeFenceRE.ReplaceAllString(content, "")
	stripped = codeFenceRE2.ReplaceAllString(stripped, "")
	stripped = inlineCodeRE.ReplaceAllString(stripped, "")
	stripped = imageRE.ReplaceAllString(stripped, "")

	var links []Link
	for _, m := range mdLinkRE.FindAllStringSubmatch(stripped, -1) {
		if len(m) < 3 {
			continue
		}
		links = append(links, Link{Target: strings.TrimSpace(m[2])})
	}

	var wikis []Wikilink
	for _, m := range wikilinkRE.FindAllStringSubmatch(stripped, -1) {
		if len(m) < 2 {
			continue
		}
		wikis = append(wikis, Wikilink{Target: strings.TrimSpace(m[1])})
	}
	return links, wikis
}

// resolveMarkdownLink joins a Markdown link target relative to the directory
// of fromRelPath and returns the wiki-root-relative path. It returns ok=false
// for targets that escape the wiki root or resolve to empty.
func resolveMarkdownLink(fromRelPath, target string) (string, bool) {
	target = stripAnchor(strings.TrimSpace(target))
	if target == "" {
		return "", false
	}
	dir := path.Dir(fromRelPath)
	if dir == "." {
		dir = ""
	}
	joined := path.Join(dir, target)
	cleaned := path.Clean(joined)
	if cleaned == "." || cleaned == ".." || strings.HasPrefix(cleaned, "../") {
		return "", false
	}
	return cleaned, true
}
