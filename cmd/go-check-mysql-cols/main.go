package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DBname string = "DB" // set this to your db (shard-0)
var Colname string = "objectId" // set this to what your looking for

func main() {
	// Create a connection to the database
	db, err := sql.Open("mysql", "root:@(localhost)/"+DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Fetch the list of tables
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatal(err)
		}

		// Check for 'ColName' column in the table
		colQuery := fmt.Sprintf("SHOW COLUMNS FROM %s LIKE '%s'", tableName, Colname)
		rows, err := db.Query(colQuery)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		if rows.Next() {
			fmt.Printf("Table '%s' is OK.\n", tableName)
		} else {
			fmt.Printf("Table '%s' does NOT have the column '%s'.\n", tableName, Colname)
		}

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
