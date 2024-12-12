package mapper

import (
	tasksv1 "github.com/VoRaX00/todoProto/gen/go/tasks"
	"todoGRPC/internal/domain/models"
)

func MapToTaskV1(task models.Task) *tasksv1.Task {
	return &tasksv1.Task{
		Id:           task.Id,
		Name:         task.Name,
		Descriptions: task.Descriptions,
		TypeTask:     task.TypeTask.TypeTask,
		Deadline:     task.Deadline.String(),
	}
}
