BINARY=runSim

all: test build
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
