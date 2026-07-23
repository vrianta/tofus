// Package cmd contains shared build and run commands for the Tofus CLI.
package cmd

import (
	_ "embed"
	"io"
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

// Build scans the local src/ directory and creates the build/ output tree.
// It compiles directories containing main.go into WebAssembly binaries and
// copies other assets into the build/ folder.
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

// createBuildFolders recursively walks the source directory tree.
// It creates matching directories under dst, copies non-Go files, and
// builds WebAssembly outputs for directories containing main.go.
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
			} else if filepath.Ext(entry.Name()) != ".go" { // copy the non go files to the build folder
				srcFile := filepath.Join(src, entry.Name())
				dstFile := filepath.Join(dst, entry.Name())

				gulog.Info("Copying file %s to %s", srcFile, dstFile)

				if err := copyFile(srcFile, dstFile); err != nil {
					gulog.Error("Failed to copy file %s to %s\nError: %s", srcFile, dstFile, err.Error())
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

// copyFile copies a single file from src to dst preserving file mode.
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	info, err := in.Stat()
	if err != nil {
		return err
	}

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

// BuildWasm compiles the Go package in src to WebAssembly and writes the
// output binary as main.wasm into dst.
func BuildWasm(src, dst string) error {
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
