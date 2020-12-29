package main

import (
	"context"
	"eatwhat/conf"
	Log "eatwhat/log"
	"eatwhat/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)
//int server
func initServer(rt *gin.Engine){
	srv := &http.Server{
		Addr: conf.APP_HOST+conf.APP_PORT,
		Handler: rt,
		ReadTimeout: conf.READ_TIMEOUT*time.Second,
		WriteTimeout: conf.WRITE_TIMEOUT*time.Second,
		//MaxHeaderBytes: 1<<20,
	}
	go func() {
		//start server
		if err:=srv.ListenAndServe();err!=nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n",err)
		}
	}()
	//handler signal
	quit := make(chan os.Signal)
	signal.Notify(quit,os.Interrupt,syscall.SIGTERM)
	c:= <- quit
	log.Println("server is shutting down:",c)
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	if err:=srv.Shutdown(ctx);err!=nil {
		log.Fatal("Server Shutdown:",err)
	}
	log.Println("Server exiting")
}

func main(){
	//init log
	logger:=new(Log.Log)
	logger.InitLog()
	//register router
	r := new(router.Router)
	routerHandler := r.InitRouter(logger)
	//start server
	initServer(routerHandler)

}