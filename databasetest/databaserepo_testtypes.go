{{define "databaserepotesttypes"}}
{{- if eq .Kind "text"}}{{.Name | title}}: getText(15),{{end}}
{{- if eq .Kind "password"}}{{.Name | title}}: getText(15),{{end}}
{{- if eq .Kind "integer"}}{{- if ne .Name "ID"}}{{.Name | title}}: rand.Uint64(),{{end}}{{end}}
{{- if eq .Kind "number"}}{{.Name | title}}: rand.Float32(),{{end}}
{{- if eq .Kind "boolean"}}{{.Name | title}}:true,{{end}}
{{- if eq .Kind "email"}}{{.Name | title}}: getEmail(),{{end}}
{{- if eq .Kind "tel"}}	{{.Name | title}}: getText(12), {{end}}
{{- if eq .Kind "longtext"}}{{.Name | title}}: getText(50),{{end}}
{{- if eq .Kind "time"}}{{.Name | title}}:time.Now(),{{end}}
{{- if eq .Kind "lookup"}}{{.Name | title}}: 1,{{end}}
{{- if eq .Kind "child"}}{{.Name | title}}:1,{{end}}
{{- end}}