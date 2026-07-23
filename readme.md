# Tofus

A Tool to Start server for wasm directories and build then

## Requirements

- The Source code should be inside the *src* directory this take *src* directory as root
- *wasm* will only be build if the folder has *main.go* in that folder

## How the Tool Works

### Build

```bash
# To build the current project
tofus --build
```
#### How Build Works

- It will create a new Directory called *build*
- it will scan the *src* directory
- It will look for main.go on each directory once it gets the main.go it replicate the same directory pattern and build the main.wasm file
- It will copy any file which do not have .go extension to the *build* folder with similer folder structure

#### Build Example

***Before Build***

src
|-assets
|--js
|---main.js
|--css
|---main.css
|-login
|--main.go
|-main.go

***After Build***

build
|-assets
|--js
|---main.js
|--css
|---main.css
|-login
|--main.wasm
|-main.wasm

### Run

You can also use this tool to server the server and create auto routing and file serve

```bash
# How to run
tofus --run # you have to be inside the build directory
```

To build and run at the same time
```bash
tofus --build --run # this will build the app and start running the it without the need to navigating to the build directory
```

#### How Routing Works

- It scans each and every folder once it found main.wasm it create route for it
- If it did not found any wasm file but got some other files in the build it will create a FileServer for it

#### Run Example

Build Directory

build
|-assets
|--js
|---main.js
|--css
|---main.css
|-login
|--main.wasm
|-main.wasm

The Routes
- /
- /login/

The File Servers
- /asset/css
- /asset/js