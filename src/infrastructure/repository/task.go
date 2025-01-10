package repository

import (
	"context"
	"log"
	"todo-app/domain/task"
	"todo-app/infrastructure/models"
	"todo-app/infrastructure/sql/myconf"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetById(id int) *task.Task {
	db := myconf.DBConfig()
	t, err := models.Tasks(
		qm.Where("id=?", id),
	).One(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	object := task.Task{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description.String,
		Status:      task.ReNewStatus(t.Status),
	}
	return task.ReNew(object.ID, object.Title, object.Description, object.Status)
}
