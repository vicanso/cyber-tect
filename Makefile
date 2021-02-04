.PHONY: default test test-cover dev generate hooks

# for dev
dev:
	nodemon main.go

doc:
	swagger generate spec -o ./api.yml && swagger validate ./api.yml 

test:
	go test -race -cover ./...

generate: 
	go generate ./ent

describe:
	entc describe ./ent/schema

test-cover:
	go test -race -coverprofile=test.out ./... && go tool cover --html=test.out

list-mod:
	go list -m -u all

tidy:
	go mod tidy

build:
	packr2
	go build -ldflags "-X main.Version=0.0.1 -X 'main.BuildedAt=`date`'" -o cybertect 

clean:
	packr2 clean

lint:
	golangci-lint run --timeout 2m --skip-dirs web

hooks:
	cp hooks/* .git/hooks/