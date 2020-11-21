.PHONY: test fmt showcover lint 

test:
	@make lint && go test -v -coverprofile cp.out ./...

showcover:
	go tool cover -html=cp.out
	
fmt:
	gofmt -s -w .

lint:
	golangci-lint run
