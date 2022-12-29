package helper

import (
	"database/sql"
	"rest-api-presence/model"
	"strconv"
	"time"
)

func CheckPrensence(rows *sql.Rows, nameFromMain string) (r bool, s string) {
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		CatchPanic(err)

		if nameFromMain == name {
			s = AppearJSON("Name already presence!")
			r = true
			break
		}
	}
	return
}

func CheckAbsence(rows *sql.Rows, idFromMain string) (r bool, s string) {
	for rows.Next() {
		var id string
		var alreadyAbsence bool
		err := rows.Scan(&id, &alreadyAbsence)
		CatchPanic(err)

		if id == idFromMain {
			if alreadyAbsence {
				s = AppearJSON("Name already absence!")
				r = true
				break
			} else {
				r = false
				break
			}
		} else {
			s = AppearJSON("ID not found, try another!")
			r = true
		}
	}
	return
}

func CheckDelete(rows *sql.Rows, idFromMain string) (r bool, s string, n string) {
	for rows.Next() {
		var id, name string
		rows.Scan(&id, &name)

		if id == idFromMain {
			r = true
			n = name
			break
		} else {
			r = false
		}
	}

	if !r {
		s = AppearJSON("ID not found, try another!")
	}
	return
}

func CheckIDWithData(rows *sql.Rows, idFromMain string) (s string) {
	idConv, err := strconv.Atoi(idFromMain)
	CatchPanic(err)

	var r bool

	for rows.Next() {
		var id int
		var name string
		var presence, absence time.Time
		var late_presence bool

		rows.Scan(&id, &name, &presence, &absence, &late_presence)

		if id == idConv {
			dataSingleRows := model.User{
				Id:            id,
				Name:          name,
				Presence:      presence,
				Absence:       absence,
				Late_presence: late_presence,
			}
			s = AppearJSON(dataSingleRows)
			r = true
			break
		} else {
			r = false
		}
	}
	if !r {
		s = AppearJSON("Data with Id " + idFromMain + " not found, try another!")
	}
	return
}
