.PHONY: default test test-cover dev generate hooks

# for dev
dev:
	air -c .air.toml

doc:
	swagger generate spec -o ./api.yml && swagger validate ./api.yml 

test:
	MAIL_USER=test@outlook.com go test -race -cover ./...

install:
	go get entgo.io/ent/cmd/entc

generate: 
	entc generate ./schema --target ./ent

describe:
	entc describe ./schema

test-cover:
	MAIL_USER=test@outlook.com go test -race -coverprofile=test.out ./... && go tool cover --html=test.out

list-mod:
	go list -m -u all

tidy:
	go mod tidy

build:
	go build -ldflags "-X main.Version=0.0.1 -X 'main.BuildedAt=`date`'" -o cybertect 


lint:
	golangci-lint run --timeout 2m --skip-dirs web

hooks:
	cp hooks/* .git/hooks/