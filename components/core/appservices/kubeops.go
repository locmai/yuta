package appservices

import (
	"context"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type KubeopsAppService struct {
	KubeClientSet kubernetes.Clientset
}

func NewKubeopsAppService() *KubeopsAppService {
	config, err := rest.InClusterConfig()
	if err != nil {
		logrus.Warn("Unable to initialize KubeOps appservicem, skipping.")
		return nil
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Warn("Unable to initialize KubeOps appservicem, skipping.")
		return nil
	}
	return &KubeopsAppService{
		KubeClientSet: *clientset,
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
	// TODO Remove this Value later
	Value float64
}

func (s KubeopsAppService) Act(action Action, metadata ObjectMeta) {
	switch action {
	case ScaleAction:
		s.scaleDeployment(metadata.Name, metadata.Namespace, metadata.Value)
	case RestartAction:
		s.restartDeployment()
	case GetAction:
		s.getDeployment()
	default:
		logrus.Printf("Action requested is not implemented %s", action)
	}
}

func (s KubeopsAppService) scaleDeployment(name, namespace string, value float64) {
	ctx := context.Background()
	deploymentClient := s.KubeClientSet.AppsV1().Deployments(namespace)

	deployment, err := deploymentClient.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		logrus.Errorf("Failed to get deployment: %v", err)
	}

	parsedValue := int32(value)
	deployment.Spec.Replicas = &parsedValue

	deploymentClient.Update(ctx, deployment, metav1.UpdateOptions{})
}

func (s KubeopsAppService) restartDeployment() {
	logrus.Printf("Action requested is not implemented")
}

func (s KubeopsAppService) getDeployment() {
	logrus.Printf("Action requested is not implemented")
}
