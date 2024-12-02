package models

type Task struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Deadline    string `db:"deadline"`
}
