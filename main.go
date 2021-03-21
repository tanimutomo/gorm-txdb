package main

import "github.com/tanimutomo/gorm-txdb/mysql"

func main() {
	_, err := mysql.NewTest("tmp")
	if err != nil {
		panic(err)
	}
}
