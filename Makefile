NAME := apito_backend

deps:
	go mod download

lint:
	golint ./...

vet:
	go vet ./...

run:
	go run main.go

test:
	go test ./...

imports:
	goimports -l -w .

build:
	go build -o ${NAME} .

api:
	swagger generate spec -o ops/api_spec-${APITO_BACKEND_VERSION}.json
	swagger serve ops/api_spec-${APITO_BACKEND_VERSION}.json

watch:
	CompileDaemon --build="go build -o ./${NAME} ." --command="./${NAME}"
