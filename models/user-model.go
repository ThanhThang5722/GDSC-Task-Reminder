package models

import (
	"TaskReminder/pkg/const/tables"
	"TaskReminder/pkg/database"
	"errors"
	"fmt"
)

type User struct {
	userID   int
	username string
	password string
	email    string
}

func (u *User) SetUserName(value string) {
	u.username = value
}
func (u *User) SetPassword(value string) {
	u.password = value
}
func (u *User) SetUserID(value int) {
	u.userID = value
}
func (u *User) SetEmail(value string) {
	u.email = value
}

type UserDB struct {
	UserID   int
	Username string
	Password string
	Email    string
}

func NewUserFromDB(u UserDB) User {
	var newUser User
	newUser.SetPassword(u.Password)
	newUser.SetUserName(u.Username)
	newUser.SetUserID(u.UserID)
	return newUser
}

func (u *User) GetUserName() string {
	return u.username
}

func (u *User) GetUserID() int {
	return u.userID
}

func (u *User) GetPassword() string {
	return u.password
}
func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetOne(Property string, value interface{}) error {
	db := database.GetDbInstance()
	table := tables.UserTable
	query := fmt.Sprintf(`
			SELECT userID, username, password
			FROM %s
			WHERE %s = ?`,
		table, Property)

	resulst, err := db.Query(query, value)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var uDB UserDB
	for resulst.Next() {
		err := resulst.Scan(&uDB.UserID, &uDB.Username, &uDB.Password)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	*u = NewUserFromDB(uDB)
	return nil
}

func (u *User) Insert(uname, password, email string) error {
	db := database.GetDbInstance()
	query := `
		INSERT INTO users (username, password, email)
		VALUES (?,?,?)
	`
	_, err := db.Exec(query, uname, password, email)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Create(uname, hashPass, email string) error {
	var existedUser User
	if err := existedUser.GetOne("username", uname); err == nil && existedUser.GetUserName() != "" {
		return errors.New("username is existed")
	}
	if err := u.Insert(uname, hashPass, email); err != nil {
		return err
	}
	return nil
}
