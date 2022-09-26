package main

import (
	"fmt"

	"actiontech.cloud/universe/ucommon/v4/log"
	"actiontech.cloud/universe/ucommon/v4/mysql"
)

func main() {
	stage := log.NewStage().Enter("test")

	db, err := mysql.OpenDb(stage, "test", "test", "", "10.186.63.29", "3333", 5, 5)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	row, err := mysql.SqlQuery(stage, db, "SHOW DATABASES")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", row)
}
