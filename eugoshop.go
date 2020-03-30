package main

import (
	"uegoshop/database"
	"uegoshop/router"
	"net/http"
)

func main() {
	err := database.InitMysql()
	if err != nil {
		panic(err)
	}
	r := router.InitRouter()
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		//TLSConfig:         nil,
		//ReadTimeout:       0,
		//ReadHeaderTimeout: 0,
		//WriteTimeout:      0,
		//IdleTimeout:       0,
		//MaxHeaderBytes:    0,
		//TLSNextProto:      nil,
		//ConnState:         nil,
		//ErrorLog:          nil,
	}
	//r.Run(":8080")

	router.Run(srv)
	//ssignal.WaitCtrlC(func(s os.Signal) bool {
	//	fmt.Println("s>>>>>>>>>>>>")
	//	return false
	//})
}
