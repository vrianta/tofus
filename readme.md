# Tofus

A Go tool and CLI for building and serving WebAssembly apps from a `src/` directory.

- Module path: `github.com/vrianta/tofus`
- Go version: `1.24.3`

## Quick Start

1. Put your source code under a `src/` directory.
2. Run `tofus --build` from the project root.
3. Run `tofus --run` from the generated `build/` directory, or use `tofus --build --run`.

## Features

- Build WebAssembly apps from Go packages under `src/`
- Copy non-`.go` files into the `build/` directory as static assets
- Auto-generate HTTP routes for `main.wasm` applications
- Serve static assets from generated directories

## Requirements

- All source code must be inside the `src/` directory.
- A WebAssembly build is created only for folders that contain `main.go`.

## Install

Recommended install:

```bash
go install github.com/vrianta/tofus@latest
```

Alternative install:

```bash
./install.sh
```

## Build

```bash
tofus --build
```

### What build does

- Creates a `build/` directory.
- Scans the `src/` directory recursively.
- For each folder that contains `main.go`, it builds `main.wasm` in the corresponding `build/` location.
- Copies all non-`.go` files into `build/` while preserving the directory structure.

### Build example

Before build:

```
src/
  assets/
    js/
      main.js
    css/
      main.css
  login/
    main.go
  main.go
```

After build:

```
build/
  assets/
    js/
      main.js
    css/
      main.css
  login/
    main.wasm
  main.wasm
```

## Run

### Run in `build/`

```bash
cd build
tofus --run
```

### Build and run together

```bash
tofus --build --run
```

This builds the project first and then runs the server without requiring you to change into the `build/` directory.

## Routing and asset serving

- The tool scans the generated `build/` tree.
- For each folder containing `main.wasm`, it creates a route for that folder.
- For folders with only static files, it creates a file server for those assets.

### Run example

Given this `build/` tree:

```
build/
  assets/
    js/
      main.js
    css/
      main.css
  login/
    main.wasm
  main.wasm
```

Generated routes:

- `/`
- `/login/`

Static file servers:

- `/assets/css/`
- `/assets/js/`

## Notes

- Only directories with `main.go` in `src/` become WebAssembly app routes.
- Non-`.go` files are copied as static assets.
- If `build/` already exists, it is updated with the new output.
- Nested `main.go` files are supported as long as they live inside `src/`.

## Troubleshooting

- If `src/` is missing or empty, the build produces no output.
- If a folder contains no `main.go` and only static files, those files are still copied as static assets.
- If a route is not created, verify that `main.go` exists in the expected `src/` folder.
