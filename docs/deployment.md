# Deployment - IN PROGRESS

The application is containerized by using [Podman] because it's a better than Docker. Every `docker` can be replaced by `podman`. It's already installed on Fedora and RHEL, if you're using a different OS, checkout [the official guide][2]

[Kind][3] is used to run a multi-node local kubernetes cluster and is very effective with Go and podman.

To ensure Kind uses Podman instad of Docker, set the below environment variable in the shell config file:

```zsh
export KIND_EXPERIMENTAL_PROVIDER=podman
```

Finally, [install Kubectl][4]

## Run Application Directly

+ Verify that the application runs correctly

  ```zsh
  go run ./cmd/app/main.go
  ```

+ On the browser visit http://localhost:8080/api/v1/hello

## Create Cluster

+ Navigate to the `build/` directory from the project root

+ Create a multi-node local cluster

  ```zsh
  kind create cluster --config cluster.yaml
  ```

+ Check cluster detail

  ```zsh
  kubectl cluster-info --context kind-playground-cluster
  ```

+ Get nodes

  ```zsh
  kubectl get nodes
  ```


## Deploy Application to Cluster

+ Navigate to the `deploy/` directory from the project root

+ Apply the application's deployment and service manifests

  ```zsh
  kubectl apply -f deployment.yaml
  kubectl apply -f service.yaml
  ```

+ Get pods

  ```zsh
  kubectl get pods
  ```

+ Get services (in detail)

  ```zsh
  kubectl get service -o wide
  ```

+ Port forward application. Note that the first port is the port for `localhost` and the second port is from the service manifest.

  ```
  kubectl port-forward service/flask-app-service 6001:6000
  ```

+ View application on browser: navigate to http://localhost:6001


---

[1]: https://podman.io
[2]: https://podman.io/docs/installation
[3]: https://kind.sigs.k8s.io
[4]: https://kubernetes.io/docs/tasks/tools/#kubectl
