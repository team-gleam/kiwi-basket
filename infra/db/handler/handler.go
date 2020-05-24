package handler

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Config struct {
	DBMS     string `yaml:"dbms"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Protocol string `yaml:"protocol"`
	DBName   string `yaml:"dbname"`
}

type DbHandler struct {
	Db *gorm.DB
}

func NewDbHandler(c Config) (*DbHandler, error) {
	connect := c.User + ":" + c.Password + "@" + c.Protocol + "/" + c.DBName + "?charset=utf8mb4" + "&parseTime=true"
	db, err := gorm.Open(c.DBMS, connect)
	return &DbHandler{db}, err
}
