package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	tasksv1 "github.com/VoRaX00/todoProto/gen/go/tasks"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"todoGRPC/internal/services/storage"
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

func (s *Storage) SaveTask(ctx context.Context, name, description, typeTask, deadline string, userId int64) (int64, error) {
	const op = "storage.postgres.SaveTask"
	query := `INSERT INTO tasks (name_task, descriptions, deadline, user_id) VALUES ($1, $2, $3, $4) RETURNING id`
	row := s.db.QueryRowContext(ctx, query, name, description, deadline, userId)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *Storage) Tasks(ctx context.Context, page, countTaskOnPage, userId int64) ([]*tasksv1.Task, error) {
	const op = "storage.postgres.Tasks"
	query := `SELECT name_task, descriptions, deadline FROM tasks WHERE user_id = $1 LIMIT $2 OFFSET $3`
	var tasks []*tasksv1.Task
	err := s.db.GetContext(ctx, tasks, query, userId, page, countTaskOnPage)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return tasks, nil
}

func (s *Storage) TaskByID(ctx context.Context, taskId int64) (*tasksv1.Task, error) {
	const op = "storage.postgres.TaskByID"
	query := `SELECT name_task, descriptions, deadline FROM tasks WHERE id = $1`
	var task tasksv1.Task
	err := s.db.GetContext(ctx, &task, query, taskId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrTaskNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &task, nil
}

func (s *Storage) TaskByName(ctx context.Context, userId int64, name string) ([]*tasksv1.Task, error) {
	const op = "storage.postgres.TaskByName"
	query := `SELECT name_task, descriptions, deadline FROM tasks WHERE name_task = $1 AND user_id = $2`
	var tasks []*tasksv1.Task
	err := s.db.SelectContext(ctx, &tasks, query, name, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrTaskNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return tasks, nil
}

func (s *Storage) UpdateTask(ctx context.Context, name, description, typeTask, deadline string, taskId int64) error {
	const op = "storage.postgres.UpdateTask"
	query := `UPDATE tasks SET name_task = $1, descriptions = $2, deadline = $3 WHERE id = $4`
	res, err := s.db.ExecContext(ctx, query, name, description, deadline, taskId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rows == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrTaskNotFound)
	}
	return nil
}

func (s *Storage) DeleteTask(ctx context.Context, taskId int64) error {
	const op = "storage.postgres.DeleteTask"

	query := `DELETE FROM tasks WHERE id = $1`
	res, err := s.db.ExecContext(ctx, query, taskId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rows != 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrTaskNotFound)
	}
	return nil
}
