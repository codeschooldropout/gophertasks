package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "./gophertask-sqlite.db")
	if err != nil {
		return err
	}
	return db.Ping()
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS gophertask (
		"idTask" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"task" TEXT,
		"description" TEXT,
		"category" TEXT,
		"status" TEXT
		);`
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Table gophertask created")
}

func InsertTask(task string, description string, category string, status string) {
	insertNoteSQL := `INSERT INTO gophertask (task, description, category, status) VALUES (?, ?, ?, ?);`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec(task, description, category, status)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Note inserted")
}

func DisplayAllTasks() {
	row, err := db.Query("SELECT * FROM gophertask ORDER BY task")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer row.Close()

	for row.Next() {
		var taskNote int
		var task string
		var description string
		var category string
		var status string

		row.Scan(&taskNote, &task, &description, &category, &status)
		log.Println("[", category, "]", task, "-", description)
	}

}
