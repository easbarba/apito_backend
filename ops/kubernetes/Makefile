up:
	kubectl apply -f /deployment/apito.yml
	kubectl apply -f ./service/apito-service.yml

down:
	kubectl delete -f ./deployment/apito.yml
	kubectl delete -f ./service/apito-service.yml

logs:
	kubectl logs deployment/apito-deployment

watch:
	kubectl get all --watch=false

podman:
	systemctl enable --now podman.socket
	ln -fs /run/podman/podman.sock /var/run/docker.sock
	podman network create podman

cluster:
	k3d cluster delete --config ./cluster/k3d.yml
	k3d cluster create --config ./cluster/k3d.yml

import:
	k3d image import -c apito localhost/easbarbosa/apito:${TAG}
