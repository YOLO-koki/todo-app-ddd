package repository

import (
	"context"
	"log"
	"todo-app/infrastructure/models"
	"todo-app/infrastructure/sql/myconf"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetById(id int) *models.Todo {
	db := myconf.DBConfig()
	todo, err := models.Todos(
		qm.Where("id=?", id),
	).One(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return todo
}
