test:
	@go build ./... && go test -v -count=1 ./...

new-%:
	@mkdir ./$*
	@cd ./$* && touch readme.md $*.go $*_test.go
	@cd ./$* && echo "package $*" > $*.go
	@cd ./$* && echo "package $*" > $*_test.go

content:
	@go run content.go

.PHONY: test, content