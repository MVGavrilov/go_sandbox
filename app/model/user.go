package model

import (
	"log"
	"time"

	"mysrvr/app/common/database"
)

type User struct {
	ID        uint32    `db:"id"`
	Name      string    `db:"name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Status    uint8     `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}

var userList []User

func UserCreate(firstName, lastName, email, password string) error {
	var err error
	_, err = database.Handler.Exec("INSERT INTO users (name, last_name, email, password, status, created_at) VALUES ($1, $2, $3, $4, 0, NOW())", firstName, lastName, email, password)
	if err != nil {
		log.Default().Println(err)
	}
	return err
}

func UserUpdate(id uint32, firstName, lastName, email, password string) error {
	var err error
	_, err = database.Handler.Exec("UPDATE users SET name=$1, last_name=$2, email=$3, password=$4 WHERE id=$5", firstName, lastName, email, password, id)
	if err != nil {
		log.Default().Println(err)
	}
	return err
}

func UserDelete(id uint32) error {
	var err error
	_, err = database.Handler.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		log.Default().Println(err)
	}
	return err
}

func UserList() ([]User, error) {
	var err error
	err = database.Handler.Select(&userList, "SELECT * FROM users")
	if err != nil {
		log.Default().Println(err)
	}
	return userList, err
}

func UserByID(id uint32) (User, error) {
	var result User
	err := database.Handler.Get(&result, "SELECT * FROM users WHERE id=$1 LIMIT 1", id)
	if err != nil {
		log.Default().Println(err)
	}
	return result, err
}

func UserByEmail(email string) (User, error) {
	var result User
	err := database.Handler.Get(&result, "SELECT * FROM users WHERE email=$1 LIMIT 1", email)
	if err != nil {
		log.Default().Println(err)
	}
	return result, err
}

