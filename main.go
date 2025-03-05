package main

import (
    "go-test/k8scrud"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    rt := addroutes()
    rt.Run(":8080")
    s := &http.Server{
	Addr:           ":8080",
	Handler:        rt,
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
    api.Handle("GET", "/listpod", k8scrud.List_pod)
    api.Handle("GET", "/listdeploy", k8scrud.List_deploy)
    return router
}
