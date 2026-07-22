package cmd

import gonfig "github.com/vrianta/gonfig/v1"

var Config = gonfig.New[struct {
	Build bool `arg:"build" default:"false"`
}](true)
