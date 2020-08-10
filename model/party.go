package model

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"strconv"
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

type PartyInfo struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	MeetTime      string `json:"meetTime"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	TotalPeople   int    `json:"totalPeople"`
	CurrentPeople int    `json:"currentPeople"`
}

func CreateParty(party PartyForm) int64 {
	db, err := sql.Open(config.C.DBType, config.C.DBPath)
	checkErr(err)
	defer db.Close()

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

func SearchParties(latitude string, longitude string, dist float64) []PartyInfo {
	db, err := sql.Open(config.C.DBType, config.C.DBPath)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT id, name, meet_time, latitude, longitude, total_people, current_people FROM party")

	parties := []PartyInfo{}
	for rows.Next() {
		var party PartyInfo
		err := rows.Scan(&party.Id, &party.Name, &party.MeetTime, &party.Latitude, &party.Longitude, &party.TotalPeople, &party.CurrentPeople)
		checkErr(err)
		if calDistence(party.Latitude, party.Longitude, latitude, longitude) <= dist {
			parties = append(parties, party)
		}
	}

	return parties
}

func calDistence(latitude1 string, longitude1 string, latitude2 string, longitude2 string) float64 {
	lat1, err := strconv.ParseFloat(latitude1, 4)
	if err != nil {
		panic(err)
	}
	lon1, err := strconv.ParseFloat(longitude1, 4)
	if err != nil {
		panic(err)
	}
	lat2, err := strconv.ParseFloat(latitude2, 4)
	if err != nil {
		panic(err)
	}
	lon2, err := strconv.ParseFloat(longitude2, 4)
	if err != nil {
		panic(err)
	}

	theta := lon1 - lon2
	dist := math.Sin(degreeToRadian(lat1))*math.Sin(degreeToRadian(lat2)) + math.Cos(degreeToRadian(lat1))*math.Cos(degreeToRadian(lat2))*math.Cos(degreeToRadian(theta))
	dist = math.Acos(dist)
	dist = radianToDegree(dist)
	dist = dist * 60 * 1.515

	dist = dist * 1.609344

	return dist
}
func degreeToRadian(degree float64) float64 {
	return degree * 3.14 / 180.0
}

func radianToDegree(radian float64) float64 {
	return radian / 3.14 * 180.0
}

func JoinPartyById(partyId string, userId string) error {
	db, err := sql.Open(config.C.DBType, config.C.DBPath)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT current_people, total_people FROM party WHERE id = " + partyId)
	checkErr(err)

	var currentPeople int
	var totalPeople int
	for rows.Next() {
		rows.Scan(&currentPeople, &totalPeople)
	}

	if currentPeople+1 >= totalPeople {
		return errors.New("over total People")
	} else {
		_, err := db.Query("UPDATE party SET current_people = " + strconv.FormatInt(int64(currentPeople+1), 10) + " WHERE id = " + partyId)
		checkErr(err)

		stmt, err := db.Prepare("INSERT INTO party_member(party_id, user_id) VALUES (?,?)")
		checkErr(err)

		_, err = stmt.Exec(partyId, userId)
		checkErr(err)

		return nil
	}
}

func LeaveParty(partyId string, userId string) {
	db, err := sql.Open(config.C.DBType, config.C.DBPath)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT current_people FROM party WHERE id = " + partyId)
	checkErr(err)

	var currentPeople int
	for rows.Next() {
		rows.Scan(&currentPeople)
	}

	if currentPeople <= 1 {
		_, err = db.Query("DELETE FROM party WHERE id=" + partyId)
		checkErr(err)
	} else {
		_, err = db.Query("UPDATE party SET current_people = " + strconv.FormatInt(int64(currentPeople-1), 10) + " WHERE id = " + partyId)
		checkErr(err)
	}

	_, err = db.Query("DELETE FROM party_member WHERE user_id=" + userId)
	checkErr(err)
}

func GetPartyById(id string) PartyInfo {
	db, err := sql.Open(config.C.DBType, config.C.DBPath)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT id, name, meet_time, latitude, longitude, current_people, total_people FROM party WHERE id = " + id)
	checkErr(err)

	var party PartyInfo
	for rows.Next() {
		err := rows.Scan(&party.Id, &party.Name, &party.MeetTime, &party.Latitude, &party.Longitude, &party.TotalPeople, &party.CurrentPeople)
		checkErr(err)
	}

	return party
}
