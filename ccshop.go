package main

import (
	"ccshop/database"
	"ccshop/router"
)

func main() {
	err := database.InitMysql()
	if err != nil {
		panic(err)
	}
	r := router.InitRouter()
	r.Run(":8080")
}
