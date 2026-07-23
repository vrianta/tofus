// Package main builds the Tofus CLI binary.
//
// Tofus scans a project `src/` directory, builds folders containing `main.go`
// into `main.wasm`, copies static assets into `build/`, and can run an
// HTTP server for the generated output.
package main

import (
	gonfig "github.com/vrianta/gonfig/v1"
	"github.com/vrianta/gulog"
	"github.com/vrianta/tofus/cmd"
)

// Version is the current CLI release version.
var Version = "v1.0"

// Args defines CLI flags for the Tofus tool.
var Args = gonfig.New[struct {
	Build   bool `arg:"build" default:"false"`   // build the project from src/ into build/
	Version bool `arg:"version" default:"false"` // print the current version
	Run     bool `arg:"run" default:"false"`     // run the Tofus HTTP server from build/
}](true)

// main executes the selected Tofus command.
func main() {
	if Args.Version {
		gulog.Info("Version - %s", Version)
	}
	if Args.Build {
		cmd.Build()
	}
	if Args.Run {
		cmd.Run()
	}

	// Keep the process alive until log output is flushed.
	gulog.Wait()
}
