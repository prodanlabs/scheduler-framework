# scheduler-framework
An example of CPU scheduler based on kubernetes scheduler framework

```
 git clone https://github.com/prodanlabs/scheduler-framework.git
 cd scheduler-framework/
 kubectl create configmap kube-scheduler-kubeconfig --from-file=/etc/kubernetes/kube-scheduler.kubeconfig -n kube-system
 kubectl create -f  sample-scheduler.yaml
```

<img src="https://raw.githubusercontent.com/prodanlabs/scheduler-framework/main/image/weixin.png" width="460">