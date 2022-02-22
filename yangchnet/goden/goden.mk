GODEN_PACKAGE:=goden
GODEN_USER := yangchnet

GODEN_DIR := $(ROOT_DIR)/$(GODEN_USER)/$(GODEN_PACKAGE)

.PHONY: $(GODEN_USER).$(GODEN_PACKAGE).new-%
$(GODEN_USER).$(GODEN_PACKAGE).new-%:
	@mkdir -p $(GODEN_DIR)/$(GODEN_PACKAGE)-$*
	@touch $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/readme.md \
	  $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/$(GODEN_PACKAGE)-$*.go \
	  $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/$(GODEN_PACKAGE)-$*_test.go
	@echo "package $(GODEN_PACKAGE)_$*" > $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/$(GODEN_PACKAGE)-$*.go
	@echo "package $(GODEN_PACKAGE)_$*" > $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/$(GODEN_PACKAGE)-$*_test.go

.PHONY: $(GODEN_USER).$(GODEN_PACKAGE).list-%
$(GODEN_USER).$(GODEN_PACKAGE).list-%: $(ROOT_DIR)/$(GODEN_USER)/base/LinkList/LinkList.go
	@mkdir -p $(GODEN_DIR)/$(GODEN_PACKAGE)-$*
	@touch $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/readme.md \
	  $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/$(GODEN_PACKAGE)-$*.go \
	  $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/$(GODEN_PACKAGE)-$*_test.go
	@echo "package $(GODEN_PACKAGE)_$*" > $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/$(GODEN_PACKAGE)-$*.go
	@echo "package $(GODEN_PACKAGE)_$*" > $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/$(GODEN_PACKAGE)-$*_test.go
	@cp $< $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/
	@sed -i 's/package list/package $(GODEN_PACKAGE)_$*/g' $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/LinkList.go

.PHONY: $(GODEN_USER).$(GODEN_PACKAGE).tree-%
$(GODEN_USER).$(GODEN_PACKAGE).tree-%: $(ROOT_DIR)/$(GODEN_USER)/base/Tree/Tree.go
	@mkdir -p $(GODEN_DIR)/$(GODEN_PACKAGE)-$*
	@touch $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/readme.md \
	  $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/$(GODEN_PACKAGE)-$*.go \
	  $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/$(GODEN_PACKAGE)-$*_test.go
	@echo "package $(GODEN_PACKAGE)_$*" > $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/$(GODEN_PACKAGE)-$*.go
	@echo "package $(GODEN_PACKAGE)_$*" > $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/$(GODEN_PACKAGE)-$*_test.go
	@cp $< $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/
	@sed -i 's/package tree/package $(GODEN_PACKAGE)_$*/g' $(GODEN_DIR)/$(GODEN_PACKAGE)-$*/Tree.go

.PHONY: $(GODEN_USER).$(GODEN_PACKAGE).test
$(GODEN_USER).$(GODEN_PACKAGE).test:
	@go build $(ROOT_DIR)/$(GODEN_USER)/$(GODEN_PACKAGE)/... && \
		go test -v -count=1 $(ROOT_DIR)/$(GODEN_USER)/$(GODEN_PACKAGE)/...
