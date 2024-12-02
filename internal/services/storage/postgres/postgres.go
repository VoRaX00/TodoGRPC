package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"todoGRPC/internal/domain/models"
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

func (s *Storage) SaveTask(ctx context.Context, name, description, deadline, token string) (*models.Task, error) {
	panic("implement me")
}

func (s *Storage) Tasks(ctx context.Context, page, countTaskOnPage int) ([]models.Task, error) {
	panic("implement me")
}

func (s *Storage) TaskByID(ctx context.Context, id string) (models.Task, error) {
	panic("implement me")
}

func (s *Storage) TaskByName(ctx context.Context, userId int64, name string) (models.Task, error) {
	panic("implement me")
}

func (s *Storage) UpdateTask(ctx context.Context, name, description, deadline, token string) (string, error) {
	panic("implement me")
}

func (s *Storage) DeleteTask(ctx context.Context, id string) (string, error) {
	panic("implement me")
}
