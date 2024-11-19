.PHONY: build
build:
	go build -o ./dbservice cmd/dbservice.go

clean:
	rm ./dbservice

run:
	./dbservice

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
