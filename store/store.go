package store

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Note Id
type Note struct {
	Id              uint   `json:"id"`
	Title           string `json:"title"`
	Body            string `json:"body"`
	IsDir           uint   `json:"isDir"`
	ParentId        uint   `json:"parentId"`
	IsFav           uint   `json:"isFav"`
	CreateTimestamp uint   `json:"createTimestamp"`
	UpdateTimestamp uint   `json:"updateTimestamp"`
}

// getDatabase returns a *sql.DB and an error.
//
// This function opens a connection to a SQLite database with the file name "notee.db".
// It returns the *sql.DB object representing the connection, and an error if any occurred.
func getDatabase() (*sql.DB, error) {
	return sql.Open("sqlite3", "notee.db")
}

// CreateSchema creates a database schema.
//
// It does the following:
// - Gets the database connection.
// - Prepares a statement to create a table named "notes" if it does not already exist.
// - Executes the statement to create the table.
//
// This function does not take any parameters.
// It does not return anything.
func CreateSchema() {
	db, err := getDatabase()
	// Create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS notes (id INTEGER PRIMARY KEY AUTOINCREMENT, title VARCHAR(255), body TEXT, isDir INTEGER, parentId INTEGER, isFav INTEGER, createTimestamp INTEGER, updateTimestamp INTEGER)")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("---------------------------------------Successfully created tables---------------------------------------")
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

// CreateNote creates a new note in the database.
//
// It takes a pointer to a Note struct as a parameter.
// It does not return anything.
func CreateNote(note *Note) uint {
	db, err := getDatabase()
	if err != nil {
		log.Println(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	stmt, _ := db.Prepare("INSERT INTO notes (title, body, isDir, parentId, isFav, createTimestamp, updateTimestamp) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id")
	_, err = stmt.Exec(note.Title, note.Body, note.IsDir, note.ParentId, note.IsFav, note.CreateTimestamp, note.UpdateTimestamp)

	rows, err := db.Query("SELECT last_insert_rowid()")
	for rows.Next() {
		err := rows.Scan(&note.Id)
		if err != nil {
			log.Println(err)
			return 0
		}
	}
	fmt.Println(note.Id)

	if err != nil {
		log.Println(err)
		return 0
	} else {
		return note.Id
	}

}

func DeleteNote(id *uint) bool {
	db, err := getDatabase()
	if err != nil {
		log.Println(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	stmt, _ := db.Prepare("DELETE FROM notes WHERE id=? OR parentId=?")
	_, err = stmt.Exec(id, id)
	if err != nil {
		return false
	} else {
		return true
	}

}

func UpdateNoteTitle(id *uint, title *string) bool {
	db, err := getDatabase()
	if err != nil {
		log.Println(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	stmt, _ := db.Prepare("UPDATE notes SET title=? WHERE id=?")
	_, err = stmt.Exec(title, id)
	if err != nil {
		return false
	} else {
		return true
	}

}

func UpdateFav(id *uint, isFav *uint) bool {
	db, err := getDatabase()
	if err != nil {
		log.Println(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	stmt, _ := db.Prepare("UPDATE notes SET isFav=? WHERE id=?")
	_, err = stmt.Exec(isFav, id)
	if err != nil {
		return false
	} else {
		return true
	}

}

func UpdateNote(id *uint, body *string) bool {
	db, err := getDatabase()
	if err != nil {
		log.Println(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	stmt, _ := db.Prepare("UPDATE notes SET body=? WHERE id=?")
	_, err = stmt.Exec(body, id)
	if err != nil {
		return false
	} else {
		return true
	}

}

func GetNote(id *uint) *Note {
	db, err := getDatabase()
	if err != nil {
		log.Println(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	rows, err := db.Query("SELECT * FROM notes WHERE id=?", id)

	if err != nil {
		log.Fatal(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var note Note

	for rows.Next() {

		err = rows.Scan(&note.Id, &note.Title, &note.Body, &note.IsDir, &note.ParentId, &note.IsFav, &note.CreateTimestamp, &note.UpdateTimestamp)
		if err != nil {
			log.Fatal(err)
		}

	}
	return &note

}

// GetNotes retrieves all the notes from the database.
//
// It does not take any parameters.
// It returns a pointer to a slice of Note objects.
func GetNotes() *[]Note {
	db, err := getDatabase()
	if err != nil {
		log.Println(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	rows, err := db.Query("SELECT * FROM notes ORDER BY updateTimestamp DESC")

	if err != nil {
		log.Fatal(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var notes []Note

	for rows.Next() {

		note := Note{}
		err = rows.Scan(&note.Id, &note.Title, &note.Body, &note.IsDir, &note.ParentId, &note.IsFav, &note.CreateTimestamp, &note.UpdateTimestamp)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(note)
		notes = append(notes, note)

	}
	return &notes

}
