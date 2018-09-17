BINARY=runSim
PACKAGE=github.com/StStep/go-test-simulation

all: format test build
format:
	find $(GOPATH)/src/$(PACKAGE)/cmd/ -iname "*.go" -exec gofmt -w {} \;
	find $(GOPATH)/src/$(PACKAGE)/internal/ -iname "*.go" -exec gofmt -w {} \;
build:
	go build $(PACKAGE)/cmd/$(BINARY)
test:
	go test -v ./...
clean:
	go clean
	rm -f $(BINARY)
run:
	go build $(PACKAGE)/cmd/$(BINARY)
	./$(BINARY)
deps:
	dep ensure
