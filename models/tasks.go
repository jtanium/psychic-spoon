package models

import "database/sql"

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type TaskCollection []Task

func GetTasks(db *sql.DB) TaskCollection {
	s := "SELECT id, name FROM tasks"
	rows, err := db.Query(s)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.ID, &task.Name)
		result = append(result, task)
	}

	return result
}

func CreateTask(db *sql.DB, t *Task) error {
	s := "INSERT INTO tasks (name) VALUES (?)"
	stmt, err := db.Prepare(s)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(t.Name)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = int(id)

	return nil
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	s := "DELETE FROM tasks WHERE id = ?"

	stmt, err := db.Prepare(s)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}