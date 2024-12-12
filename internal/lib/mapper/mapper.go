package mapper

import (
	tasksv1 "github.com/VoRaX00/todoProto/gen/go/tasks"
	"todoGRPC/internal/domain/models"
)

//func MapToTask(task *tasksv1.Task) (models.Task, error) {
//	deadline, err := time.Parse("02.01.2006", task.Deadline)
//	if err != nil {
//		return models.Task{}, err
//	}
//	return models.Task{
//		Name:         task.Name,
//		Descriptions: task.Descriptions,
//		TypeTask: models.TypeTask{
//			TypeTask: task.TypeTask,
//		},
//		Deadline: deadline,
//	}, nil
//}

func MapToTaskV1(task models.Task) *tasksv1.Task {
	return &tasksv1.Task{
		Id:           task.Id,
		Name:         task.Name,
		Descriptions: task.Descriptions,
		TypeTask:     task.TypeTask.TypeTask,
		Deadline:     task.Deadline.String(),
	}
}
