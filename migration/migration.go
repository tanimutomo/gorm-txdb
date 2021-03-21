package main

import (
	"github.com/tanimutomo/gorm-txdb/entity"
	"github.com/tanimutomo/gorm-txdb/mysql"
)

func main() {
	db, err := mysql.New()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{})
}
