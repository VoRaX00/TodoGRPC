package models

import "time"

type Task struct {
	Id           int64     `db:"id"`
	Name         string    `db:"name"`
	Descriptions string    `db:"description"`
	TypeTask     TypeTask  `db:"type_task"`
	Deadline     time.Time `db:"deadline"`
}
