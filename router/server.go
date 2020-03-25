package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run(srv *http.Server)  {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server closed by err:%s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill) // 貌似kill信号收不到
	<- quit
	log.Println("Shutdown Server ...")
	//for i := 0; i < 5; i++ {
	//	fmt.Printf("after %d second ,well be closed server\n", 5 - i)
	//	time.Sleep(time.Second)
	//}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
