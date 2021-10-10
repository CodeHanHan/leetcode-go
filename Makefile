test:
	@go build ./... && go test -v -count=1 ./...

content:
	@go run content.go

new-%:
	mkdir -p ./yangchnet/$*
	cd ./yangchnet/$* && touch readme.md $*.go $*_test.go

.PHONY: test, content