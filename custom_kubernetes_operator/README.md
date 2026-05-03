# Go Concepts
Repo for basic tutorial-based Golang study  

---

### todo
- [ ] maybe this should be a separate repo
- [ ] mini-docs / 'tutorial for tutorial' to speed up starting and preparing stuff after break
- [x] move CI into project-level CI
- [ ] fix CI (Go version stuff)

### custom kubernetes operator
* `Controller` - loop that forevewer and observes state of some resource
* `Idempotence` - controller should be idempotent, meaning that after running it 1 up to N times, the result should be the same
* `Reconcile` - loop is going to eventually match the desired state via the process of reconciliation (implementing this logic is important)
* ... so Kubernetes is `eventually consistent`
* `k3d` is used as local Kubernetes cluster
* `kubebuilder` is used as a operator-building framework
* entry points (`cmd/main.go`) exposes some useful stuff, like metrics

#### kubebuilder
* `Spec` and `Status` are really important, they basically define what the resource will be showing and doing
* whatever we see on kubernetes resources is defined in `TypeMeta`, for example `Kind`
* ... there is also an object metadata, in `ObjectMeta`

#### controller-runtime
* ... this is also important, don't know why yet

#### commands
* `k3d version` (sanity check)
* `k3d cluster create --config cluster.yaml`
* `kubectl cluster-info` (verify creation)
* `k3d cluster delete --all`

#### reference
* [freecodecamp](https://www.freecodecamp.org/news/build-your-own-kubernetes-operators-with-go-and-kubebuilder/)
* [act - useful for testing CI (github)](https://github.com/nektos/act)
* [act - useful for testing CI (docs)](https://nektosact.com/)