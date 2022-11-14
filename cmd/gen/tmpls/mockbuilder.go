package tmpls

import (
	_ "embed"
)

//go:embed "mockbuilder.tmpl"
var MockBuilderTmpl string
