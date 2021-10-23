test:
	@go build ./... && go test -v -count=1 ./...

content:
	@cd ./yangchnet && make content
	@cd ./HT-CHEN520 && make content
	@go run content.go

lc-new-%:
	@cd ./yangchnet && make new-$*

lc-list-%:yangchnet/base/LinkList/LinkList.go
	@cd ./yangchnet && make list-$*

cht-new-%:
	@cd ./HT-CHEN520 && make new-$*

.PHONY: test, content