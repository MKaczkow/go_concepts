# Go Concepts
Repo for basic tutorial-based Golang study  

---

### todo
- [x] maybe this should be a separate repo
- [ ] mini-docs / 'tutorial for tutorial' to speed up starting and preparing stuff after break
- [ ] run the operator on actual (k3d, so not really actual) cluster
- [x] move CI into project-level CI
- [ ] fix CI (Go version stuff)

### custom kubernetes operator
* `Controller` - loop that runs forever and regulates the state of a system
* `Controller Pattern` - a controller tracks at least one Kubernetes resource type. Those have to have a specification field, which represents the desired state. The controllers are basically responsible for moving the current state (from the field) closer to desired state.
* `Operator Pattern` - allows modification the cluster state, without modifying the code of Kubernetes itself by linking controllers to one or more custom resources.
* `Idempotence` - controller should be idempotent, meaning that after running it 1 up to N times, the result should be the same
* `Reconcile` - loop is going to eventually match the desired state via the process of reconciliation; implementing this logic is v. important, actual business logic goes there, implementing `Reconciler` interface
* ... so Kubernetes is `eventually consistent`
* `k3d` is used as local Kubernetes cluster
* `kubebuilder` is used as a operator-building framework
* ... actually it uses `controller-runtime` under the hood
* ... and another options is using `Operator SDK`
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
* [controllers - control loop ref docs](https://kubernetes.io/docs/concepts/architecture/controller/)
* [operators ref docs](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)
* [google cloud best practices from 2018](https://cloud.google.com/blog/products/containers-kubernetes/best-practices-for-building-kubernetes-operators-and-stateful-apps)
* [custom resources docs](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)
* [custom resource definition docs](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/)