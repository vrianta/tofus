package main

import (
	gonfig "github.com/vrianta/gonfig/v1"
	"github.com/vrianta/gulog"
	"github.com/vrianta/tofus/cmd"
)

var Version = "v1.0"

var Args = gonfig.New[struct {
	Build   bool `arg:"build" default:"false"`
	Version bool `arg:"version" default:"false"`
}](true)

func main() {
	if Args.Version {
		gulog.Info("Version - %s", Version)
		return
	}
	if Args.Build {
		cmd.Build()
	}

	// http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`./wasm/`)))

	gulog.Wait()

}
