{{define "repo" -}}
// Package database contains structures and function for generic database access
// Generated code - do not modify it will be overwritten!!
// Time: {{.TimeStamp}}
package database

import (
	"github.com/jmoiron/sqlx"	
	"fmt"
	model "{{.PackageName}}/model"
)

{{with .Entity}}

// SQL statements
const (
	{{.Name | lowercase}}ByIDStatement = 	`SELECT * FROM {{.Name | lowercase | plural}} WHERE id=$1`
	{{.Name | lowercase}}AllStatement  = 	"SELECT * FROM {{.Name | lowercase| plural}} "+
		"ORDER BY {{range $index, $field:=.Fields}}{{if eq .IsLabel true}}{{if gt $index 0}},{{end}}{{$field.Name}} {{end}}{{end}} ASC"
	{{.Name | lowercase}}DeleteStatement = 	`DELETE FROM {{.Name | lowercase | plural}} WHERE id=$1)`
	{{.Name | lowercase}}LabelStatement  = 	"SELECT * FROM {{.Name | lowercase| plural}} "+
		"ORDER BY {{range $index, $field:=.Fields}}{{if eq .IsLabel true}}{{if gt $index 0}},{{end}}{{$field.Name}} {{end}}{{end}} ASC"
)


// {{.Name | lowercase}}UpdateStatement creates the UPDATE statement using templating technologies
func {{.Name | lowercase}}UpdateStatement() string {
	names:=[]string{ {{- range $index, $element := .Fields}}{{if ne .Name "ID"}}{{if ne .Kind "parent"}}
	{{- if gt $index 0}},{{end}} "{{$element.Name | lowercase}}" 
	{{- end}}{{end}}{{end}} }
	statement := "UPDATE {{.Name | lowercase}} SET"
	for i,name := range names {
		if i>0 {
			statement +=  ","
		}
		statement += " "+ name+"= $"+ string(i+2)
	}
	statement += " WHERE id= $1"
	
	return statement
}

// {{.Name | lowercase}}InsertStatement creates the INSERT statement using templating technologies
func {{.Name | lowercase}}InsertStatement() string {
	names:=[]string{ {{- range $index, $element := .Fields}}{{if ne .Name "ID"}}{{if ne .Kind "parent"}}
	{{- if gt $index 0}},{{end}} "{{$element.Name | lowercase}}" 
	{{- end}}{{end}}{{end}} }
	
	statement := `INSERT INTO {{.Name | lowercase}} ({{- range $index, $element := .Fields}}{{if ne .Name "ID"}}{{if ne .Kind "parent"}}
	{{- if gt $index 0}},{{end}} {{$element.Name | lowercase}} 
	{{- end}}{{end}}{{end}}) VALUES `
	for i := range names {
		if i>0 {
			statement += ", "
		}
		statement += string(i+1)
	}
	statement +=  " )"
	
	return statement
}

// {{.Name | title}}Repo is the interface for a {{.Name}} repository that will persist 
// and retrieve data and has to be implemented for concrete Databases 
// (e.g. db *sqlx.DB) or other respositories
type {{.Name | title}}Repo struct{
	// pointer to the global database
	DB *sqlx.DB
}

// New{{.Name | title}}Repo creates a new {{.Name | title}}Repo repository and assigns a database 
// connection to it
func New{{.Name | title}}Repo(db *sqlx.DB) *{{.Name | title}}Repo{
	repo := new({{.Name | title}}Repo)	
	repo.DB = db
	return repo
}

// Get queries a {{.Name | lowercase}} by id, throws an error when id is not found
func (repo {{.Name | title}}Repo) Get(id uint64) (*model.{{.Name | title}}, error) {
	{{.Name | lowercase}} := new(model.{{.Name | title}})
	if err := repo.DB.Get({{.Name | lowercase}}, {{.Name | lowercase}}ByIDStatement, id); err != nil {
		return nil, fmt.Errorf("get {{.Name | lowercase}} with id %d, %v", id, err)
	}
	return {{.Name | lowercase}}, nil
}

