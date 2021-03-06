{{define "server" -}}
// Package view all view related structs and functions
// Generated code - do not modify it will be overwritten!!
// Time: {{.TimeStamp}}
package view

import (
	model "{{.Config.PackageName}}/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"text/template"
	"strconv"
	"io"
)

// FormatSettings holds all settings for formatting outputs like currencies
// time and date formats
type FormatSettings struct {
		CurrencySymbol    string
		DecimalSeparator  string
		ThousandSeparator string
		TimeFormat        string
		DateFormat        string
}

// Server main server based on echo server that holds all handlers and database
// access functions
type Server struct {
	*echo.Echo
	Env 	*model.Env

	Settings FormatSettings
}

// Template pointer to templates, necessary for echo template handling
type Template struct {
	templates *template.Template
}

// Render renders echo templates
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// NewServer creates a new main server based on echo that holds all handlers and database
// access functions, It has to be initialised with a populated 'Env' structure that holds
// the database access functions
func NewServer(env *model.Env) *Server {
	s := new(Server)
	s.Env = env
	s.Settings = FormatSettings{
		CurrencySymbol: "{{.Config.CurrencySymbol}}",
		DecimalSeparator: "{{.Config.DecimalSeparator}}",
		ThousandSeparator: "{{.Config.ThousandSeparator}}",
		TimeFormat: "{{.Config.TimeFormat}}",
		DateFormat: "{{.Config.DateFormat}}",		
	}
	s.Echo = echo.New()

	// assets will be loaded from /static directory as /assets/*
	s.Static("/static", "assets")

	// Instantiate a template registry with an array of template set
	// Ref: https://gist.github.com/rand99/808e6e9702c00ce64803d94abff65678
	renderer := &Template{
		templates: template.Must(template.ParseGlob("view/html/*.html")),
	}
	s.Renderer = renderer

	// Middleware
	s.Use(middleware.Logger())
	s.Use(middleware.Recover())

	// Routes
	s.GET("/", s.getDashboard) // Opens Dashboard

{{- range .Entities}}
	// routes for {{.Name}}
	s.GET("/{{.Name | plural | lowercase}}", s.List{{.Name | plural | title}})
	s.GET("/{{.Name | plural | lowercase}}/:id", s.Get{{.Name | title}})
	s.GET("/{{.Name | plural | lowercase}}/new", s.New{{.Name | title}})
	s.POST("/{{.Name | plural | lowercase}}", s.Create{{.Name | title}})
	s.POST("/{{.Name | plural | lowercase}}/:id", s.Update{{.Name | title}})
	s.POST("/{{.Name | plural | lowercase}}/:id/delete", s.Delete{{.Name | title}})
{{- end}}

	return s
}

// asUint returns a string as int or in case of an error 0
func asUint(s string) uint64 {
	if i, err := strconv.ParseUint(s, 10, 64); err == nil {
		return i
	}
	return 0
}


{{end}}