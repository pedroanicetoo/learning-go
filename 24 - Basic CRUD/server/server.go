package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"crud/database"

	"github.com/gorilla/mux"
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

// SearchUsers get a specific user from database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	db, erro_r := database.Connect()
	if erro_r != nil {
		w.Write([]byte("Error to connect to database"))
		return
	}
	defer db.Close()

	lines, erro_r := db.Query("select * from users")
	if erro_r != nil {
		w.Write([]byte("Error on select users"))
		return
	}
	defer lines.Close()

	var users []user

	for lines.Next() {
		var user user

		if erro_r := lines.Scan(&user.ID, &user.Name, &user.Email); erro_r != nil {
			w.Write([]byte("Error on user scan"))
			return
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	if erro_r := json.NewEncoder(w).Encode(users); erro_r != nil {
		w.Write([]byte("Error on convert users to JSON"))
		return
	}
}

// SearchUser get a specific user from database
func SearchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, erro_r := strconv.ParseUint(params["id"], 10, 32)
	if erro_r != nil {
		w.Write([]byte("Error on convert param to integer"))
		return
	}

	db, erro_r := database.Connect()
	if erro_r != nil {
		w.Write([]byte("Error to connect to database"))
		return
	}

	line, erro_r := db.Query("select * from users where id = ?", ID)
	if erro_r != nil {
		w.Write([]byte("Error on select user"))
		return
	}

	var user user
	if line.Next() {
		if erro_r := line.Scan(&user.ID, &user.Name, &user.Email); erro_r != nil {
			w.Write([]byte("Error on user scan"))
			return
		}
	}

	if erro_r := json.NewEncoder(w).Encode(user); erro_r != nil {
		w.Write([]byte("Error on convert user to JSON"))
		return
	}

}
