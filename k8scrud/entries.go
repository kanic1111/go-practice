package k8scrud

import (
	//	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
	Namespace string `json:"namespace"` // JSON key should match the expected input
	Name      string `json:"name,omitempty"`
	Image     string `json:"image,omitempty"`
}

// c value is defined from the kuberentes.Clientset "*" means access from memory so you don't need to create it again
func List_pod(c *gin.Context) {
	podManager := NewPodManager()
	rd := get_value(c)
	podList, err := podManager.ListPods(rd)
	if err != nil {
		log.Fatalf("Failed to create pod: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"message": podList})
}

func Create_pod(c *gin.Context) {
	podManager := NewPodManager()
	rd := get_value(c)
	err := podManager.CreatePod(rd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Create Pod failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "create pod successful"})
}

func Delete_pod(c *gin.Context) {
	podManager := NewPodManager()
	rd := get_value(c)
	err := podManager.DeletePod(rd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Delete Pod failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete pod successful"})

}
func get_value(c *gin.Context) RequestData {
	var requestData RequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		requestData.Namespace = "default"
		return requestData
	}
	return requestData
}
