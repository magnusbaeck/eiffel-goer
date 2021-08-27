export RELEASE_VERSION ?= $(shell git describe --always)
export DOCKER_REGISTRY ?= registry.nordix.org/eiffel
export DEPLOY ?= goer

all: test build start
gen:
	go generate ./...
build: gen
	go get github.com/ahmetb/govvv
	govvv build -o bin/goer ./cmd/goer
clean:
	rm ./bin/* || true
	docker-compose --project-directory . -f deploy/$(DEPLOY)/docker-compose.yml rm || true
	docker volume rm goer-volume || true
test: gen
	go test -cover -timeout 30s -race $(shell go list ./... | grep -v test) 

# Start a development docker with a database that restarts on file changes.
start:
	docker-compose --project-directory . -f deploy/$(DEPLOY)/docker-compose.yml up
stop:
	docker-compose --project-directory . -f deploy/$(DEPLOY)/docker-compose.yml down

# Build a docker using the production Dockerfile
docker:
	docker build -t $(DOCKER_REGISTRY)/$(DEPLOY):$(RELEASE_VERSION) -f ./deploy/$(DEPLOY)/Dockerfile .
push:
	docker push $(DOCKER_REGISTRY)/$(DEPLOY):$(RELEASE_VERSION)
