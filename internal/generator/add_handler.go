package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Aro-M/go-micro-gen/internal/config"
	"golang.org/x/mod/modfile"
)

// AddHandler scaffold a new handler endpoint.
func AddHandler(name, route string) error {
	// First, dynamically discover the module name by parsing go.mod
	modBytes, err := os.ReadFile("go.mod")
	if err != nil {
		return fmt.Errorf("could not read go.mod (are you in the project root?): %w", err)
	}

	modPath := modfile.ModulePath(modBytes)
	if modPath == "" {
		return fmt.Errorf("could not parse module path from go.mod")
	}

	cfg := &config.ServiceConfig{
		ServiceName: name,
		ModulePath:  modPath,
	}

	tmplPath := "add/handler.go.tmpl"
	content, err := fs.ReadFile(templateFS, tmplPath)
	if err != nil {
		return fmt.Errorf("read template %s: %w", tmplPath, err)
	}

	tmpl, err := template.New("handler").
		Funcs(templateFuncs()).
		Parse(string(content))
	if err != nil {
		return fmt.Errorf("parse template %s: %w", tmplPath, err)
	}

	// Prepare data
	data := struct {
		Config *config.ServiceConfig
		Name   string
		Route  string
	}{
		Config: cfg,
		Name:   name,
		Route:  route,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	// Format code
	formattedCode, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("format generated code: %w", err)
	}

	// Save to internal/transport/httpx/{name}_handler.go
	fileName := fmt.Sprintf("%s_handler.go", strings.ToLower(name))
	outPath := filepath.Join("internal", "transport", "httpx", fileName)

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return fmt.Errorf("create dir for %s: %w", outPath, err)
	}

	if err := os.WriteFile(outPath, formattedCode, 0644); err != nil {
		return fmt.Errorf("write file %s: %w", outPath, err)
	}

	fmt.Printf("  ✔ %s\n", outPath)
	return nil
}
