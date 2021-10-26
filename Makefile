USERDIR = $(shell ls -l | grep ^d | awk '{print $$9}')

test:
	@for dir in ${USERDIR}; do make -C $${dir} test; done

content:
	@for dir in ${USERDIR}; do make -C $${dir} content; done

lc-new-%:
	@cd ./yangchnet && make new-$*

lc-list-%:yangchnet/base/LinkList/LinkList.go
	@cd ./yangchnet && make list-$*

cht-new-%:
	@cd ./HT-CHEN520 && make new-$*

.PHONY: test, content