package main  //TODO: Zeichen die nicht gehen: :, $

import (
	"bytes"
	"db"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	db.Init("root", "selorenus") //TODO: Username wird in Playerprefs oder so gespeichert - ist nur für einen PC

	//ip := os.Getenv("OPENSHIFT_GO_IP") + ":" + os.Getenv("OPENSHIFT_GO_PORT")
	ip := "localhost:8080"
	http.HandleFunc("/vasparun", respondHandler)
	http.ListenAndServe(ip, nil)
}

func respondHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		op := r.FormValue("op") //TODO: testen
		username := r.FormValue("username")
		level, _ := strconv.Atoi(r.FormValue("level"))
		fmt.Println("level:", level)
		
		switch op {
		case "ContainsUsername":
			fmt.Fprint(w, db.ContainsUser(username))
		case "AddUser":
			db.AddUser(username)
		case "SelectUserDistance":
			dist := db.SelectUserDistance(username)
			fmt.Fprint(w, dist)
		case "SelectAllUserDistances":
			m := db.SelectAllUserDistances()
			fmt.Fprint(w, createStringFromSortedMap(m))
		case "SelectHighestDistance":
			fmt.Fprint(w, db.SelectHighestDistance())
		case "SelectAverageDistance":
			fmt.Fprint(w, db.SelectAverageDistance())
		case "UpdateOrAddDistance":
			dist := r.FormValue("distance")
			updateOrAddDistance(username, getFloat(dist))
		case "SelectUserTime":
			fmt.Fprint(w, db.SelectUserTime(username, level))
		case "SelectAllUserTimes":
			m := db.SelectAllUserTimes(level)
			fmt.Fprint(w, createStringFromSortedMap(m))
		case "SelectBestTime":
			fmt.Fprint(w, db.SelectBestTime(level))
		case "SelectAverageTime":
			fmt.Fprint(w, db.SelectAverageTime(level))
		case "ContainsTime":
			fmt.Fprint(w, db.ContainsTime(username, level))
		case "UpdateOrAddTime":
			time := r.FormValue("time")
			updateOrAddTime(username, level, getFloat(time))
		}
	}
}

func createStringFromSortedMap(m db.SortedStringFloatMap) string {
	var buffer bytes.Buffer
	for i := 0; i < m.Length(); i++ {
		name, dist := m.GetFromIndex(i)
		distString := fmt.Sprintf("%v", dist)
		buffer.WriteString(name)
		buffer.WriteString(":")
		buffer.WriteString(distString)

		if i != m.Length()-1 {
			buffer.WriteString("$")
		}
	}
	return buffer.String()
}

func getFloat(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}

func updateOrAddDistance(username string, dist float32) {
	if db.ContainsDistance(username) {
		db.UpdateDistance(username, dist)
	} else {
		db.AddDistance(username, dist)
	}
}

func updateOrAddTime(username string, level int, time float32) {
	if db.ContainsTime(username, level) {
		db.UpdateTime(username, level, time)
	} else {
		db.AddTime(username, level, time)
	}
}
