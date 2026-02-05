package models

import (
	"GO_REST_API/db"
	"GO_REST_API/utils"
	"errors"
	"fmt"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `
	INSERT INTO users(email, password) 
	VALUES(?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPass)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	user.ID = userId
	return err
}


func GetAllUsers() ([]User, error){
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next(){
		var user User
		err = rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (user User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPass string
	err := row.Scan(&retrievedPass) // in Scan have to be used &varName - without can be cause to error
	fmt.Println(retrievedPass)

	if err != nil {
		return errors.New("Credential errors!")
	}

	validPassword := utils.ComparePassword(user.Password, retrievedPass)
	fmt.Println(validPassword)


	if !validPassword {
		return errors.New("Credential errors!")
	}

	return nil

}
