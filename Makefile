GOOS?=lniux
GOARCH?=amd64
NAMESPACE?=lucchmielowski
IMAGE_NAMESPACE?=$(NAMESPACE)
APP=testbed
COMMIT=`git rev-parse --short HEAD`
TAG?=dev
BUILD?=-dev
CWD=$(PWD)

all: binaries

bindir:
	@mkdir -p bin

binaries: daemon
	@echo " -> Built $(TAG) version ${COMMIT} (${GOOS}/${GOARCH})"

daemon: bindir
	@>&2 echo " -> building daemon ${COMMIT}${BUILD}"
	@cd cmd/$(APP) && CGO_ENABLED=0 go build -installsuffix cgo -ldflags "-w -X github.com/$(REPO)/version.GitCommit=$(COMMIT) -X github.com/$(REPO)/version.Build=$(BUILD)" -o ../../bin/$(APP) .

docker-generate:
	@echo "** This uses a separate Dockerfile (Dockerfile.dev) **"
	@docker build -t $(APP)-dev -f Dockerfile.dev .
	@docker run --rm -w /go/src/github.com/$(NAMESPACE)/$(APP) $(APP)-dev sh -c "make generate; find api -name \"*.pb.go\" | tar -T - -cf -" | tar -xvf -

install:
	@install -D -m 755 cmd/$(APP)/$(APP) /usr/local/bin

clean:
	@rm -rf bin/

.PHONY: bindir binaries daemon docker-generate install clean