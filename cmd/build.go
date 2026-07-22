package cmd

import (
	_ "embed"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/vrianta/gulog"
)

type Routes struct {
	route string // http route
}

//go:embed ui_templates/index.html
var indexHtml []byte

//go:embed ui_templates/wasm_exec.js
var wasmJs []byte

var src_folder_path = filepath.Join(".", "src")
var ViewFolder string = "views"

/*
Scans entire directory
*/
func Build() {

	gulog.Info("Starting the Build Process")
	buildDir := filepath.Join(".", "build")
	gulog.Info("Creating Build Directory")
	if err := os.Mkdir(buildDir, 0777); err != nil {
		gulog.Error("Failed to Create build Directory")
		os.Exit(-1)
	}

	// routes := buildDir
	// scan entire views directory

	if _, err := os.ReadDir(src_folder_path); err != nil {
		gulog.Error("Folder Structure is wrong src folder not present")
		os.Exit(-1)
	}
	err := create_build_files(src_folder_path, buildDir)
	if err != nil {
		panic(err)
	}

}

// loop through folder and paths and create folder and wasm files to build
// src what to scan and dst is what to create or where to create
func create_build_files(src string, dst string) error {

	err := filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			// build the route to create the folders from routes
			// gulog.Info("Creating Folder - %s", d)
			// os.Mkdir(string(routes), 0777)

			// // resetting the route to build
			// routes = buildDir
			// gulog.Info("Not a dir it's a file time to creat the route")
			// return nil
		}

		// folder I have to create
		f_to_create := filepath.Join(dst, d.Name())
		gulog.Info("Creating Tofus Module %s", f_to_create)
		os.Mkdir(f_to_create, 0777)
		gulog.Info("-- Done %s", f_to_create)

		create_build_files(filepath.Join(src, d.Name()), f_to_create)

		return nil
	})
	return err
}
