package version

import "runtime"

var (
	// Name of the main daemon
	Name = "Testbed"

	// Version of the daemon
	Version = "0.1.0"

	// Description of the daemon
	Description = ""

	// Build stage
	Build = "-dev"

	// GitCommit to use to deploy
	GitCommit = "HEAD"
)

func BuildVersion() string {
	return Version + Build + " (" + GitCommit + ") " + runtime.GOOS + "/" + runtime.GOARCH
}

func FullVersion() string {
	return Name + "/" + BuildVersion()
}
