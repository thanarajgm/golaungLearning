package main //Package declaration

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Import Packages

func main() { //function declaration
	fmt.Println("Database connection module")

	db, err := sql.Open("mysql", "root:November@2025@tcp(localhost:3306)/testdb")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()
	fmt.Println("Successfully connected to the database.")

	query, err := db.Query("INSERT into users VALUES('ABCD')")

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer query.Close()
	fmt.Println("Query executed successfully.")
}
