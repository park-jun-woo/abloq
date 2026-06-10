// Package template embeds the abloq blog template payload (degit-style copy
// source for `abloq init`). The payload lives under files/ so the embedded
// paths can be re-rooted with fs.Sub before copying.
package template

import (
	"embed"
	"io/fs"
)

//go:embed all:files
var payload embed.FS

// Files returns the template payload rooted at the blog-root level: paths are
// exactly what `abloq init` writes into the new blog directory. The fs.Sub
// error is impossible for an embedded directory and is deliberately dropped.
func Files() fs.FS {
	sub, _ := fs.Sub(payload, "files")
	return sub
}
