package handler

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DbHandler struct {
	Db *gorm.DB
}

func NewDbHandler(dbms, user, password, dbname string) (*DbHandler, error) {
	connect := user + ":" + password + "@/" + dbname
	db, err := gorm.Open(dbms, connect)
	return &DbHandler{db}, err
}
