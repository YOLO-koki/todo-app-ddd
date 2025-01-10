package repository

import (
	"context"
	"log"
	"todo-app/domain/task"
	"todo-app/infrastructure/models"
	"todo-app/infrastructure/sql/myconf"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TaskRepository struct{}

func (r *TaskRepository) FindById(id int) *task.Task {
	db := myconf.DBConfig()

	t, err := models.Tasks(
		qm.Where("id=?", id),
	).One(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}

	return task.ReNew(t.ID, t.Title, t.Description.String, task.ReNewStatus(t.Status))
}

func (r *TaskRepository) FindAll() []*task.Task {
	db := myconf.DBConfig()

	tasks, err := models.Tasks().All(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var tasksDTO []*task.Task
	for _, t := range tasks {
		object := task.ReNew(
			t.ID,
			t.Title,
			t.Description.String,
			task.ReNewStatus(t.Status),
		)
		tasksDTO = append(tasksDTO, object)
	}
	return tasksDTO
}
