package appservices

import (
	"context"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type KubeopsAppService struct {
	KuberClientSet kubernetes.Clientset
}

func NewKubeClient() *KubeopsAppService {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return &KubeopsAppService{
		KuberClientSet: *clientset,
	}
}

type Action string

const (
	ScaleAction   Action = "scale"
	RestartAction Action = "restart"
	GetAction     Action = "get"
)

type ObjectMeta struct {
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
}

func (s KubeopsAppService) Act(action Action, metadata ObjectMeta) {
	switch action {
	case ScaleAction:
		s.scaleDeployment(metadata.Name, metadata.Namespace)
	case RestartAction:
		s.restartDeployment()
	case GetAction:
		s.getDeployment()
	default:
		logrus.Printf("Action requested is not implemented %s", action)
	}
}

func (s KubeopsAppService) scaleDeployment(name, namespace string) {
	ctx := context.Background()
	deploymentClient := s.KuberClientSet.AppsV1().Deployments(namespace)

	deployment, err := deploymentClient.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		logrus.Errorf("Failed to get deployment: %v", err)
	}

	deployment.Spec.Replicas = int32Ptr(10)

	deploymentClient.Update(ctx, deployment, metav1.UpdateOptions{})
}

func (s KubeopsAppService) restartDeployment() {
	logrus.Printf("Action requested is not implemented")
}

func (s KubeopsAppService) getDeployment() {
	logrus.Printf("Action requested is not implemented")
}

func int32Ptr(i int32) *int32 { return &i }
