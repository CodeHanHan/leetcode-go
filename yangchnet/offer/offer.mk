OFFER_PACKAGE:=offer
OFFER_USER := yangchnet

OFFER_DIR := $(ROOT_DIR)/$(OFFER_USER)/$(OFFER_PACKAGE)

.PHONY: $(OFFER_USER).$(OFFER_PACKAGE).new-%
$(OFFER_USER).$(OFFER_PACKAGE).new-%:
	@mkdir -p $(OFFER_DIR)/$(OFFER_PACKAGE)-$*
	@touch $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/readme.md \
	  $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/$(OFFER_PACKAGE)-$*.go \
	  $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/$(OFFER_PACKAGE)-$*_test.go
	@echo "package $(OFFER_PACKAGE)_$*" > $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/$(OFFER_PACKAGE)-$*.go
	@echo "package $(OFFER_PACKAGE)_$*" > $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/$(OFFER_PACKAGE)-$*_test.go

.PHONY: $(OFFER_USER).$(OFFER_PACKAGE).list-%
$(OFFER_USER).$(OFFER_PACKAGE).list-%: $(ROOT_DIR)/$(OFFER_USER)/base/LinkList/LinkList.go
	@mkdir -p $(OFFER_DIR)/$(OFFER_PACKAGE)-$*
	@touch $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/readme.md \
	  $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/$(OFFER_PACKAGE)-$*.go \
	  $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/$(OFFER_PACKAGE)-$*_test.go
	@echo "package $(OFFER_PACKAGE)_$*" > $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/$(OFFER_PACKAGE)-$*.go
	@echo "package $(OFFER_PACKAGE)_$*" > $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/$(OFFER_PACKAGE)-$*_test.go
	@cp $< $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/
	@sed -i 's/package list/package $(OFFER_PACKAGE)_$*/g' $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/LinkList.go

.PHONY: $(OFFER_USER).$(OFFER_PACKAGE).tree-%
$(OFFER_USER).$(OFFER_PACKAGE).tree-%: $(ROOT_DIR)/$(OFFER_USER)/base/Tree/Tree.go
	@mkdir -p $(OFFER_DIR)/$(OFFER_PACKAGE)-$*
	@touch $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/readme.md \
	  $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/$(OFFER_PACKAGE)-$*.go \
	  $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/$(OFFER_PACKAGE)-$*_test.go
	@echo "package $(OFFER_PACKAGE)_$*" > $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/$(OFFER_PACKAGE)-$*.go
	@echo "package $(OFFER_PACKAGE)_$*" > $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/$(OFFER_PACKAGE)-$*_test.go
	@cp $< $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/
	@sed -i 's/package tree/package $(OFFER_PACKAGE)_$*/g' $(OFFER_DIR)/$(OFFER_PACKAGE)-$*/Tree.go

.PHONY: $(OFFER_USER).$(OFFER_PACKAGE).test
$(OFFER_USER).$(OFFER_PACKAGE).test:
	@go build $(ROOT_DIR)/$(OFFER_USER)/$(OFFER_PACKAGE)/... && \
		go test -v -count=1 $(ROOT_DIR)/$(OFFER_USER)/$(OFFER_PACKAGE)/...
