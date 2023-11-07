package sqltmpl

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// This handles the fields that are then stored in the SQL database.
type Employee struct {
	Name       string
	LastName   string
	Department string
	JobTitle   string
	Salary     int
}

const (
	TEMPLATE_INSERTSTMT_PATH    = "templates/insertStmt.tmpl"
	TEMPLATE_CREATEDB_STMT_PATH = "templates/createDBStmt.tmpl"
)

func (e Employee) CreateDBStmt() (string, error) {
	var buf bytes.Buffer
	cwd, _ := os.Getwd()
	tmplFile := filepath.Join(cwd, TEMPLATE_CREATEDB_STMT_PATH)
	tmpl, err := template.New("createDBStmt.tmpl").ParseFiles(tmplFile)
	if err != nil {
		return "", fmt.Errorf("Error: in parsing template %s\n", err)
	}
	err = tmpl.Execute(&buf, e)
	if err != nil {
		return "", fmt.Errorf("Error: in template execute %s\n", err)
	}
	return buf.String(), nil
}

func (e Employee) InsertStmt() (string, error) {
	var buf bytes.Buffer
	cwd, _ := os.Getwd()
	tmplFile := filepath.Join(cwd, TEMPLATE_INSERTSTMT_PATH)
	tmpl, err := template.New("insertStmt.tmpl").ParseFiles(tmplFile)
	if err != nil {
		return "", fmt.Errorf("Error: in parsing template %s\n", err)
	}
	err = tmpl.Execute(&buf, e)
	if err != nil {
		return "", fmt.Errorf("Error: in template execute %s\n", err)
	}
	return buf.String(), nil
}
