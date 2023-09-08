package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "mydb",
	}

	// Open a MySQL database connection
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Execute a simple query
	rows, err := db.Query("SELECT name, email FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Iterate through query results
	for rows.Next() {
		var name, email string
		if err := rows.Scan(&name, &email); err != nil {
			panic(err.Error())
		}
		fmt.Printf("Name: %s, Email: %s\n", name, email)
	}
}

/*

Install MySQL on MacOS:
$ brew install mysql

Start MySQL Service:
$ brew services start mysql

Log into MySQL:
$ mysql -u root

Create a DB:
mysql> CREATE DATABASE mydb;

Use the DB:
mysql> USE mydb;

Create the users table
mysql> CREATE TABLE users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

Populate the users table
mysql> INSERT INTO users (name, email) VALUES
    ('John Doe', 'john.doe@example.com'),
    ('Jane Smith', 'jane.smith@example.com'),
    ('Alice Johnson', 'alice.johnson@example.com');

*/
