# Table of contents

- [Kubernetes Cluster](#kubernetes-cluster)
  - [Overview](#overview)
  - [Demo](#demo)
    - [Backend Connected](#backend-connected)
    - [Mongo Connected](#mongo-connected)
    - [Redis Connected](#redis-connected)
  - [Requirements](#requirements)
  - [Cluster Diagram](#cluster-diagram)
  - [Create Cluster](#create-cluster)
    - [Using Scripts Inside `scripts` Directory](#using-scripts-inside-scripts-directory)
    - [Manually](#manually)
    

# Kubernetes Cluster

## Overview

Kubernetes cluster that shows how to deploy and connect Go backend, React frontend, MongoDB, Redis cache with each other.

After running the cluster you can access the frontend by going to [simple-frontend.io](simple-frontend.io) in your browser.

If you can increment the counter in:

- [simple-frontend.io](http://simple-frontend.io) then react frontend and go backend are successfully connected to each other.

- [simple-frontend.io/mongo](http://simple-frontend.io/mongo) then go backend and mongoDB are successfully connected to each other.

- [simple-frontend.io/redis](http://simple-frontend.io/redis) then go backend and redis are successfully connected to each other.

## Demo

### Backend Connected

![](readme-assets/backend-connected.gif)

### Mongo Connected

![](readme-assets/mongo-connected.gif)

### Redis Connected

![](readme-assets/redis-connected.gif)

## Requirements

1. Minikube
2. Kubectl
3. Docker
4. Git

## Cluster Diagram

![](readme-assets/diagram.png)

## Create Cluster

You can create the cluster:

1. [Using scripts inside `scripts` directory](#using-scripts-inside-`scripts`-directory)
2. [Manually](#manually)

### Using Scripts Inside `scripts` Directory

Run the following command

```bash
bash scripts/init.sh && bash scripts/build-and-push.sh && bash scripts/apply-config-files.sh
```

Explanation:

1. `scripts/init.sh`: starts minikube, enables required addons, runs local container image registry on port `5000`
2. `scripts/build-and-push.sh`: builds the required images and pushes them to local container image registry
3. `scripts/apply-config-files.sh`: applies configuration files in `configuration` folder

If you want to delete all created resources by these scripts you can use use the following command `bash scripts/clean.sh`

### Manually

Follow the following steps to create the cluster from scratch using `minikube`, `docker`, `kubectl` and `git`

1. Start minikube

    ``` bash
    minikube start
    ```

2. To make the development process easier **(Optional)**:

    1. Create in-cluster container image registry

        ``` bash
        REGISTRY_PORT=5000
        
        minikube start
        
        minikube addons enable registry
        
        kubectl port-forward --namespace kube-system service/registry "$REGISTRY_PORT":80 &
        
        [ -n "$(docker images -q alpine)" ] || docker pull alpine
        
        docker run --rm -it --network=host alpine ash -c "apk add socat && socat TCP-LISTEN:$REGISTRY_PORT,reuseaddr,fork TCP:$(minikube ip):$REGISTRY_PORT" 
        ```

    2. Push images to in-cluster registry

        ```bash
        PORT=5000
        BACKEND_IMAGE_TAG="localhost:$PORT/simple-go-backend"
        FRONTEND_IMAGE_TAG="localhost:$PORT/simple-react-frontend"
        
        docker build backend -t "$BACKEND_IMAGE_TAG"
        docker push "$BACKEND_IMAGE_TAG"
        
        docker build frontend -t "$FRONTEND_IMAGE_TAG"
        docker push "$FRONTEND_IMAGE_TAG"
        ```

    If you want to skip the step of creating in-cluster registry and pushing images to it, you need to use the image uploaded previously do dockerhub. You have to replace:

    `localhost:5000/simple-react-frontend` in `connecting-frontend-backend/frontend-deployment.yaml` with `oradwan/simple-react-frontend`

    and 

    `localhost:5000/simple-go-backend` in `connecting-frontend-backend/backend-deployment.yaml` with `oradwan/simple-go-backend`

3. Enable ingress addon 

    ``` bash
    minikube addons enable ingress
    ```


4. Add DNS records of ingress endpoint to `/etc/hosts` file using the following command

    ``` bash
    DNS_RECORD="$(minikube ip) simple-backend.io simple-frontend.io"
    HOSTS_PATH="/etc/hosts"
    
    grep -Fxq "$DNS_RECORD" "$HOSTS_PATH" || echo "$DNS_RECORD" | sudo tee -a "$HOSTS_PATH"
    ```
    
5. Create the persistent volume, persistent volume claim, config map

    ``` bash
    kubectl apply -f configuration/pv-pvc-cm-ingress
    ```

6. Create services, deployments

    ``` bash
    kubectl apply -f configuration/service-deployment
    ```

7. View minikube dashboard in the browser **(Optional)**

    ``` bash
    minikube dashboard
    ```

8. You can now access the frontend by openning [simple-frontend.io](http://simple-frontend.io) in the browser

