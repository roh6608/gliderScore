
build:
	go build -o ./bin/gliderScore src/main.go ./src/supplementary.go

test:
	go test -v ./src

benchmark:
	go test ./src -v -bench=.

compile:
