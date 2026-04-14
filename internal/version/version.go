// Package version reports the wikilint binary's version string.
package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
)

const binaryName = "wikilint"

// String returns a one-line version descriptor suitable for
// "wikilint -version" output. It derives version, commit, and Go
// toolchain information from runtime/debug.ReadBuildInfo.
func String() string {
	info, ok := debug.ReadBuildInfo()
	return format(info, ok, runtime.GOOS, runtime.GOARCH)
}

func format(info *debug.BuildInfo, ok bool, goos, goarch string) string {
	goVer := runtime.Version()
	if ok && info != nil && info.GoVersion != "" {
		goVer = info.GoVersion
	}

	if !ok || info == nil {
		return fmt.Sprintf("%s dev (%s, %s/%s)", binaryName, goVer, goos, goarch)
	}

	revision := ""
	for _, s := range info.Settings {
		if s.Key == "vcs.revision" {
			revision = s.Value
			break
		}
	}
	if len(revision) > 7 {
		revision = revision[:7]
	}

	ver := info.Main.Version
	hasSemver := strings.HasPrefix(ver, "v")

	switch {
	case hasSemver && revision != "":
		return fmt.Sprintf("%s %s (commit %s, %s, %s/%s)", binaryName, ver, revision, goVer, goos, goarch)
	case hasSemver:
		return fmt.Sprintf("%s %s (%s, %s/%s)", binaryName, ver, goVer, goos, goarch)
	case revision != "":
		return fmt.Sprintf("%s %s (%s, %s/%s)", binaryName, revision, goVer, goos, goarch)
	default:
		return fmt.Sprintf("%s dev (%s, %s/%s)", binaryName, goVer, goos, goarch)
	}
}
