module github.com/prodanlabs/scheduler-framework

go 1.14

require (
	k8s.io/api v0.19.1
	k8s.io/apimachinery v1.19.1
	k8s.io/client-go v0.19.1
	k8s.io/klog/v2 v2.2.0
	k8s.io/kubernetes v1.19.1
	k8s.io/metrics v0.0.0
)

replace (
	k8s.io/api => /opt/workspaces/kubernetes/vendor/k8s.io/api
	k8s.io/apiextensions-apiserver => /opt/workspaces/kubernetes/vendor/k8s.io/apiextensions-apiserver
	k8s.io/apimachinery => /opt/workspaces/kubernetes/vendor/k8s.io/apimachinery
	k8s.io/apiserver => /opt/workspaces/kubernetes/vendor/k8s.io/apiserver
	k8s.io/cli-runtime => /opt/workspaces/kubernetes/vendor/k8s.io/cli-runtime
	k8s.io/client-go => /opt/workspaces/kubernetes/vendor/k8s.io/client-go
	k8s.io/cloud-provider => /opt/workspaces/kubernetes/vendor/k8s.io/cloud-provider
	k8s.io/cluster-bootstrap => /opt/workspaces/kubernetes/vendor/k8s.io/cluster-bootstrap
	k8s.io/code-generator => /opt/workspaces/kubernetes/vendor/k8s.io/code-generator
	k8s.io/component-base => /opt/workspaces/kubernetes/vendor/k8s.io/component-base
	k8s.io/cri-api => /opt/workspaces/kubernetes/vendor/k8s.io/cri-api
	k8s.io/csi-translation-lib => /opt/workspaces/kubernetes/vendor/k8s.io/csi-translation-lib
	k8s.io/kube-aggregator => /opt/workspaces/kubernetes/vendor/k8s.io/kube-aggregator
	k8s.io/kube-controller-manager => /opt/workspaces/kubernetes/vendor/k8s.io/kube-controller-manager
	k8s.io/kube-proxy => /opt/workspaces/kubernetes/vendor/k8s.io/kube-proxy
	k8s.io/kube-scheduler => /opt/workspaces/kubernetes/vendor/k8s.io/kube-scheduler
	k8s.io/kubectl => /opt/workspaces/kubernetes/vendor/k8s.io/kubectl
	k8s.io/kubelet => /opt/workspaces/kubernetes/vendor/k8s.io/kubelet
	k8s.io/legacy-cloud-providers => /opt/workspaces/kubernetes/vendor/k8s.io/legacy-cloud-providers
	k8s.io/metrics => /opt/workspaces/kubernetes/vendor/k8s.io/metrics
	k8s.io/sample-apiserver => /opt/workspaces/kubernetes/vendor/k8s.io/sample-apiserver
)
