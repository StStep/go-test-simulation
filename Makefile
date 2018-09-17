BINARY=runSim

all: format test build
format:
	find $(GOPATH)/src/github.com/StStep/go-test-simulation/cmd/ -iname "*.go" -exec gofmt -w {} \;
	find $(GOPATH)/src/github.com/StStep/go-test-simulation/internal/ -iname "*.go" -exec gofmt -w {} \;
build:
	go build github.com/StStep/go-test-simulation/cmd/runSim
test:
	go test -v ./...
clean:
	go clean
	rm -f $(BINARY)
run:
	go build github.com/StStep/go-test-simulation/cmd/runSim
	./$(BINARY)
deps:
	dep ensure
