A work in progress k8s operator

The controller logic is, it checks if a service and deployment is present, if they are not it deploys those, if they are already then do nothing.

`make install` installs CRD onto a real GKE cluster, it also puts a deployment of a simple hello pod, and a service into the cluster.

`make run` runs controller on the localhost outside of the cluster.

`make test` runs the traveller_test.go in the controllers directory.

```
kubectl get travellers.my.domain traveller-sample -o yaml
apiVersion: my.domain/v1alpha1
kind: Traveller
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"my.domain/v1alpha1","kind":"Traveller","metadata":{"annotations":{},"name":"traveller-sample","namespace":"default"},"spec":{"foo":"bar"}}
  creationTimestamp: "2022-03-29T15:17:19Z"
  generation: 1
  name: traveller-sample
  namespace: default
  resourceVersion: "1699772"
  uid: 3e3b1071-ca80-4731-9827-8f4afd4572f2
spec:
  foo: bar
```



kubebuilder version
Version: main.version{KubeBuilderVersion:"3.3.0", KubernetesVendor:"1.23.1", GitCommit:"47859bf2ebf96a64db69a2f7074ffdec7f15c1ec", BuildDate:"2022-01-18T17:03:29Z", GoOs:"linux", GoArch:"amd64"}

go version go1.17.3 linux/amd64