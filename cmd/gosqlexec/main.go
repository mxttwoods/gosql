package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// create a database object which can be
	// used to connect with database.
	db, err := sql.Open("mysql", "root:passwd@tcp(0.0.0.0:3306)/user")

	// handle error, if any.
	if err != nil {
		panic(err)
	}

	// Here a SQL query is used to return all
	// the data from the table user.
	result, err := db.Query("SELECT * FROM user")

	// handle error
	if err != nil {
		panic(err)
	}

	// the result object has a method called Next,
	// which is used to iterate through all returned rows.
	for result.Next() {

		var id int
		var name string

		// The result object provided Scan method
		// to read row data, Scan returns error,
		// if any. Here we read id and name returned.
		err = result.Scan(&id, &name)

		// handle error
		if err != nil {
			panic(err)
		}

		fmt.Printf("Id: %d Name: %s\n", id, name)
	}

	// database object has a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
}
