.PHONY: binary
binart: dodo

.PHONY: dodo
dodo:
	go build -o bin/dodo ./cmd/dodo

.PHONY: clean
clean:
	rm -rf bin