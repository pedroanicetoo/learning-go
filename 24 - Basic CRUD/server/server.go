package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"crud/database"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser insert user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro_r := ioutil.ReadAll(r.Body)
	if erro_r != nil {
		w.Write([]byte("Failed to read que request body"))
		return
	}

	var user user
	if erro_r = json.Unmarshal(requestBody, &user); erro_r != nil {
		w.Write([]byte("Error to convert the user to struct"))
		return
	}

	db, erro_r := database.Connect()
	if erro_r != nil {
		w.Write([]byte("Error to connect to database"))
		return
	}

	// PREPARE STATEMENT

	statement, erro_r := db.Prepare("insert into users (name, email) values (?, ?)")
	if erro_r != nil {
		w.Write([]byte("Error to create statement!"))
		return
	}
	defer statement.Close()

	insert, erro_r := statement.Exec(user.Name, user.Email)
	if erro_r != nil {
		w.Write([]byte("Error to execute the statement"))
		return
	}

	insertID, erro_r := insert.LastInsertId()
	if erro_r != nil {
		w.Write([]byte("Error to get insert ID"))
		return
	}

	// STATUS CODES
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserted with success! Id: %d", insertID)))
}
