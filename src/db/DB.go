package db

import (
	"database/sql"
	"fmt"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init(username string, password string) {
	var err error
	db, err = sql.Open("mysql", username+":"+password+"@tcp(localhost:3307)/vasparun")
	handleError(err)
}

func AddUser(username string) {
	executeQuery("INSERT INTO users VALUES(null, ?)", username)
}

func ContainsUser(username string) bool {
	rows := getRows("SELECT Username FROM Users WHERE Username = ?", username)
	return getNumRows(rows) > 0
}

func SelectUserDistance(username string) float32 {
	return getFloat(`SELECT Distance FROM users
					JOIN distances
					WHERE users.ID = distances.U_ID
					AND users.Username = ?`, username)
}

func SelectAllUserDistances() SortedStringFloatMap {
	rows := getRows(`SELECT Username, Distance FROM users 
					JOIN distances
					WHERE users.ID = distances.U_ID 
					ORDER BY Distance DESC`)
	defer rows.Close()
	
	distMap := NewSortedStringFloatMap()
	
	var username string
	var distance float32
	for rows.Next() {
		err := rows.Scan(&username, &distance)
		handleError(err)
		distMap.Put(username, distance)
	}
	return *distMap
}

func SelectHighestDistance() float32 {
	return getFloat(`SELECT MAX(Distance) FROM distances`)
}

func SelectAverageDistance() float32 {
	return getFloat("SELECT AVG(Distance) FROM distances")
}

func AddDistance(username string, distance float32) {
	id := getID(username)
	executeQuery("INSERT INTO distances VALUES (null, ?, ?)", distance, id)
}

func ContainsDistance(username string) bool {
	id := getID(username)
	rows := getRows("SELECT Distance FROM distances WHERE U_ID = ?", id)
	return getNumRows(rows) > 0
}

func UpdateDistance(username string, distance float32) {
	id := getID(username)
	fmt.Println(id)
	if id == -1 {
		handleError(errors.New("This user does not exist"))
		return
	}
	executeQuery(`update distances
				set Distance = ?
				where U_ID = ?`, distance, id)
}

func SelectUserTime(username string, level int) float32 {
	id := getID(username)
	return getFloat("SELECT Time FROM times WHERE U_ID = ? AND Level = ?", id, level)
	
}

func SelectAllUserTimes(level int) SortedStringFloatMap {
	rows := getRows(`SELECT Username, Time FROM users 
					JOIN times
					WHERE users.ID = times.U_ID
					AND Level = ?
					ORDER BY Time`, level)
	defer rows.Close()
	
	times := NewSortedStringFloatMap()
	var username string
	var time float32
	
	for rows.Next() {
		rows.Scan(&username, &time)
		times.Put(username, time)
	}
	return *times
}

func SelectBestTime(level int) float32 {
	return getFloat("SELECT MIN(Time) FROM times WHERE Level = ?", level)
}

func SelectAverageTime(level int) float32 {
	return getFloat("SELECT AVG(Time) FROM times WHERE Level = ?", level)
}

func AddTime(username string, level int, time float32) {
	id := getID(username)
	executeQuery("INSERT INTO times VALUES (null, ?, ?, ?)", level, time, id)
}

func ContainsTime(username string, level int) bool {
	id := getID(username)
	rows := getRows("SELECT Time FROM times WHERE U_ID = ? AND Level = ?", id, level)
	return getNumRows(rows) > 0
}

func UpdateTime(username string, level int, time float32) {
	id := getID(username)
	executeQuery(`UPDATE Times
				SET Time = ?
				WHERE U_ID = ?
				AND Level = ?`, time, id, level)
}

func getID(username string) int {
	return getInt("SELECT ID FROM users WHERE Username = ?", username)
}

func getInt(statement string, args ...interface{}) int {
	rows := getRows(statement, args...)
	defer rows.Close()
	
	var i int
	if rows.Next() {
		err := rows.Scan(&i)
		handleError(err)
		if err != nil {
			return -1
		}
		return i
	}
	return -1
}

func getFloat(statement string, args ...interface{}) float32 {
	rows := getRows(statement, args...)
	defer rows.Close()
	
	var f float32
	if rows.Next() {
		err := rows.Scan(&f)
		handleError(err)
		if err != nil {
			return -1
		}
		return f
	}
	return -1
}

func getNumRows(rows *sql.Rows) int {
	count := 0
	for rows.Next() {
		count++
	}
	return count
}

func getRows(statement string, args ...interface{}) *sql.Rows {
	stmt, err := db.Prepare(statement)
	handleError(err)
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	handleError(err)
	return rows
}

func executeQuery(statement string, args ...interface{}) {
	stmt, err := db.Prepare(statement)
	handleError(err)
	defer stmt.Close()
	_, err = stmt.Exec(args...)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Print("ERROR: ")
		fmt.Println(err)
	}
}
