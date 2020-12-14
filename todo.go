package main

import (
	"database/sql"
	"examples/go-echo-vue/handlers"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDB("storage.db")
	migrate(db)

	e := echo.New()
	th := &handlers.TaskHandler{Db: db}

	e.File("/", "public/index.html")
	e.GET("/tasks", th.GetTasks)
	e.PUT("/tasks", th.PutTasks)
	e.DELETE("/tasks/:id", th.DeleteTasks)

	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("nil db!")
	}

	return db
}

func migrate(db *sql.DB) {
	s := `
CREATE TABLE IF NOT EXISTS tasks(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name VARCHAR NOT NULL 
);
`
	_, err := db.Exec(s)
	if err != nil {
		panic(err)
	}
}