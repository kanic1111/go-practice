package main

import (
	"go-test/k8scrud"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	rt := addroutes()
	rt.Run(":8080")
	s := &http.Server{
		Addr:    ":8080",
		Handler: rt,
	}
	go func() {
		s.ListenAndServe()
	}()
}

func addroutes() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(gin.Logger())
	api := router.Group("/api")
	api.Handle("POST", "/listpod", k8scrud.List_pod)
	api.Handle("GET", "/listdeploy", k8scrud.List_deploy)
	api.Handle("POST", "/createpod", k8scrud.Create_pod)
	api.Handle("POST", "/deletepod", k8scrud.Delete_pod)
	return router
}
