#!/usr/bin/env bash

##### Description ####
# Setup script for creating a new WASM project structure
######################
set -e

PROJECT_NAME="${1:-wasnui}"

mkdir -p "$PROJECT_NAME"/{
app,
widget,
material,
layout,
state,
router,
browser,
dom,
wasm
}

touch 
"$PROJECT_NAME/app/app.go" 
"$PROJECT_NAME/app/header.go" 
"$PROJECT_NAME/app/footer.go" 
"$PROJECT_NAME/widget/widget.go" 
"$PROJECT_NAME/widget/element.go" 
"$PROJECT_NAME/widget/context.go" 
"$PROJECT_NAME/material/button.go" 
"$PROJECT_NAME/material/text.go" 
"$PROJECT_NAME/material/column.go" 
"$PROJECT_NAME/material/row.go" 
"$PROJECT_NAME/layout/flex.go" 
"$PROJECT_NAME/layout/stack.go" 
"$PROJECT_NAME/layout/spacer.go" 
"$PROJECT_NAME/state/state.go" 
"$PROJECT_NAME/state/notifier.go" 
"$PROJECT_NAME/router/router.go" 
"$PROJECT_NAME/router/route.go" 
"$PROJECT_NAME/browser/viewport.go" 
"$PROJECT_NAME/browser/language.go" 
"$PROJECT_NAME/browser/charset.go" 
"$PROJECT_NAME/browser/meta.go" 
"$PROJECT_NAME/dom/node.go" 
"$PROJECT_NAME/dom/diff.go" 
"$PROJECT_NAME/dom/patch.go" 
"$PROJECT_NAME/wasm/js.go"

echo "Created project structure:"
tree "$PROJECT_NAME" 2>/dev/null || find "$PROJECT_NAME"
