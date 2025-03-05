package k8scrud


type PodInfo struct {
	Name      string
	Namespace string
}

type DeploymentInfo struct {
        Name      string
        Namespace string
        Replicas  int32
}
