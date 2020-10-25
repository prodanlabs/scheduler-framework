package score

import (
	"context"

	"github.com/prodanlabs/scheduler-framework/pkg/client"
	"github.com/prodanlabs/scheduler-framework/pkg/filter"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

const (
	small  int64 = 0
	middle int64 = 1
	large  int64 = 2
	great  int64 = 3
)

func IsSamePod(node, namespace string, podLabs map[string]string) int64 {

	clientset, _ := client.Connect()
	for k, v := range podLabs {
		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
			LabelSelector: k + "=" + v,
		})
		if err != nil {
			klog.Error(err)
		}
		for i := 0; i < len(pods.Items); i++ {
			if node == pods.Items[i].Spec.NodeName {
				return small
			}
		}
	}

	return large
}

func CPURate(nodeName string) int64 {
	_, _, _, _, cpu, _ := filter.ResourceStatus(nodeName)

	if cpu > 0.8 {
		return small
	} else if cpu > 0.6 && cpu < 0.8 {
		return middle
	} else if cpu > 0.3 && cpu < 0.5 {
		return large
	} else {
		return great
	}
}

func MemoryRate(nodeName string) int64 {
	_, _, _, _, _, men := filter.ResourceStatus(nodeName)

	if men > 0.8 {
		return small
	} else if men > 0.6 && men < 0.8 {
		return middle
	} else if men > 0.3 && men < 0.5 {
		return large
	} else {
		return great
	}
}

func CpuCore(nodeName string) int64 {
	clientset, _ := client.Connect()
	nodeInfo, err := clientset.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		klog.Fatalln(err)
	}
	c := nodeInfo.Status.Capacity.Cpu().Value()
	if c >= 128 {
		return great
	} else if c >= 64 && c < 128 {
		return large
	} else if c >= 16 && c < 64 {
		return middle
	} else {
		return small
	}
}
