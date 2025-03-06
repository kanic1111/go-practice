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
	Replicas  int32  `json:"replicas,omitempty"`
}

// c value is defined from the kuberentes.Clientset "*" means access from memory so you don't need to create it again
func List_pod(c *gin.Context) {
	kc := NewK8sClient()
	rd := get_value(c)
	podList, err := kc.ListPods(rd)
	if err != nil {
		log.Fatalf("Failed to create pod: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"message": podList})
}

func Create_pod(c *gin.Context) {
	kc := NewK8sClient()
	rd := get_value(c)
	err := kc.CreatePod(rd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "create pod successful"})
}

func Delete_pod(c *gin.Context) {
	kc := NewK8sClient()
	rd := get_value(c)
	err := kc.DeletePod(rd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete pod successful"})

}

func List_deploy(c *gin.Context) {
	kc := NewK8sClient()
	rd := get_value(c)
	dl, err := kc.ListDeploy(rd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"get deployment": dl})

}

func Create_deploy(c *gin.Context) {
	kc := NewK8sClient()
	rd := get_value(c)
	err := kc.CreateDeployment(rd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "create deployment successful", "deployment_name": rd.Name})
}

func Delete_deploy(c *gin.Context) {
	kc := NewK8sClient()
	rd := get_value(c)
	err := kc.DeleteDeploy(rd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete deployment successful", "deployment_name": rd.Name})
}

func Update_deploy(c *gin.Context) {
	kc := NewK8sClient()
	rd := get_value(c)
	err := kc.UpdateDeploy(rd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update deployment successful", "deployment_name": rd.Name})
}

func get_value(c *gin.Context) RequestData {
	var requestData RequestData
	requestData.Namespace = c.Param("namespace")
	if err := c.ShouldBindJSON(&requestData); err != nil {
		return requestData
	}
	return requestData
}
