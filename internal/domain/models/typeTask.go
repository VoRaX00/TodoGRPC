package models

type TypeTask struct {
	Id       int64  `db:"id"`
	TypeTask string `db:"type_task"`
}
