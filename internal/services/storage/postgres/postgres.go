package postgres

import (
	"context"
	"fmt"
	tasksv1 "github.com/VoRaX00/todoProto/gen/go/tasks"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sqlx.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) SaveTask(ctx context.Context, name, description, deadline string, userId int64) (int64, error) {
	panic("implement me")
}

func (s *Storage) Tasks(ctx context.Context, page, countTaskOnPage, userId int64) ([]*tasksv1.Task, error) {
	panic("implement me")
}

func (s *Storage) TaskByID(ctx context.Context, userId, taskId int64) (*tasksv1.Task, error) {
	panic("implement me")
}

func (s *Storage) TaskByName(ctx context.Context, userId int64, name string) ([]*tasksv1.Task, error) {
	panic("implement me")
}

func (s *Storage) UpdateTask(ctx context.Context, name, description, deadline string, userId int64) (string, error) {
	panic("implement me")
}

func (s *Storage) DeleteTask(ctx context.Context, taskId, userId int64) (string, error) {
	panic("implement me")
}
