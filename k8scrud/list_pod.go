package k8scrud

import (
	"context"
	"log"
        "go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        "github.com/gin-gonic/gin"
        "net/http"
)

// c value is defined from the kuberentes.Clientset "*" means access from memory so you don't need to create it again
func List_pod(c *gin.Context){
    // connect_k8s
    client, err := connect_k8s()
    log.Print("test:", c.GetHeader("Content-Type"))
    if err != nil {
        log.Fatalf("Error connecting k8s: %v", err)
    }
    // new logger
    logger := zap.NewExample()
    defer logger.Sync()
    sugar := logger.Sugar()
    // Get the list of pods in the default namespace
    pods, err := client.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        log.Fatalf("Error fetching pods: %v", err)
    }
    // list pod 
    var podList []PodInfo
    for _, pod := range pods.Items {
        if pod.Status.Phase == "Running" {
            podList = append(podList, PodInfo{
                    Name:      pod.Name,
                    Namespace: pod.Namespace,
                })
            sugar.Infow("get pod",
                zap.Any("Running Pods:", PodInfo{
			Name:      pod.Name,
			Namespace: pod.Namespace,
		}),
            )
        }
    }
    c.JSON(http.StatusOK, gin.H{"message": podList})
}
