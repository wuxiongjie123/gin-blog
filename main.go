package main

import (
	"context"
	"fmt"
	"gin-blog/models"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init()  {
	setting.Setup()
	models.Setup()
	logging.Setup()
}

func main() {
	//router := gin.Default()
	//router.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err:=s.Shutdown(ctx);err!=nil{
		log.Fatal("Server Shutdown:",err)
	}

	log.Println("Server exiting")

	// 使用endless热更新,但是不推荐
	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1<<20
	//endPoint := fmt.Sprintf(":%d",setting.HTTPPort)
	////endless.NewServer 返回一个初始化的 endlessServer 对象，在 BeforeBegin
	//// 时输出当前进程的 pid，调用 ListenAndServe 将实际“启动”服务
	//server := endless.NewServer(endPoint,routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d",syscall.Getpid())
	//}
	//err:=server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}

}
