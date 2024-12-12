package tasks

import (
	"context"
	"errors"
	tasksv1 "github.com/VoRaX00/todoProto/gen/go/tasks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	taskService "todoGRPC/internal/services/tasks"
)

type Tasks interface {
	Create(ctx context.Context, name, description, typeTask, deadline string, userId int64) (id int64, err error)
	Get(ctx context.Context, page, countTaskOnPage, userId int64) (tasks []*tasksv1.Task, err error)
	GetByName(ctx context.Context, name string, userId int64) (tasks []*tasksv1.Task, err error)
	GetById(ctx context.Context, taskId int64) (task *tasksv1.Task, err error)
	Update(ctx context.Context, name, description, typeTask, deadline string, taskId int64) error
	Delete(ctx context.Context, taskId int64) error
}

type serverAPI struct {
	tasksv1.UnimplementedTasksServer
	tasks Tasks
}

func Register(gRPC *grpc.Server, tasks Tasks) {
	tasksv1.RegisterTasksServer(gRPC, &serverAPI{
		tasks: tasks,
	})
}

func (s *serverAPI) Create(ctx context.Context, req *tasksv1.CreateTaskRequest) (*tasksv1.CreateTaskResponse, error) {
	if err := validateCreate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := s.tasks.Create(ctx, req.Name, req.Descriptions, req.TypeTask, req.Deadline, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &tasksv1.CreateTaskResponse{
		TaskId: id,
	}, nil
}

func validateCreate(req *tasksv1.CreateTaskRequest) error {
	_, err := time.Parse("02.01.2006", req.Deadline)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid deadline")
	}

	if req.Name == "" {
		return status.Error(codes.InvalidArgument, "invalid name")
	}
	return nil
}

func (s *serverAPI) Get(ctx context.Context, req *tasksv1.GetAllRequest) (*tasksv1.GetAllResponse, error) {
	if err := validateGet(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tasks, err := s.tasks.Get(ctx, req.Page, req.CountTaskOnPage, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &tasksv1.GetAllResponse{
		Task: tasks,
	}, nil
}

func validateGet(req *tasksv1.GetAllRequest) error {
	if req.UserId <= 0 {
		return status.Error(codes.InvalidArgument, "invalid user id")
	}

	if req.Page <= 0 {
		return status.Error(codes.InvalidArgument, "invalid page")
	}

	if req.CountTaskOnPage <= 0 {
		return status.Error(codes.InvalidArgument, "invalid count task on page")
	}
	return nil
}

func (s *serverAPI) GetByName(ctx context.Context, req *tasksv1.GetByNameRequest) (*tasksv1.GetByNameResponse, error) {
	err := validateGetByName(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tasks, err := s.tasks.GetByName(ctx, req.Name, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &tasksv1.GetByNameResponse{
		Task: tasks,
	}, nil
}

func validateGetByName(req *tasksv1.GetByNameRequest) error {
	if req.UserId <= 0 {
		return status.Error(codes.InvalidArgument, "invalid user id")
	}

	if req.Name == "" {
		return status.Error(codes.InvalidArgument, "invalid name")
	}
	return nil
}

func (s *serverAPI) GetById(ctx context.Context, req *tasksv1.GetByIdRequest) (*tasksv1.GetByIdResponse, error) {
	if err := validateGetById(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	task, err := s.tasks.GetById(ctx, req.TaskId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &tasksv1.GetByIdResponse{
		Task: task,
	}, nil
}

func validateGetById(req *tasksv1.GetByIdRequest) error {
	if req.TaskId <= 0 {
		return status.Error(codes.InvalidArgument, "invalid task id")
	}
	return nil
}

func (s *serverAPI) Update(ctx context.Context, req *tasksv1.UpdateTaskRequest) (*tasksv1.UpdateTaskResponse, error) {
	if err := validateUpdate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.tasks.Update(ctx, req.Task.Name, req.Task.Descriptions, req.Task.TypeTask, req.Task.Deadline, req.Task.Id)
	if err != nil {
		if errors.Is(err, taskService.ErrTaskNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &tasksv1.UpdateTaskResponse{
		Message: "success",
	}, nil
}

func validateUpdate(req *tasksv1.UpdateTaskRequest) error {
	if req.Task.Name == "" {
		return status.Error(codes.InvalidArgument, "invalid name")
	}

	if _, err := time.Parse("02.01.2006", req.Task.Deadline); err != nil {
		return status.Error(codes.InvalidArgument, "invalid deadline")
	}

	if req.Task.Id <= 0 {
		return status.Error(codes.InvalidArgument, "invalid task id")
	}
	return nil
}

func (s *serverAPI) Delete(ctx context.Context, req *tasksv1.DeleteTaskRequest) (*tasksv1.DeleteTaskResponse, error) {
	if err := validateDelete(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.tasks.Delete(ctx, req.TaskId)
	if err != nil {
		if errors.Is(err, taskService.ErrTaskNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &tasksv1.DeleteTaskResponse{
		Message: "success",
	}, nil
}

func validateDelete(req *tasksv1.DeleteTaskRequest) error {
	if req.TaskId <= 0 {
		return status.Error(codes.InvalidArgument, "invalid task id")
	}
	return nil
}
