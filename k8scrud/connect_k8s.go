package k8scrud

import (
    "log"
    "path/filepath"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
)

func homeDir() string {
    if h := filepath.Join("/home/kanic"); h != "" {
        return h
    }
    return "/"
}

func connect_k8s() (client *kubernetes.Clientset, err error) {
    // load kubeconfig
    kubeconfig := filepath.Join(homeDir(), ".kube", "config")
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        log.Fatalf("Error loading kubeconfig: %v", err)
        return 
    }
    // get kubernetes context
    clientset, err := kubernetes.NewForConfig(config)
    return clientset , err
}
