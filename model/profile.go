package model

import (
	"database/sql"

	"github.com/changmink/shafoo/config"
)

type ProfileInfo struct {
	Name  string
	Image string
	Score int
}

func GetProfileById(id string) ProfileInfo {
	db, err := sql.Open(config.C.DBType, config.C.DBPath)
	checkErr(err)

	rows, err := db.Query("SELECT name, image, score FROM profile WHERE user_id=" + id)
	checkErr(err)

	var profile ProfileInfo
	for rows.Next() {
		err := rows.Scan(&profile.Name, &profile.Image, &profile.Score)
		checkErr(err)
	}

	return profile
}
