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

	routes := buildRoutes(".")

	var buf bytes.Buffer

	gulog.Debug("Routes we got:\n- %s", strings.Join(routes, "\n- "))

	for _, route := range routes {

		if route != "/" {
			route = "/" + route
		} else { // means root and we need to remove / for root
			route = ""
		}

		tmpl, err := template.New(route).Parse(string(indexHtml))
		if err != nil {
			gulog.Error("Failed to create template for %s\nError: %s", route, err.Error())
			continue
		}
		err = tmpl.Execute(&buf, map[string]any{
			"wasm_url": route + "/main.wasm/",
		})
		if err != nil {
			gulog.Error("Failed to execute template: %s\nError: %s", route, err.Error())
			continue
		}

		output := buf.String()

		// http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		// 	http.Redirect(w, r, route+"/", http.StatusMovedPermanently)
		// })

		// // creating route for wasmfile for the route
		// http.HandleFunc(route+"/main.wasm", func(w http.ResponseWriter, r *http.Request) {
		// 	http.Redirect(w, r, route+"/main.wasm/", http.StatusMovedPermanently)
		// })
		wasm_file_path := filepath.Join(route, "main.wasm")
		if wasm, err := os.ReadFile(wasm_file_path); err == nil {
			http.HandleFunc(route+"/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(output))
			})
			http.HandleFunc(route+"/main.wasm/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(wasm)
			})
		} else {
			gulog.Error("Failed to Read the Wasm Code: %s", wasm_file_path)
			continue
		}

		gulog.Debug("Created Route for %s", route+"/")
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

	for _, entry := range entries {
		if entry.IsDir() {
			for _, r := range buildRoutes(filepath.Join(src, entry.Name())) {
				routes = append(routes, filepath.Join(entry.Name(), r))
			}
			continue
		}

		if entry.Name() == "main.wasm" && len(routes) < 1 {
			routes = append(routes, "/")
		}
	}

	return routes
}
