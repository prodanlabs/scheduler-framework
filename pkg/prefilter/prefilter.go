package prefilter

import (
	"context"

	"github.com/prodanlabs/scheduler-framework/pkg/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func IsExist(namespace string) bool {
	clientset, _ := client.Connect()
	_, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		return false
	}
	return true
}

func IsReady(namespace, podLabs string) bool {
	clientset, _ := client.Connect()
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: podLabs,
	})
	if err != nil {
		return false
	}

	if len(pods.Items) == 0 {
		return false
	}
	for _, podname := range pods.Items {
		for i := 0; i < len(podname.Status.ContainerStatuses); i++ {
			if podname.Status.ContainerStatuses[i].Ready == false {
				return false
			}
		}
	}

	return true
}
