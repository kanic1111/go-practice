package k8scrud

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
)

// PodManager struct holds the Kubernetes clientset
type K8sClient struct {
	Clientset *kubernetes.Clientset
}

// CreatePod creates a new pod with the given name and namespace
func (kc *K8sClient) CreatePod(rd RequestData) error {
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

	_, err := kc.Clientset.CoreV1().Pods(rd.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

// ListPods retrieves all pods in a given namespace
func (kc *K8sClient) ListPods(rd RequestData) ([]PodInfo, error) {
	pods, err := kc.Clientset.CoreV1().Pods(rd.Namespace).List(context.TODO(), metav1.ListOptions{})
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
func (kc *K8sClient) DeletePod(rd RequestData) error {
	err := kc.Clientset.CoreV1().Pods(rd.Namespace).Delete(context.TODO(), rd.Name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (kc *K8sClient) ListDeploy(rd RequestData) ([]DeploymentInfo, error) {
	deployments, err := kc.Clientset.AppsV1().Deployments(rd.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var DeploymentList []DeploymentInfo
	for _, deployment := range deployments.Items {
		DeploymentList = append(DeploymentList, DeploymentInfo{
			Name:      deployment.Name,
			Namespace: deployment.Namespace,
			Replicas:  *deployment.Spec.Replicas,
		})
	}
	return DeploymentList, err
}

func (kc *K8sClient) CreateDeployment(rd RequestData) error {
	dc := kc.Clientset.AppsV1().Deployments(rd.Namespace)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rd.Name,
			Namespace: rd.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &rd.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": rd.Name},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": rd.Name},
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
			},
		},
	}

	_, err := dc.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (kc *K8sClient) DeleteDeploy(rd RequestData) error {
	dc := kc.Clientset.AppsV1().Deployments(rd.Namespace)
	err := dc.Delete(context.TODO(), rd.Name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (kc *K8sClient) UpdateDeploy(rd RequestData) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		dc := kc.Clientset.AppsV1().Deployments(rd.Namespace)
		result, getErr := dc.Get(context.TODO(), rd.Name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}
		result.Spec.Replicas = &rd.Replicas
		result.Spec.Template.Spec.Containers[0].Image = rd.Image
		_, updateErr := dc.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		return retryErr
	}
	return nil
}

// Initialize Kubernetes client
func NewK8sClient() *K8sClient {
	clientset, err := connect_k8s()
	if err != nil {
		return nil
	}
	return &K8sClient{Clientset: clientset}
}
