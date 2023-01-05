back:
	podman build -f ./backend/Dockerfile --tag easbarbosa/apito:backend-${TAG}

front:
	 podman build -f ./frontend/Dockerfile --tag easbarbosa/apito:frontend-${TAG}

release:
	podman push easbarbosa/apito:backend-${TAG}
	podman push easbarbosa/apito:frontend-${TAG}
