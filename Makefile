all : clean deps test build
.PHONY: all

# export GOOS ?= linux
# export GOARCH ?= amd64
# export CGO_ENABLED ?= 1
# export CI_BUILD_NUMBER ?= 0

LDFLAGS += -X "main.buildDate=$(shell date -u '+%Y-%m-%d %H:%M:%S %Z')"
LDFLAGS += -X "main.build=$(CI_BUILD_NUMBER)"

EXECUTABLE ?= vsphere-metrics
COMMIT ?= $(shell git rev-parse --short HEAD)

LDFLAGS += -X "main.buildCommit=$(COMMIT)"
PACKAGES = $(shell go list ./... | grep -v /vendor/)
# '-extldflags "-static" -X github.com/drone/drone/version.VersionDev=$(CI_BUILD_NUMBER)' 

watch:
	go get github.com/onsi/ginkgo/ginkgo
	ginkgo watch -r -cover
savedeps:
	rm -rf vendor Godeps
	godep save -t ./...

clean:
	go clean -v -i ./...

deps:
	go get -t -v ./...

# test:
# 	go test `go list ./... | grep -v /vendor/` -cover -ginkgo.failFast

test:
	@for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

$(EXECUTABLE): $(wildcard *.go)
	go build -ldflags '-s -w $(LDFLAGS)'

build: $(EXECUTABLE)