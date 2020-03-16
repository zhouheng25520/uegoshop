package main

import "ccshop/router"

func main() {
	r := router.InitRouter()
	r.Run(":8080")
}
