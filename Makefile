export GO111MODULE = on

.PHONY: default test test-cover dev

# for dev
dev:
	fresh

doc:
	swagger generate spec -o ./api.yml && swagger validate ./api.yml 

test: export GO_ENV=test
test:
	go test -cover ./...

test-cover: export GO_ENV=test
test-cover:
	go test -race -coverprofile=test.out ./... && go tool cover --html=test.out

list-mod:
	go list -m -u all

build:
	packr2
	go build -ldflags "-X main.Version=0.0.1 -X 'main.BuildAt=`date`' -X 'main.GO=`go version`'" -o cybertect 

clean:
	packr2 clean

lint:
	golangci-lint run
