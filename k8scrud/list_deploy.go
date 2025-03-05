package k8scrud

import (
	"context"
        "log"
        "go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        "github.com/gin-gonic/gin"
        "net/http"
)


func List_deploy(c *gin.Context) {
    // connect_k8s
    client, err := connect_k8s()
    if err != nil {
        log.Fatalf("Error connecting k8s: %v", err)
    }
    logger := zap.NewExample()
    defer logger.Sync()
    sugar := logger.Sugar()
    // Get the list of pods in the default namespace
        deployments, err := client.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})
        if err != nil {
                log.Fatalf("Failed to list deployments: %v", err)
        }
    var DeploymentList []DeploymentInfo
    for _, deployment := range deployments.Items {
            DeploymentList = append(DeploymentList, DeploymentInfo{
                Name: deployment.Name,
                Namespace: deployment.Namespace,
                Replicas: *deployment.Spec.Replicas,
            })
            sugar.Infow("", zap.Any("Deployment:", DeploymentInfo{
                Name: deployment.Name,
                Namespace: deployment.Namespace,
                Replicas: *deployment.Spec.Replicas,
                }),
            )
        }
    c.JSON(http.StatusOK, gin.H{"message": DeploymentList})
}
