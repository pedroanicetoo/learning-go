package main

// To test: Create a mysql database named 'devbook' with user 'golang' and password 'golang'

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConnection := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro_r := sql.Open("mysql", stringConnection)

	if erro_r != nil {
		log.Fatal(erro_r)
	}
	defer db.Close()

	if erro_r = db.Ping(); erro_r != nil {
		log.Fatal(erro_r)
	}

	fmt.Println("Open Connection")
	lines, erro_r := db.Query("select * from users")
	if erro_r != nil {
		log.Fatal(erro_r)
	}
	defer lines.Close()
	fmt.Println(lines)
}
