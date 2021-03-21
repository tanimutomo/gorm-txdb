package dao

import (
	"github.com/tanimutomo/gorm-txdb/pkg/mysql"
	"gorm.io/gorm"
)

func prepare(name string, seeds []interface{}) (*gorm.DB, error) {
	db, err := mysql.NewTest(name)
	if err != nil {
		return nil, err
	}
	for _, s := range seeds {
		if err := db.Create(s).Error; err != nil {
			return nil, err
		}
	}
	return db, nil
}
