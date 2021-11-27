package data

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type Recipe struct {
	Name string
}

const (
	dbname   = "mealbuddy"
	host     = "localhost"
	password = "postgres"
	port     = 5432
	table    = "recipes"
	user     = "efrenulloa"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func connect() *sql.DB {
	conn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname,
	)

	db, connErr := sql.Open("postgres", conn)

	CheckError(connErr)

	pingErr := db.Ping()

	CheckError(pingErr)

	return db
}

func GetRecipes() ([]byte, error) {
	db := connect()

	stmt := fmt.Sprintf("SELECT * FROM %s", table)

	rows, err := db.Query(stmt)

	CheckError(err)

	rs := make([]Recipe, 0)

	for rows.Next() {
		var r Recipe
		e := rows.Scan(&r.Name)
		CheckError(e)
		rs = append(rs, r)
	}

	defer rows.Close()
	defer db.Close()

	bs, err := json.Marshal(rs)
	return bs, err
}
