package myconf

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func DBConfig() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(todo-mysql-db:3306)/todo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	boil.SetDB(db)

	return db
}
