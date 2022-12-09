NAME := apito_backend

deps:
	./mvnw clean compile

lint:

vet:

test:
	./mvnw test

run: deps
	 ./mvnw spring-boot:run

run-local:
	 ./mvnw spring-boot:run -Dspring.profiles.active=local

imports:

build:

api:

watch:

image:
	sudo podman build -f Dockerfile --tag easbarbosa/apito:backend-${APITO_VERSION}

release: image
	sudo podman push easbarbosa/apito:backend-${APITO_VERSION}
