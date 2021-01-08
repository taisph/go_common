package commonversion

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"text/template"
)

var (
	Version string
	Branch  string
)

var tmpl = `
{{.name}} version {{.version}} ({{if .branch}}branch: {{.branch}}, {{end}}go: {{.go_version}})
`

func Print(name string) string {
	t := template.Must(template.New("version").Parse(tmpl))
	m := map[string]string{
		"name":       name,
		"version":    Version,
		"branch":     Branch,
		"go_version": runtime.Version(),
	}

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", m); err != nil {
		panic(err)
	}
	return strings.TrimSpace(buf.String())
}

func Info() string {
	return fmt.Sprintf("%s", Version)
}
