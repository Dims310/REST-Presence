package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api-presence/helper"
	"rest-api-presence/model"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// configuration
	router := httprouter.New()
	db := helper.DBConnection()
	defer db.Close()

	// all method
	router.GET("/dashboard", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "All data will appear here!")

		query_script := "SELECT id, name, presence, absence, late_presence FROM user"
		rows, err := db.Query(query_script)
		helper.CatchPanic(err)

		for rows.Next() {
			var id int
			var name string
			var presence, absence time.Time
			var late_presence bool

			err := rows.Scan(&id, &name, &presence, &absence, &late_presence)
			helper.CatchPanic(err)

			dataRows := model.User{
				Id:            id,
				Name:          name,
				Presence:      presence,
				Absence:       absence,
				Late_presence: late_presence,
			}

			byte, err := json.Marshal(dataRows)
			helper.CatchPanic(err)

			fmt.Fprint(w, "\n"+string(byte))
		}
	})

	router.POST("/presence", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := r.ParseForm()
		helper.CatchPanic(err)

		name := r.PostForm.Get("name")

		query_script := "SELECT name FROM user"
		rows, err := db.Query(query_script)
		helper.CatchPanic(err)

		validate, info := helper.CheckPrensence(rows, name)

		if name == "" {
			fmt.Fprint(w, helper.AppearJSON("Fill name field first!"))
		} else if validate {
			fmt.Fprint(w, info)
		} else {
			exec_script := "INSERT INTO user(name, presence, late_presence) VALUES (?, ?, ?)"
			var late_presence = helper.PresenceTime()

			_, err = db.Exec(exec_script, name, helper.TimeNow().Format("2006-01-02 15:04:05"), late_presence)
			helper.CatchPanic(err)

			fmt.Fprint(w, helper.AppearJSON("Presence success, have good day! :)"))
		}
	})

	router.PUT("/absence", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := r.ParseForm()
		helper.CatchPanic(err)

		id := r.PostForm.Get("id")

		query_script := "SELECT id, already_absence FROM user"
		rows, err := db.Query(query_script)
		helper.CatchPanic(err)

		validate, info := helper.CheckAbsence(rows, id)

		if id == "" {
			fmt.Fprint(w, helper.AppearJSON("Fill ID field first!"))
		} else if validate {
			fmt.Fprint(w, info)
		} else {
			exec_script := "UPDATE user SET absence=(?), already_absence=(?) WHERE id=(?)"
			_, err := db.Exec(exec_script, time.Now().Format("2006-01-02 15:04:05"), true, id)
			helper.CatchPanic(err)

			fmt.Fprint(w, helper.AppearJSON(("Absence success, thank you! :)")))
		}
	})

	router.GET("/dashboard/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		idParam := p.ByName("id")
		fmt.Fprint(w, "All data will appear here!\n")

		query_script := "SELECT id, name, presence, absence, late_presence FROM user WHERE id=(?)"
		rows, err := db.Query(query_script, idParam)
		helper.CatchPanic(err)

		info := helper.CheckIDWithData(rows, idParam)

		fmt.Fprint(w, info)

	})

	router.DELETE("/del/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		idParam := p.ByName("id")
		query_script := "SELECT id, name FROM user"
		rows, err := db.Query(query_script)
		helper.CatchPanic(err)

		validate, info, name := helper.CheckDelete(rows, idParam)

		if idParam == "" {
			fmt.Fprint(w, helper.AppearJSON("Fill ID field first!"))
		} else if validate {
			exec_script := "DELETE FROM user WHERE id=(?)"
			_, err := db.Exec(exec_script, idParam)
			helper.CatchPanic(err)

			fmt.Fprint(w, helper.AppearJSON("Data with Id "+idParam+" and Name "+name+" succcessfully deleted."))
		} else {
			fmt.Fprint(w, info)
		}
	})

	// running server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	defer server.Close()

	// information server
	fmt.Println("Server running at:", server.Addr, "\nPress CTRL+C to stop.")
	server.ListenAndServe()
}
