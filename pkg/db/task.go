package db

import "database/sql"

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

func Tasks(limit int) ([]*Task, error) {
	query := `SELECT * FROM scheduler ORDER BY date LIMIT :limit`

	rows, err := db.Query(query, sql.Named("limit", limit))

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
