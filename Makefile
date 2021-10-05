test:
	@go build ./... && go test -v -count=1 ./...

content:
	@go run content.go

.PHONY: test, content