{{define "mockmaintest" -}}
// Package database contains structures and function for mock database access
// Generated code - do not modify it will be overwritten!!
// Time: {{.TimeStamp}}
package database

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
{{range .Entities}}
	{{.Name | lowercase}}db = New{{.Name | title}}Repo()
{{- end}}

	os.Exit(m.Run())
}

{{end}}





