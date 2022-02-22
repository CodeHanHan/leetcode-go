TOPIC_PACKAGE:=topic
TOPIC_USER := yangchnet

TOPIC_DIR := $(ROOT_DIR)/$(TOPIC_USER)/$(TOPIC_PACKAGE)

.PHONY: $(TOPIC_USER).$(TOPIC_PACKAGE).new-%
$(TOPIC_USER).$(TOPIC_PACKAGE).new-%:
	@mkdir -p $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*
	@touch $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/readme.md \
	  $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/$(TOPIC_PACKAGE)-$*.go \
	  $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/$(TOPIC_PACKAGE)-$*_test.go
	@echo "package $(TOPIC_PACKAGE)_$*" > $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/$(TOPIC_PACKAGE)-$*.go
	@echo "package $(TOPIC_PACKAGE)_$*" > $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/$(TOPIC_PACKAGE)-$*_test.go

.PHONY: $(TOPIC_USER).$(TOPIC_PACKAGE).list-%
$(TOPIC_USER).$(TOPIC_PACKAGE).list-%: $(ROOT_DIR)/$(TOPIC_USER)/base/LinkList/LinkList.go
	@mkdir -p $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*
	@touch $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/readme.md \
	  $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/$(TOPIC_PACKAGE)-$*.go \
	  $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/$(TOPIC_PACKAGE)-$*_test.go
	@echo "package $(TOPIC_PACKAGE)_$*" > $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/$(TOPIC_PACKAGE)-$*.go
	@echo "package $(TOPIC_PACKAGE)_$*" > $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/$(TOPIC_PACKAGE)-$*_test.go
	@cp $< $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/
	@sed -i 's/package list/package $(TOPIC_PACKAGE)_$*/g' $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/LinkList.go

.PHONY: $(TOPIC_USER).$(TOPIC_PACKAGE).tree-%
$(TOPIC_USER).$(TOPIC_PACKAGE).tree-%: $(ROOT_DIR)/$(TOPIC_USER)/base/Tree/Tree.go
	@mkdir -p $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*
	@touch $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/readme.md \
	  $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/$(TOPIC_PACKAGE)-$*.go \
	  $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/$(TOPIC_PACKAGE)-$*_test.go
	@echo "package $(TOPIC_PACKAGE)_$*" > $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/$(TOPIC_PACKAGE)-$*.go
	@echo "package $(TOPIC_PACKAGE)_$*" > $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/$(TOPIC_PACKAGE)-$*_test.go
	@cp $< $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/
	@sed -i 's/package tree/package $(TOPIC_PACKAGE)_$*/g' $(TOPIC_DIR)/$(TOPIC_PACKAGE)-$*/Tree.go

.PHONY: $(TOPIC_USER).$(TOPIC_PACKAGE).test
$(TOPIC_USER).$(TOPIC_PACKAGE).test:
	@go build $(ROOT_DIR)/$(TOPIC_USER)/$(TOPIC_PACKAGE)/... && \
		go test -v -count=1 $(ROOT_DIR)/$(TOPIC_USER)/$(TOPIC_PACKAGE)/...
