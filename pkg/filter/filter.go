/*
Copyright 2020 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package filter

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/prodanlabs/scheduler-framework/pkg/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

func GetEnvFloat(key string, defaultValue float64) float64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	i, err := strconv.ParseFloat(value, 64)
	if err != nil {
		klog.Errorf("Environment variable value for %s must be an float: %v", key, err)
	}
	return i
}

func ResourceStatus(node string) (usedCPU float64, usedMen, cpu, men int64, rateCPU, rateMen float64) {
	clientset, config := client.Connect()
	mc, err := metrics.NewForConfig(config)
	if err != nil {
		klog.Fatalln(err)
	}
	nodeInfo, err := clientset.CoreV1().Nodes().Get(context.TODO(), node, metav1.GetOptions{})
	if err != nil {
		klog.Fatalln(err)
	}
	nodeGet, err := mc.MetricsV1beta1().NodeMetricses().Get(context.TODO(), node, metav1.GetOptions{})
	if err != nil {
		klog.Errorln(err)
	}

	//CPU
	//cpu n (1 core=1000m  1m = 1000*1000n )
	CPUTotal := nodeInfo.Status.Capacity.Cpu().Value() * (1000 * 1000 * 1000)
	CPUUsed, _ := strconv.ParseFloat(strings.Replace(nodeGet.Usage.Cpu().String(), "n", "", 1), 64)

	//Memory Ki
	memoryTotal := nodeInfo.Status.Capacity.Memory().Value() / 1024
	memoryUsed := nodeGet.Usage.Memory().Value() / 1024

	return CPUUsed, memoryUsed, CPUTotal, memoryTotal, CPUUsed / float64(CPUTotal), float64(memoryUsed) / float64(memoryTotal)
}
