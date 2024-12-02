package tasks

import (
	"context"
	"log/slog"
	"todoGRPC/internal/domain/models"
)

type Tasks struct {
	log *slog.Logger
}

func New(log *slog.Logger) *Tasks {
	return &Tasks{
		log: log,
	}
}

func (s *Tasks) Create(ctx context.Context, name, description, deadline, token string) (id int, err error) {
	panic("implement me")
}
func (s *Tasks) Get(ctx context.Context, page, countTaskOnPage int) (tasks []models.Task, err error) {
	panic("implement me")
}

func (s *Tasks) GetByName(ctx context.Context, userId int64, name, token string) (tasks []models.Task, err error) {
	panic("implement me")
}

func (s *Tasks) GetById(ctx context.Context, userId int64, token string, taskId int) (task models.Task, err error) {
	panic("implement me")
}

func (s *Tasks) Update(ctx context.Context, name, description, deadline, token string) (message string, err error) {
	panic("implement me")
}

func (s *Tasks) Delete(ctx context.Context, token string, taskId int) (message string, err error) {
	panic("implement me")
}
