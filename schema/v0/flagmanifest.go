// Package flagmanifest embeds the flag manifest into a code module.
package flagmanifest

import _ "embed"

// Schema contains the embedded flag manifest schema.
//
//go:embed flag_manifest.json
var Schema string

// SchemaPath provides the current path and version of flag manifest.
const SchemaPath = "github.com/open-feature/cli/schema/v0/flag_manifest.json"
