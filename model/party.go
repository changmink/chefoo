package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/changmink/shafoo/config"
)

type PartyForm struct {
	Name        string `json:"name"`
	MeetTime    string `json:"meetTime"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	TotalPeople int    `json:"totalPeople"`
}

func CreateParty(party PartyForm) int64 {
	db, err := sql.Open(config.C.DBType, config.C.DBPath)
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO party(name, meet_time, latitude, longitude, total_people, create_date, edit_date) values(?,?,?,?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(party.Name, party.MeetTime, party.Latitude, party.Longitude, party.TotalPeople, time.Now(), time.Now())
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("userId: ", id)

	db.Close()

	return id
}