// GetAll returns all records ordered by the fields  with isLabel=true
func (repo {{.Name | title}}Repo) GetAll() (model.{{.Name | title}}List, error) {
	list := model.{{.Name | title}}List{}

	rows, err := repo.DB.Queryx({{.Name | lowercase}}AllStatement)
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		{{.Name | lowercase}} := new(model.{{.Name}})
		if err := rows.StructScan({{.Name | lowercase}}); err != nil {
			return nil, fmt.Errorf("parsing {{.Name | lowercase| plural}} struct, err %v", err)
		}
		list = append(list, *{{.Name | lowercase}})
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	return list, nil
}

// Delete deletes the {{.Name | lowercase}} with id, throws an error when id is not found
func (repo {{.Name | title}}Repo) Delete(id uint64) error {
	if _, err := repo.DB.Exec({{.Name | lowercase}}DeleteStatement, id); err != nil {
		return fmt.Errorf("delete {{.Name | lowercase}} with id %d, %v", id, err)
	}
	return nil
}

// Update updates all fields in the database table with data from *{{.Name}})
func (repo {{.Name | title}}Repo) Update({{.Name | lowercase}} *model.{{.Name | title}}) error {
		{{- $name := .Name  | title}}
	if _, err := repo.DB.Exec({{.Name | lowercase}}UpdateStatement(), {{.Name | lowercase}}.ID,{{- range $index, $element := .Fields}}{{if ne $element.Name "ID"}}{{if ne .Kind "parent"}}{{if gt $index 0}}, {{end}}{{$name | lowercase}}.{{$element.Name}}{{end}}{{end}}{{end}} ); err != nil {
		return fmt.Errorf("update {{.Name | lowercase| plural}}, %v", err)
	}
	return nil
}

// Insert inserts a new record in the database table with data from *{{.Name}})
func (repo {{.Name | title}}Repo) Insert({{.Name | lowercase}} *model.{{.Name | title}}) error {
	{{- $name := .Name  | title}}
	if _, err := repo.DB.Exec({{.Name | lowercase}}InsertStatement(), {{- range $index, $element := .Fields}}{{if ne $element.Name "ID"}}{{if ne .Kind "parent"}}{{if gt $index 0}}, {{end}}{{$name | lowercase}}.{{$element.Name}}{{end}}{{end}}{{end}}); err != nil {
		return fmt.Errorf("insert {{.Name | lowercase| plural}}, %v", err)
	}
	return nil
}

// GetLabels returns a map with the key id and the value of
// all fields tagged with isLabel=true and separated by a blank
func (repo {{.Name | title}}Repo) GetLabels() (model.Labels, error) {
	l := make(model.Labels)

	rows, err := repo.DB.Queryx({{.Name | lowercase}}LabelStatement)
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		{{.Name | lowercase}} := new(model.{{.Name | title}})
		if err := rows.StructScan({{.Name | lowercase}}); err != nil {
			return nil, fmt.Errorf("parsing {{.Name | lowercase| plural}} struct, err %v", err)
		}
		{{$name:= .Name | title}}
		label := fmt.Sprintf("{{range .Fields}}{{if eq .IsLabel true}}%s {{end}}{{end}}"{{range .Fields}}{{if eq .IsLabel true}},{{$name | lowercase}}.{{.Name}} {{end}}{{end}})
		l[{{.Name | lowercase}}.ID] = label
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	return l, nil
}

{{$name:=.Name | title}}
{{- range .Fields}}{{if eq .Kind "parent"}}
// GetAll{{$name | plural | title}}ForParentID returns a map with the key id and the value of
// all fields tagged with isLabel=true and separated by a blank
func (repo {{$name}}Repo) GetAll{{.Name | plural | title}}ByParentID(parentID uint64) (model.{{.Name | title}}List, error)	{
	list := model.{{.Name | title}}List{}

	query:= fmt.Sprintf("SELECT * FROM {{.Name | lowercase| plural}} WHERE id=%d", parentID)
	rows, err := repo.DB.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("selecting {{.Name | lowercase| plural}} with id: '%d' from database, %v",parentID, err)
	}
	
	for rows.Next() {
		{{.Name | lowercase}} := new(model.{{.Name | title}})
		if err := rows.StructScan({{.Name | lowercase}}); err != nil {
			return nil, fmt.Errorf("parsing {{.Name | lowercase| plural}} struct, %v", err)
		}
		list = append(list, *{{.Name | lowercase}})
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	return list, nil
}			
{{- end}}{{end}}

{{end}}
{{end}}


