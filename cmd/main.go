package main

import (
	"embed"
	"io/fs"
	"log/slog"
	"os"

	"github.com/go-micro-gen/go-micro-gen/internal/cli"
	"github.com/go-micro-gen/go-micro-gen/internal/generator"
)

//go:embed all:templates
var embeddedTemplates embed.FS

func main() {
	// Trim the top-level "templates" dir so generator paths are relative to it
	sub, err := fs.Sub(embeddedTemplates, "templates")
	if err != nil {
		slog.Error("failed to sub templates FS", "error", err)
		os.Exit(1)
	}
	generator.SetTemplateFS(sub)

	if err := cli.Execute(); err != nil {
		slog.Error("go-micro-gen failed", "error", err)
		os.Exit(1)
	}
}
