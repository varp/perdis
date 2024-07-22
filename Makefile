.PHONY: test

test:
	go test -v -p 4 -race ./...
