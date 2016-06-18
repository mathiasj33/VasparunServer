package main

import (
	"bytes"
	"db"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	db.Init("root", "selorenus")

	//ip := os.Getenv("OPENSHIFT_GO_IP") + ":" + os.Getenv("OPENSHIFT_GO_PORT")
	ip := "localhost:8080"
	http.HandleFunc("/vasparun", respondHandler)
	http.ListenAndServe(ip, nil)
}

func respondHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		op := r.FormValue("op") TESTEN
		switch op {
		case "ContainsUsername":
			fmt.Fprint(w, db.ContainsUser(op)) TEXT SPLITTEN UND SO
		case "AddUser":
			db.AddUser(op)
		case "SelectUserDistance":
			dist := db.SelectUserDistance(op)
			fmt.Fprint(w, dist)
		case "SelectAllUserDistances":
			m := db.SelectAllUserDistances()
			fmt.Fprint(w, createStringFromSortedMap(m))
		case "SelectHighestDistance":
			fmt.Fprint(w, db.SelectHighestDistance())
		case "SelectAverageDistance":
			fmt.Fprint(w, db.SelectAverageDistance())
		case "UpdateOrAddDistance":
			updateOrAddDistance(getStringAndFloat(op))
		}
	}
}

func createStringFromSortedMap(m db.SortedStringFloatMap) string {
	var buffer bytes.Buffer
	for i := 0; i < m.Length(); i++ {
		name, dist := m.GetFromIndex(i)
		distString := fmt.Sprintf("%v", dist)
		buffer.WriteString(name)
		buffer.WriteString("+")
		buffer.WriteString(distString)

		if i != m.Length()-1 {
			buffer.WriteString("$")
		}
	}
	return buffer.String()
}

func getStringAndFloat(text string) (string, float32) {
	arr := strings.Split(text, "+")
	s := arr[0]
	f, _ := strconv.ParseFloat(arr[1], 32)
	return s, float32(f)
}

func updateOrAddDistance(username string, dist float32) {
	if db.ContainsDistance(username) {
		db.UpdateDistance(username, dist)
	} else {
		db.AddDistance(username, dist)
	}
}
