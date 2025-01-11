package repository

import (
	"context"
	"database/sql"
	"log"
	"todo-app/domain/task"
	"todo-app/infrastructure/models"
	"todo-app/infrastructure/sql/myconf"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository() task.TaskRepository {
	db := myconf.DBConfig()
	return &TaskRepository{db}
}

// FindById is find a task by task id.
func (r *TaskRepository) FindById(id int) *task.Task {
	t, err := models.Tasks(
		qm.Where("id=?", id),
	).One(context.Background(), r.db)
	if err != nil {
		log.Fatal(err)
	}

	return task.ReNew(t.ID, t.Title, t.Description.String, task.ReNewStatus(t.Status))
}

// FindAll is find all task records;
func (r *TaskRepository) FindAll() []*task.Task {
	tasks, err := models.Tasks().All(context.Background(), r.db)
	if err != nil {
		log.Fatal(err)
	}
	defer r.db.Close()

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
