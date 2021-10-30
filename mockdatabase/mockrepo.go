{{define "mockrepo" -}}
// Package database contains structures and function for mock database access
// Generated code - do not modify it will be overwritten!!
// Time: {{.TimeStamp}}
package database

import (
	"fmt"
	model "{{.PackageName}}/model"
)

{{with .Entity}}

// {{.Name}}Repo is the interface for a {{.Name}} repository that will persist 
// and retrieve data and has to be implemented for concrete Databases 
// (e.g. db *sqlx.DB) or other respositories
type {{.Name | title}}Repo struct{
	data map[uint64]model.{{.Name | title}}
	count uint64
}

var {{.Name | lowercase}}repo *{{.Name | title}}Repo


// New{{.Name}}Repo creates a new instance of {{.Name}}Repo and initializes the 
// map and sets counter for ID to 1
func New{{.Name | title}}Repo() *{{.Name | title}}Repo {
	{{.Name | lowercase}}repo = new({{.Name | title}}Repo)
	{{.Name | lowercase}}repo.data = make(map[uint64]model.{{.Name | title}})
	{{.Name | lowercase}}repo.count = 1
	return {{.Name | lowercase}}repo
}


// Get queries a {{.Name | lowercase}} by id, throws an error when id is not found
func (repo {{.Name | title}}Repo) Get(id uint64) (*model.{{.Name | title}}, error) {
	value, ok := {{.Name | lowercase}}repo.data[id]
	if !ok {
		return nil, fmt.Errorf("get project with id %d, record not found", id)
	}
	return &value, nil
}

// GetAll returns all records ordered by the fields  with isLabel=true
func (repo {{.Name | title}}Repo) GetAll() (model.{{.Name | title}}List, error) {
	var list model.{{.Name | title}}List
	for _,value:=range {{.Name | lowercase}}repo.data {
		{{- range .Fields}}{{if eq .Kind "Lookup" "Child"}}
		if {{.Object | lowercase}},_:= {{.Object | lowercase}}repo.Get(value.{{.Name | title}}); {{.Object | lowercase}}!=nil {
			value.{{.Object}} = {{.Object | lowercase}}.Label()
		}
		{{- end}}{{end}}
		list = append(list,value)
	}
			
	return list, nil
}

// Delete deletes the {{.Name | lowercase}} with id, throws an error when id is not found
func (repo {{.Name | title}}Repo) Delete(id uint64) error {
	_, ok := {{.Name | lowercase}}repo.data[id]
	if !ok {
		return fmt.Errorf("delete project with id '%d', record not found", id)
	}

	delete({{.Name | lowercase}}repo.data, id)
	return nil
}

// Update updates all fields in the database table with data from *{{.Name | title}})
func (repo {{.Name| title}}Repo) Update({{.Name | lowercase}} *model.{{.Name | title}}) error {
	_, ok := {{.Name | lowercase}}repo.data[{{.Name | lowercase}}.ID]
	if !ok {
		return fmt.Errorf("update project with id '%d', record not found", {{.Name | lowercase}}.ID)
	}	
	{{.Name | lowercase}}repo.data[{{.Name | lowercase}}.ID] = *{{.Name | lowercase}}
	return nil
}

// Insert inserts a new record in the database table with data from *{{.Name | title}})
func (repo {{.Name | title}}Repo) Insert({{.Name | lowercase}} *model.{{.Name | title}}) error {
	{{.Name | lowercase}}repo.count++
	{{.Name | lowercase}}.ID ={{.Name | lowercase}}repo.count
	{{.Name | lowercase}}repo.data[{{.Name | lowercase}}repo.count] = *{{.Name | lowercase}}
	return nil
}

// GetLabels returns a map with the key id and the value of
// all fields tagged with isLabel=true and separated by a blank
func (repo {{.Name | title}}Repo) GetLabels() (model.Labels, error) {
	labels := make(model.Labels)
	for _, value := range {{.Name | lowercase}}repo.data {
		labels[value.ID] = value.Label()
	}
	return labels, nil
}

{{$name:=.Name}}
{{- range .Fields}}{{if eq .Kind "Parent"}}
// GetAll{{.Name | plural}}ByParentID returns a map with the key id and the value of
// all fields tagged with isLabel=true and separated by a blank
func (repo {{$name | title}}Repo) GetAll{{.Name | plural}}ByParentID(parentID uint64) (model.{{.Name | title}}List)	{
	list := model.{{.Name}}List{}
	{{.Name | plural | lowercase}}, err := {{.Name | lowercase}}repo.GetAll()
	if err!=nil {
		return list
	}
	for _, {{.Name | lowercase}} := range {{.Name | plural | lowercase}} {
		if {{.Name | lowercase}}.{{$name | title}}ID == parentID {
			list = append(list, {{.Name | lowercase}})
		}
	}
	return list
}			
{{- end}}{{end}}

{{end}}
{{end}}


