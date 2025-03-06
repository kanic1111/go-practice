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
	crud_api := router.Group("/k8s")
	list_api := crud_api.Group("/list/:namespace")
	list_api.Handle("GET", "/pod", k8scrud.List_pod)
	list_api.Handle("GET", "/deploy", k8scrud.List_deploy)
	crud_api.Handle("POST", "/createpod", k8scrud.Create_pod)
	crud_api.Handle("POST", "/createdeploy", k8scrud.Create_deploy)
	crud_api.Handle("DELETE", "/deletepod", k8scrud.Delete_pod)
	crud_api.Handle("DELETE", "/deletedeploy", k8scrud.Delete_deploy)
	crud_api.Handle("PATCH", "/updatedeploy", k8scrud.Update_deploy)
	return router
}
