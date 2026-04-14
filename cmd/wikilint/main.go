// wikilint validates a wiki directory of Markdown pages and prints issues.
//
// Usage:
//
//	wikilint -wiki ./wiki
//	wikilint -version
//
// On success it prints "wikilint: OK" to stdout and exits 0.
// With -version it prints a one-line version descriptor and exits 0.
// On validation failures it prints one issue per line to stderr in the
// form "<relative-path>: <message>" and exits 1.
// On internal errors (missing directory, I/O failure) it exits 2.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/olegiv/llm-wiki-go/internal/version"
	"github.com/olegiv/llm-wiki-go/internal/wikilint"
)

func main() {
	wikiDir := flag.String("wiki", "./wiki", "path to the wiki directory")
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println(version.String())
		return
	}

	report, err := wikilint.Lint(*wikiDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wikilint: %v\n", err)
		os.Exit(2)
	}
	if !report.OK() {
		for _, issue := range report.Issues {
			fmt.Fprintf(os.Stderr, "%s: %s\n", issue.Path, issue.Message)
		}
		os.Exit(1)
	}
	fmt.Println("wikilint: OK")
}
