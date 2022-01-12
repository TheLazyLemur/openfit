package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var driver string = "sqlite3"
var connectionString string = "./workout.db"

func InitDb() {
	println("Creating Database....")
	os.Remove(connectionString)

	db, err := sql.Open(driver, connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
create table exercise (id integer not null primary key, name text);
	delete from exercise;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func AddExercise(e Exercise) {
	db, err := sql.Open(driver, connectionString)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into exercise(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id, e.Name)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func ListExercises() []Exercise {
	var exerciseList []Exercise

	db, err := sql.Open(driver, connectionString)
	rows, err := db.Query("select id, name from exercise")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
		var e Exercise
		e.Name = name
		e.Id = id
		exerciseList = append(exerciseList, e)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return exerciseList
}

func GetExercise(id int) Exercise {
	db, err := sql.Open(driver, connectionString)

	stmt, err := db.Prepare("select name from exercise where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow(fmt.Sprintf("%03d", id)).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	var e Exercise
	e.Name = name
	e.Id = id
	return e
}

func DeleteExercise(id int) {
	db, err := sql.Open(driver, connectionString)
	tx, err := db.Begin()

	stmt, err := db.Prepare("delete from exercise where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(fmt.Sprintf("%03d", id))
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
}
