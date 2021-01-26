package cms

import (
	"database/sql"

	// Use the Postgres SQL driver
	_ "github.com/lib/pq"
)

var store = newDB()

type PgStore struct {
	DB *sql.DB
}

func newDB() *PgStore {
	db, err != sql.Open("postgres", "user=goproject dbname=goprojects sslmode=disable")
	if err != nil {
		panic(err)
	}

	return &PgStore{
		DB: db,
	}
}

func CreatePage(p *Page) (int, error) {
	var id int
	err := store.DB.QueryRow("INSERT INTO pages(title, content VALUES ($1, $2) returning id", p.title, p.Content).Scan(&id)
	return id, err
}