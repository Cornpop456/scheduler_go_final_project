package db

import (
	"database/sql"
	"errors"
	"time"
)

const searchDateLayout = "02.01.2006"
const timeLayout = "20060102"

var ErrWrongId = errors.New("wrong id for task")

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)`

	res, err := db.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)

	if err == nil {
		id, err = res.LastInsertId()
	}

	return id, err
}

func GetTask(id string) (*Task, error) {
	query := `SELECT * FROM scheduler WHERE id = :id`
	row := db.QueryRow(query, sql.Named("id", id))

	task := &Task{}

	if err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrWrongId
		}

		return nil, err
	}

	return task, nil
}

func UpdateTask(task *Task) error {
	query := `UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id`

	res, err := db.Exec(query,
		sql.Named("id", task.ID),
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if count == 0 {
		return ErrWrongId
	}

	return nil
}

func UpdateDate(next string, id string) error {
	query := `UPDATE scheduler SET date = :date WHERE id = :id`

	res, err := db.Exec(query,
		sql.Named("id", id),
		sql.Named("date", next),
	)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if count == 0 {
		return ErrWrongId
	}

	return nil
}

func DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id = :id`

	res, err := db.Exec(query, sql.Named("id", id))

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if count == 0 {
		return ErrWrongId
	}

	return nil
}

func Tasks(limit int, search string) ([]*Task, error) {
	query := `SELECT * FROM scheduler WHERE title LIKE :search OR comment LIKE :search ORDER BY date LIMIT :limit`
	queryDate := `SELECT * FROM scheduler WHERE date == :date ORDER BY date LIMIT :limit`

	var rows *sql.Rows
	var err error

	t, err := time.Parse(searchDateLayout, search)

	if err == nil {
		search = t.Format(timeLayout)
		rows, err = db.Query(queryDate, sql.Named("date", search), sql.Named("limit", limit))
	} else {
		search = "%" + search + "%"
		rows, err = db.Query(query, sql.Named("search", search), sql.Named("limit", limit))
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]*Task, 0)

	for rows.Next() {
		task := &Task{}

		if err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
