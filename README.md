# kubetest

Learning how to use Kubernetes using Minikube.

## Resources

[Hello Minikube](https://kubernetes.io/docs/tutorials/hello-minikube) (kubernetes.io)

[Running Kubernetes Locally via Minikube](https://kubernetes.io/docs/setup/minikube/) (kubernetes.io)

[Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/) (kubernetes.io)

[Learn Kubernetes in Under 3 Hours: A Detailed Guide to Orchestrating Containers](https://medium.freecodecamp.org/learn-kubernetes-in-under-3-hours-a-detailed-guide-to-orchestrating-containers-114ff420e882) (freecodecamp.org)

[Learning Kubernetes: Getting Started with Minikube](https://sweetcode.io/learning-kubernetes-getting-started-minikube/) (sweetcode.io)

[Setting up Ingress on Minikube](https://medium.com/@Oskarr3/setting-up-ingress-on-minikube-6ae825e98f82) (medium.com)

## Terminology

* Pod - the smallest unit of the Kubernetes object model. Typically is a single container, but can be multiple containers that need to share resources.
* Deployment - orchestrates the creation, deletion, and updating of Pods.
  * Technically there's an object in between--ReplicaSets--but this is abstracted, and the Kubernetes docs [recommend using Deployments in most cases](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/#when-to-use-a-replicaset).
* Service - defines a set of Pods (using a Label Selector) and how to access them (e.g. exposing ports). Can be configured as a load balancer.
* Ingress - exposes HTTP(S) routes from outside the cluster to services within the cluster.

## Getting Started

Install kubectl, minikube, virtualbox

```
brew install kubernetes-cli
brew cask install minikube virtualbox
```

Start minikube in a new VirtualBox instance:

```
minikube start --vm-driver=virtualbox
```

Add an entry to `/etc/hosts` for ingress:

```
# IP of minikube vm (use `minikube ip`)
192.168.99.101 kube.test
```

## Helpful Commands

Open a minikube dashboard to visually inspect the cluster:

```
minikube dashboard
```

Open a service in the browser (not sure what happens if there are multiple ports defined):

```
minikube service xxx
```

Kill the minikube vm:

```
minikube stop
```

Submitting a specification to the cluster:

```
kubectl create -f xxx.yaml
```

Updating an existing specification:

```
kubectl apply -f xxx.yaml
```

Deleting a specification:

```
kubectl delete pod xxx
kubectl delete svc xxx
kubectl delete deployment xxx
kubectl delete ing xxx
```

Inspecting pods, deployments, and services in the cluster:

```
kubectl get pod
kubectl describe pod xxx

kubectl get svc
kubectl describe svc xxx

kubectl get deployments
kubectl describe deployments.apps xxx

kubectl get ingress
kubectl describe ing xxx
```

Creating a port-forward to a specific Pod (could be handy for testing):

```
kubectl port-forward xxx 8080:80
# ports are [pod]:[host]
```

Zero-downtime deployment / rollback:

```
kubectl apply -f xxx.yaml --record
kubectl rollout status deployment xxx

# rollback
kubectl rollout history deployment xxx
kubectl rollout undo deployment xxx --to-revision=N
```

Using local docker images:

```
eval $(minikube docker-env)
# AFAIK submitting an unchanged spec is a no-op, so I've been incrementing a version tag for every build
docker build -t xxx:v1 .
# use the tag in the pod/deployment spec
```

Tailing logs from multiple pods using [kubetail](https://github.com/johanhaleby/kubetail):

```
kubetail xxx
```
