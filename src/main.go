package main

import (
	"log"

	dbsqlite "github.com/eoea/cgosqlite/src/pkg"
	sqltmpl "github.com/eoea/cgosqlite/src/templ"
)

func main() {
	employee := sqltmpl.Employee{
		Name:       "John",
		LastName:   "Doe",
		Department: "Sales",
		JobTitle:   "Assistant to the regional manager",
		Salary:     69420,
	}

	// Example 1:
	// Create the database: From a Go Template `./templates/createDBStmt.tmpl`
	// but in this case you can simply use the same contents from that file and
	// use a `.sql` extension.
	stmt, err := employee.CreateDBStmt()
	if err != nil {
		log.Fatalln(err)
	}
	err = dbsqlite.ExecuteSqliteFileOrStmt(stmt, false)
	if err != nil {
		log.Fatalln(err)
	}

	// Example 2:
	// Insert the data: From a Go Template `./templates/InsertStmt.tmpl`
	stmt, err = employee.InsertStmt()
	if err != nil {
		log.Fatalln(err)
	}
	err = dbsqlite.ExecuteSqliteFileOrStmt(stmt, false)
	if err != nil {
		log.Fatalln(err)
	}
}
