package tasks

import (
	"context"
	tasksv1 "github.com/VoRaX00/todoProto/gen/go/tasks"
	"google.golang.org/grpc"
	"todoGRPC/internal/domain/models"
)

type Tasks interface {
	Create(ctx context.Context, name, description, deadline, token string) (id int, err error)
	Get(ctx context.Context, page, countTaskOnPage int) (tasks []models.Task, err error)
	GetByName(ctx context.Context, userId int64, name, token string) (tasks []models.Task, err error)
	GetById(ctx context.Context, userId int64, token string, taskId int) (task models.Task, err error)
	Update(ctx context.Context, name, description, deadline, token string) (message string, err error)
	Delete(ctx context.Context, token string, taskId int) (message string, err error)
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
	panic("implement me")
}

func (s *serverAPI) Get(ctx context.Context, req *tasksv1.GetAllRequest) (*tasksv1.GetAllResponse, error) {
	panic("implement me")
}

func (s *serverAPI) GetByName(ctx context.Context, req *tasksv1.GetByNameRequest) (*tasksv1.GetByNameResponse, error) {
	panic("implement me")
}

func (s *serverAPI) GetById(ctx context.Context, req *tasksv1.GetByIdRequest) (*tasksv1.GetByIdResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Update(ctx context.Context, req *tasksv1.UpdateTaskRequest) (*tasksv1.UpdateTaskResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Delete(ctx context.Context, req *tasksv1.DeleteTaskRequest) (*tasksv1.DeleteTaskResponse, error) {
	panic("implement me")
}
