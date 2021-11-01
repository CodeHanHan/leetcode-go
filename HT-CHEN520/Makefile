test:
	@go build ./... && go test -v -count=1 ./...

new-%:
	@mkdir ./$*
	@cd ./$* && touch readme.md $*.go $*_test.go
	@cd ./$* && echo "package $*" > $*.go
	@cd ./$* && echo "package $*" > $*_test.go

content:
	@go run content.go

clean:
	@find . -name '*.exe' -type f -print -exec rm -rf {} \;
	@find . -name '*.o' -type f -print -exec rm -rf {} \;
	@find . -name '*.dll' -type f -print -exec rm -rf {} \;

.PHONY: test, content, clean