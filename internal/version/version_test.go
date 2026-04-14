package version

import (
	"runtime/debug"
	"testing"
)

func TestFormat(t *testing.T) {
	cases := []struct {
		name string
		info *debug.BuildInfo
		ok   bool
		want string
	}{
		{
			name: "no build info",
			info: nil,
			ok:   false,
			want: "wikilint dev (go1.26.0, linux/amd64)",
		},
		{
			name: "build info without vcs",
			info: &debug.BuildInfo{
				GoVersion: "go1.26.0",
				Main:      debug.Module{Version: "(devel)"},
			},
			ok:   true,
			want: "wikilint dev (go1.26.0, linux/amd64)",
		},
		{
			name: "revision only",
			info: &debug.BuildInfo{
				GoVersion: "go1.26.0",
				Main:      debug.Module{Version: "(devel)"},
				Settings: []debug.BuildSetting{
					{Key: "vcs.revision", Value: "abcdef1234567"},
				},
			},
			ok:   true,
			want: "wikilint abcdef1 (go1.26.0, linux/amd64)",
		},
		{
			name: "semver and revision",
			info: &debug.BuildInfo{
				GoVersion: "go1.26.0",
				Main:      debug.Module{Version: "v0.1.0"},
				Settings: []debug.BuildSetting{
					{Key: "vcs.revision", Value: "abcdef1234567"},
				},
			},
			ok:   true,
			want: "wikilint v0.1.0 (commit abcdef1, go1.26.0, linux/amd64)",
		},
		{
			name: "semver only",
			info: &debug.BuildInfo{
				GoVersion: "go1.26.0",
				Main:      debug.Module{Version: "v0.1.0"},
			},
			ok:   true,
			want: "wikilint v0.1.0 (go1.26.0, linux/amd64)",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := format(tc.info, tc.ok, "linux", "amd64")
			if got != tc.want {
				t.Errorf("format() = %q, want %q", got, tc.want)
			}
		})
	}
}
