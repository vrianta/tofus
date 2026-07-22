package cmd

import (
	"io/fs"
	"path/filepath"

	"github.com/vrianta/gulog"
)

type Routes struct {
	route string // http route
}

var ViewFolder string = "views"

/*
Scans entire directory
*/
func Build() {

	if !Config.Build {
		return
	}
	var route []byte = []byte{'/'}
	// scan entire views directory
	err := filepath.WalkDir("views", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			gulog.Info("Not a dir it's a file time to creat the route")
			return nil
		}

		route = append(route)

		return nil
	})

	if err != nil {
		panic(err)
	}

}
