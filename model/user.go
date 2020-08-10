package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/changmink/shafoo/config"
	_ "github.com/mattn/go-sqlite3"
)

type UserForm struct {
	Email string
	Auth  string
}
type LoginUser struct {
	Email string
	Auth  string
}

func AddUser(user UserForm) int64 {
	db, err := sql.Open(config.C.DBType, config.C.DBPath)
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO user(email, auth, create_date, edit_date) values(?,?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(user.Email, user.Auth, time.Now(), time.Now())

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("userId: ", id)

	db.Close()

	return id
}

func checkErr(err error) {
	if err != nil {
		panic(err) //500에러로 떨어짐
	}
}

func ExistUser(user LoginUser) bool {
	db, err := sql.Open(config.C.DBType, config.C.DBPath)
	checkErr(err)

	rows, err := db.Query("SELECT auth FROM user WHERE email='" + user.Email + "'")
	checkErr(err)

	var auth string

	for rows.Next() {
		err = rows.Scan(&auth)
		checkErr(err)
	}

	db.Close()

	if auth == user.Auth {
		return true
	}
	return false
}
