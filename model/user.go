package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/changmink/shafoo/config"
	_ "github.com/mattn/go-sqlite3"
)

type UserForm struct {
	Name  string
	Email string
	Auth  string
}

func AddUser(user UserForm) int64 {
	db, err := sql.Open(config.C.DBType, config.C.DBPath)
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO user(name, email, auth, create_date, edit_date) values(?,?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(user.Name, user.Email, user.Auth, time.Now(), time.Now())

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
