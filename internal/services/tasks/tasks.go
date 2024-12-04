package tasks

import (
	"context"
	tasksv1 "github.com/VoRaX00/todoProto/gen/go/tasks"
	"log/slog"
)

type Tasks struct {
	log          *slog.Logger
	taskSaver    TaskSaver
	taskProvider TaskProvider
	taskUpdater  TaskUpdater
	taskDeleter  TaskDeleter
}

type TaskSaver interface {
	SaveTask(ctx context.Context, name, description, deadline, token string) (*tasksv1.Task, error)
}

type TaskProvider interface {
	Tasks(ctx context.Context, page, countTaskOnPage int) ([]*tasksv1.Task, error)
	TaskByID(ctx context.Context, id string) (*tasksv1.Task, error)
	TaskByName(ctx context.Context, userId int64, name string) (*tasksv1.Task, error)
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

func (s *Tasks) Create(ctx context.Context, name, description, deadline string, userId int64) (id int64, err error) {
	panic("implement me")
}

func (s *Tasks) Get(ctx context.Context, page, countTaskOnPage, userId int64) (tasks []*tasksv1.Task, err error) {
	panic("implement me")
}

func (s *Tasks) GetByName(ctx context.Context, name string, userId int64) (tasks []*tasksv1.Task, err error) {
	panic("implement me")
}

func (s *Tasks) GetById(ctx context.Context, userId, taskId int64) (task *tasksv1.Task, err error) {
	panic("implement me")
}

func (s *Tasks) Update(ctx context.Context, taskId int64, name, description, deadline string, userId int64) (message string, err error) {
	panic("implement me")
}

func (s *Tasks) Delete(ctx context.Context, taskId, userId int64) (message string, err error) {
	panic("implement me")
}
