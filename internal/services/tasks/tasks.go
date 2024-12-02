package tasks

import (
	"context"
	"log/slog"
	"todoGRPC/internal/domain/models"
)

type Tasks struct {
	log          *slog.Logger
	taskSaver    TaskSaver
	taskProvider TaskProvider
	taskUpdater  TaskUpdater
	taskDeleter  TaskDeleter
}

type TaskSaver interface {
	SaveTask(ctx context.Context, name, description, deadline, token string) (*models.Task, error)
}

type TaskProvider interface {
	Tasks(ctx context.Context, page, countTaskOnPage int) ([]models.Task, error)
	TaskByID(ctx context.Context, id string) (models.Task, error)
	TaskByName(ctx context.Context, userId int64, name string) (models.Task, error)
}

type TaskUpdater interface {
	UpdateTask(ctx context.Context, name, description, deadline, token string) (string, error)
}

type TaskDeleter interface {
	DeleteTask(ctx context.Context, id string) (string, error)
}

func New(log *slog.Logger,
	taskSaver TaskSaver,
	taskProvider TaskProvider,
	taskUpdater TaskUpdater,
	taskDeleter TaskDeleter) *Tasks {
	return &Tasks{
		log:          log,
		taskSaver:    taskSaver,
		taskProvider: taskProvider,
		taskUpdater:  taskUpdater,
		taskDeleter:  taskDeleter,
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
