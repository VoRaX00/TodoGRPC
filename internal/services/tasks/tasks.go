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
	SaveTask(ctx context.Context, name, description, deadline string, userId int64) (int64, error)
}

type TaskProvider interface {
	Tasks(ctx context.Context, page, countTaskOnPage, userId int64) ([]*tasksv1.Task, error)
	TaskByID(ctx context.Context, taskId int64) (*tasksv1.Task, error)
	TaskByName(ctx context.Context, userId int64, name string) ([]*tasksv1.Task, error)
}

type TaskUpdater interface {
	UpdateTask(ctx context.Context, name, description, deadline string, taskId int64) error
}

type TaskDeleter interface {
	DeleteTask(ctx context.Context, taskId int64) error
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

func (s *Tasks) Create(ctx context.Context, name, description, deadline string, userId int64) (int64, error) {
	const op = "tasks.Create"
	log := s.log.With(
		slog.String("op", op),
		slog.String("name", name),
		slog.Int64("userId", userId))

	log.Info("creating task")

	id, err := s.taskSaver.SaveTask(ctx, name, description, deadline, userId)
	if err != nil {
		// TODO: обработка ошибки
	}

	log.Info("success create task")
	return id, err
}

func (s *Tasks) Get(ctx context.Context, page, countTaskOnPage, userId int64) (tasks []*tasksv1.Task, err error) {
	const op = "tasks.Get"
	log := s.log.With(
		slog.String("op", op),
		slog.Int64("page", page),
		slog.Int64("userId", userId))

	log.Info("getting tasks")
	res, err := s.taskProvider.Tasks(ctx, page, countTaskOnPage, userId)
	if err != nil {
		// TODO: обработка ошибки
	}

	log.Info("success get tasks")
	return res, nil
}

func (s *Tasks) GetByName(ctx context.Context, name string, userId int64) (tasks []*tasksv1.Task, err error) {
	const op = "tasks.GetByName"
	log := s.log.With(
		slog.String("op", op),
		slog.Int64("userId", userId))

	log.Info("getting tasks")
	res, err := s.taskProvider.TaskByName(ctx, userId, name)
	if err != nil {
		// TODO: обработка ошибки
	}

	log.Info("success get tasks")
	return res, nil
}

func (s *Tasks) GetById(ctx context.Context, taskId int64) (task *tasksv1.Task, err error) {
	const op = "tasks.GetById"
	log := s.log.With(
		slog.String("op", op),
		slog.Int64("taskId", taskId),
	)

	log.Info("getting task")
	res, err := s.taskProvider.TaskByID(ctx, taskId)
	if err != nil {
		// TODO: обработка ошибки
	}
	log.Info("success get task")
	return res, nil
}

func (s *Tasks) Update(ctx context.Context, name, description, deadline string, taskId int64) error {
	const op = "tasks.Update"
	log := s.log.With(
		slog.String("op", op),
		slog.Int64("taskId", taskId),
	)

	log.Info("updating task")
	err := s.taskUpdater.UpdateTask(ctx, name, description, deadline, taskId)
	if err != nil {
		// TODO: обработка ошибки
	}

	log.Info("success update task")
	return err
}

func (s *Tasks) Delete(ctx context.Context, taskId int64) error {
	const op = "tasks.Delete"
	log := s.log.With(
		slog.String("op", op),
		slog.Int64("taskId", taskId),
	)

	log.Info("deleting task")
	err := s.taskDeleter.DeleteTask(ctx, taskId)
	if err != nil {
		// TODO: обработка ошибки
	}
	return nil
}
