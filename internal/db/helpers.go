package db

import (
	"fmt"
	"log"
)

func CreateUser(username, password, role string) {
	log.Println("Adding User")
	addUserQuery := `INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3) RETURNING id`
	var id int
	var err error = DB.QueryRow(addUserQuery, username, password, role).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted record ID: %d\n", id)
}

func UpdateUser(userid, username, password, role string) {
	updateUserQuery := `
	UPDATE users
	SET username = $2, password = $3, role = $4
	WHERE id = $1
	RETURNING id, name;`
	var name string
	var id int
	var err error = DB.QueryRow(updateUserQuery, userid, username, password, role).Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, name)
}

func DeleteUser(userid string) {
	log.Println("Deleting User")
	deleteUserQuery := `DELETE FROM users WHERE id = $1`
	var err error
	_,err = DB.Exec(deleteUserQuery, userid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted record ID: %s\n", userid)
}