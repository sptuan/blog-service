package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sptuan/blog-service/global"
	"github.com/sptuan/blog-service/internal/routers"
	"github.com/sptuan/blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	err := SetupSetting()
	if err != nil {
		log.Fatalf("[Fatal] Init config failed: %v", err)
	}
}
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}

func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second

	log.Print(global.ServerSetting)

	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	log.Print(global.AppSetting)

	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	log.Print(global.DatabaseSetting)

	return nil
}
