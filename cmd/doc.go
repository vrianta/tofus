// Package cmd provides build and run commands used by the Tofus CLI.
//
// Tofus scans a src/ directory and generates a build/ directory for
// WebAssembly applications. It compiles each package containing main.go
// into main.wasm, copies static assets, and serves the generated output
// with automatic route handling.
package cmd
