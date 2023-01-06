back:
	${RUNNER} build -f ./backend/Dockerfile --tag easbarbosa/apito:backend-${TAG}

front:
	 ${RUNNER} build -f ./frontend/Dockerfile --tag easbarbosa/apito:frontend-${TAG}

release:
	${RUNNER} push easbarbosa/apito:backend-${TAG}
	${RUNNER} push easbarbosa/apito:frontend-${TAG}
