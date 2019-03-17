export GO111MODULE = on

.PHONY: default test test-cover dev

# for dev
dev:
	fresh

test: export GO_ENV=test
test:
	go test -cover ./...

test-cover: export GO_ENV=test
test-cover:
	go test -race -coverprofile=test.out ./... && go tool cover --html=test.out

build:
	packr2
	go build -tags netgo -o cyber-tect 

clean:
	packr2 clean

release:
	go mod tidy