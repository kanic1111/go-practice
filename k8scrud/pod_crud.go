package k8scrud

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// PodManager struct holds the Kubernetes clientset
type PodManager struct {
	Clientset *kubernetes.Clientset
}

// CreatePod creates a new pod with the given name and namespace
func (p *PodManager) CreatePod(rd RequestData) error {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rd.Name,
			Namespace: rd.Namespace,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  rd.Name,
					Image: rd.Image,
					Ports: []corev1.ContainerPort{
						{ContainerPort: 80},
					},
				},
			},
		},
	}

	_, err := p.Clientset.CoreV1().Pods(rd.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

// ListPods retrieves all pods in a given namespace
func (p *PodManager) ListPods(rd RequestData) ([]PodInfo, error) {
	pods, err := p.Clientset.CoreV1().Pods(rd.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var podList []PodInfo
	for _, pod := range pods.Items {
		podList = append(podList, PodInfo{
			Name:      pod.Name,
			Namespace: pod.Namespace,
		})
	}

	return podList, nil
}

// DeletePod deletes a pod by name in a given namespace
func (p *PodManager) DeletePod(rd RequestData) error {
	err := p.Clientset.CoreV1().Pods(rd.Namespace).Delete(context.TODO(), rd.Name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}

// Initialize Kubernetes client
func NewPodManager() *PodManager {
	clientset, err := connect_k8s()
	if err != nil {
		return nil
	}

	return &PodManager{Clientset: clientset}
}
