.PHONY: binary
binart: dodo

.PHONY: dodo
dodo:
	go build -o bin/dodo ./cmd/dodo

.PHONY: clean
clean:
	rm -rf bin

.PHONY: test
test:
	go test -covermode=atomic ./...
