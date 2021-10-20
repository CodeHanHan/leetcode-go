test:
	@go build ./... && go test -v -count=1 ./...

content:
	@go run content.go

new-%:
	@mkdir -p ./yangchnet/$*
	@cd ./yangchnet/$* && touch readme.md $*.go $*_test.go
	@cd ./yangchnet/$* && echo "package $*" > $*.go
	@cd ./yangchnet/$* && echo "package $*" > $*_test.go

new-list-%:yangchnet/base/LinkList/LinkList.go
	@mkdir -p ./yangchnet/$*
	@cd ./yangchnet/$* && touch readme.md $*.go $*_test.go
	@cd ./yangchnet/$* && echo "package $*" > $*.go
	@cd ./yangchnet/$* && echo "package $*" > $*_test.go
	@cp $< ./yangchnet/$*/
	@sed -i 's/package list/package $*/g' ./yangchnet/$*/LinkList.go

.PHONY: test, content