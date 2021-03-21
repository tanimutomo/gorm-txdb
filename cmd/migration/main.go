package main

import (
	"github.com/tanimutomo/gorm-txdb/pkg/entity"
	"github.com/tanimutomo/gorm-txdb/pkg/mysql"
)

func main() {
	db, err := mysql.New()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{})
}
