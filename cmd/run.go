package cmd

import (
	"bytes"
	_ "embed"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/vrianta/gulog"
)

type routes struct {
	routes   [][]byte
	html     []byte
	WasmPath []byte
}

//go:embed ui_templates/index.html
var indexHtml []byte

//go:embed ui_templates/wasm_exec.js
var wasmJs []byte

var runDir = "."

/*
Run the wasm application
to run we have to build the application and routes and urls
structure of the application
root directory -> every_folder is a route and the go file is the end path
for example we have a folder called test -> subpath -> test.go then our end route will be /test/subpath/test.go
*/
func Run() {

	http.HandleFunc("/wasm_exec.js", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/wasm_exec.js/", http.StatusMovedPermanently)
	})
	http.HandleFunc("/wasm_exec.js/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(wasmJs)
	})

	routes := buildRoutes(runDir)

	wd, _ := os.Getwd()
	os.Chdir(runDir)
	defer os.Chdir(wd)

	gulog.Debug("Routes we got:\n- %s", strings.Join(routes, "\n- "))

	for _, route := range routes {

		var buf bytes.Buffer

		wasmFilePath := filepath.Join(route, "main.wasm")

		httpRoute := route
		if httpRoute != "." {
			httpRoute = "/" + httpRoute
		} else {
			httpRoute = ""
		}

		tmpl, err := template.New(route).Parse(string(indexHtml))
		if err != nil {
			gulog.Error("Failed to create template for %s\nError: %s", route, err.Error())
			continue
		}
		err = tmpl.Execute(&buf, map[string]any{
			"wasm_url": httpRoute + "/main.wasm/",
		})
		if err != nil {
			gulog.Error("Failed to execute template: %s\nError: %s", route, err.Error())
			continue
		}

		output := buf.String()

		wasm, err := os.ReadFile(wasmFilePath)
		if err != nil {
			continue
		}

		html := output
		wasmData := string(wasm)
		routePath := httpRoute

		http.HandleFunc(routePath+"/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(html))
		})

		http.HandleFunc(routePath+"/main.wasm/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(wasmData))
		})

		gulog.Debug("Created Route for %s", httpRoute+"/")
	}
	gulog.Info("Startinng the server http://localhost:8080")

	// start the http server
	http.ListenAndServe(":8080", nil)

}

// Go thorugh each folder and files look for wasm files
// it will create http.HandleFunc where it will create indexHtml template
// pupulate the value of {{ .wasm_url }}
// wasm url will be /folder1/subf/main.wasm
func buildRoutes(src string) []string {
	var routes []string

	entries, err := os.ReadDir(src)
	if err != nil {
		return routes
	}

	hasWasm := false
	hasOtherFiles := false

	for _, entry := range entries {
		if entry.IsDir() {
			for _, r := range buildRoutes(filepath.Join(src, entry.Name())) {
				routes = append(routes, filepath.Join(entry.Name(), r))
			}
			continue
		}

		if entry.Name() == "main.wasm" {
			hasWasm = true
		} else if filepath.Ext(entry.Name()) != "" {
			hasOtherFiles = true
		}
	}

	if hasWasm {
		if src == runDir {
			routes = append(routes, ".")
		} else {
			routes = append(routes, "")
		}
	}
	if hasOtherFiles {

		prefix := "/"
		if src != "." {
			prefix = "/" + filepath.ToSlash(src) + "/"
		}

		http.Handle(
			prefix,
			http.StripPrefix(
				prefix,
				http.FileServer(http.Dir(src)),
			),
		)

		gulog.Debug("Created File Server for %s -> %s", prefix, src)

	}

	return routes
}
