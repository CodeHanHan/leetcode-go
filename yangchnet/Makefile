test:
	@go build ./... && go test -v -count=1 ./...
	
new-%:
	@mkdir $*
	@cd ./$* && touch readme.md $*.go $*_test.go
	@cd ./$* && echo "package $*" > $*.go
	@cd ./$* && echo "package $*" > $*_test.go

list-%:./base/LinkList/LinkList.go
	@mkdir $*
	@cd ./$* && touch readme.md $*.go $*_test.go
	@cd ./$* && echo "package $*" > $*.go
	@cd ./$* && echo "package $*" > $*_test.go
	@cp $< ./$*/
	@sed -i 's/package list/package $*/g' ./$*/LinkList.go

content:
	@go run content.go

.PHONY: content