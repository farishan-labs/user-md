package buildinfo

import "runtime/debug"

// Version is overridden via -ldflags "-X .../buildinfo.Version=...".
var Version = "dev"

func init() {
	if Version != "dev" {
		return
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		Version = info.Main.Version
	}
}
