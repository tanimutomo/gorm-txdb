package dao

import (
	"github.com/tanimutomo/gorm-txdb/mysql"
	"gorm.io/gorm"
)

func prepare(name string) *gorm.DB {
	db, err := mysql.NewTest(name)
	if err != nil {
		panic(err)
	}
	return db
}
