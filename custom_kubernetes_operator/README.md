# Go Concepts
Repo for basic tutorial-based Golang study  

---

### custom kubernetes operator
* `Controller` - loop that forevewer and observes state of some resource
* `Idempotence` - controller should be idempotent, meaning that after running it 1 up to N times, the result should be the same
* `Reconcile` - loop is going to eventually match the desired state via the process of reconciliation
* ... so Kubernetes is `eventually consistent`
* `k3d` is used as local Kubernetes cluster
* `kubebuilder` is used as a operator-building framework

#### commands
* `kubectl cluster create --config cluster.yaml`

#### reference
* [freecodecamp](https://www.freecodecamp.org/news/build-your-own-kubernetes-operators-with-go-and-kubebuilder/)