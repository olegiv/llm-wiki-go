package wikilint

import (
	"regexp"
	"sort"
	"strings"
)

var slugNonAlnumRE = regexp.MustCompile(`[^a-z0-9]+`)

// normalizeSlug lowercases s and replaces runs of non-alphanumeric characters
// with a single hyphen, trimming leading/trailing hyphens.
func normalizeSlug(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = slugNonAlnumRE.ReplaceAllString(s, "-")
	return strings.Trim(s, "-")
}

// findDuplicates groups pages by normalized title slug and returns the groups
// that contain more than one page. Keys are the duplicated slugs; values are
// the sorted wiki-root-relative page paths.
func findDuplicates(pages map[string]*Page) map[string][]string {
	groups := map[string][]string{}
	for rel, p := range pages {
		if p.Title == "" {
			continue
		}
		slug := normalizeSlug(p.Title)
		if slug == "" {
			continue
		}
		groups[slug] = append(groups[slug], rel)
	}
	out := map[string][]string{}
	for slug, rels := range groups {
		if len(rels) > 1 {
			sort.Strings(rels)
			out[slug] = rels
		}
	}
	return out
}
