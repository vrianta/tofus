package cmd

import (
	_ "embed"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/vrianta/gulog"
)

type Routes struct {
	route string // http route
}

var src_folder_path = filepath.Join(".", "src")
var ViewFolder string = "views"

/*
Scans entire directory
*/
func Build() {
	wd, _ := os.Getwd()
	gulog.Info("Starting the Build Process")
	buildDir := filepath.Join(wd, "build")
	runDir = buildDir
	gulog.Info("Creating Build Directory")

	// check if DIR exists
	if err := os.Mkdir(buildDir, 0777); err != nil && !os.IsExist(err) {
		gulog.Error("Failed to Create build Directory\nError: %s", err.Error())
		return
	}

	if _, err := os.ReadDir(src_folder_path); err != nil && os.IsExist(err) {
		gulog.Error("Folder Structure is wrong src folder not present")
		return
	}
	gulog.Info("src folder found")
	err := createBuildFolders(src_folder_path, buildDir)
	if err != nil {
		gulog.Error("%s", err.Error())
		return
	}

}

// loop through folder and paths and create folder and wasm files to build
// src what to scan and dst is what to create or where to create
func createBuildFolders(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// Create the current destination folder.
	if err := os.MkdirAll(dst, 0777); err != nil {
		return err
	}

	// wasmsToBuild := map[string]struct {
	// 	src string
	// 	dst string
	// }{}
	for _, entry := range entries {
		if !entry.IsDir() {

			// if _, ok := wasmsToBuild[src]; !ok {
			// 	wasmsToBuild[src] = struct {
			// 		src string
			// 		dst string
			// 	}{
			// 		src: src,
			// 		dst: dst,
			// 	}
			// }
			if entry.Name() == "main.go" {
				if err := BuildWasm(src, dst); err != nil {
					gulog.Error("Failed to build the wasm of %s\nError: %s", src, err.Error())
				}
			}

			continue
		}

		srcDir := filepath.Join(src, entry.Name())
		dstDir := filepath.Join(dst, entry.Name())

		gulog.Info("Creating module %s", dstDir)

		if err := os.Mkdir(dstDir, 0777); err != nil && !os.IsExist(err) {
			return err
		}

		// Recursively process children.
		if err := createBuildFolders(srcDir, dstDir); err != nil {
			return err
		}
	}

	return nil
}

// this function will build wasm GOOS=js GOARCH=wasm go build -o main.wasm in src
//
//	in the destination directory
func BuildWasm(src, dst string) error {
	// Ensure the destination exists.
	gulog.Info("Building Wasm----\nSrc: %s\ndst: %s", src, dst)

	out := filepath.Join(dst, "main.wasm")

	cmd := exec.Command(
		"go",
		"build",
		"-o",
		out,
	)

	// Build from the source directory.
	cmd.Dir = src

	// Set the WASM environment.
	cmd.Env = append(os.Environ(),
		"GOOS=js",
		"GOARCH=wasm",
	)

	// Show compiler output.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
