.PHONY: HT-CHEN520.test
HT-CHEN520.test:
	@go build ./... && go test -v -count=1 ./...

.PhONY: new-%
new-%:
	@mkdir ./$*
	@cd ./$* && touch readme.md $*.go $*_test.go
	@cd ./$* && echo "package $*" > $*.go
	@cd ./$* && echo "package $*" > $*_test.go

.PHONY: HT-CHEN520.content
HT-CHEN520.content:
	@$(GO) run content.go -d ./HT-CHEN520 -o ./HT-CHEN520


