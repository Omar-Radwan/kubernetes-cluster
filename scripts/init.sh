REGISTRY_PORT="5000"

minikube start

minikube addons enable registry

kubectl port-forward --namespace kube-system service/registry "$REGISTRY_PORT":80 &

[ -n "$(docker images -q alpine)" ] || docker pull alpine

docker run --rm -it --network=host alpine ash -c "apk add socat && socat TCP-LISTEN:$REGISTRY_PORT,reuseaddr,fork TCP:$(minikube ip):$REGISTRY_PORT" 

minikube addons enable ingress


DNS_RECORD="$(minikube ip) simple-backend.io simple-frontend.io"
HOSTS_PATH="/etc/hosts"

grep -Fxq "$DNS_RECORD" "$HOSTS_PATH" || echo "$DNS_RECORD" | sudo tee -a "$HOSTS_PATH"