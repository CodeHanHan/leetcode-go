COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

EMPTY :=
SPACE := $(EMPTY) $(EMPTY)

ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR) && pwd -P))
endif

USERS ?= yangchnet htchen

## content: 在每个用户文件夹下生成目录.
.PHONY: content
content: $(addprefix content.,$(USERS))

## content.<username>: 指定用户文件夹生成目录.
.PHONY: content.%
content.%:
	@go run content.go -d ./$* -o ./$*

## richtest: 做带输出测试.
.PHONY: richtest
richtest: $(addprefix richtest.,$(USERS))

## richtest.<username>: 指定用户目录做带输出测试.
.PHONY: richtest.%
richtest.%:
	@go build ./$*/... && go test -v -count=1 ./$*/...

## test: 做无输出测试.
.PHONY: test
test: $(addprefix test.,$(USERS))

## test.<username>: 指定用户目录做无输出测试.
.PHONY: test.%
test.%:
	@go test ./$*/... -count=1

## new.<username>.<dir>.<num>: 生成新的题目目录.
.PHONY: new.%
new.%:
	@$(eval list:=$(subst .,$(SPACE),$*))
	@$(eval username:=$(word 1,$(list)))
	@$(eval dir:=$(word 2,$(list)))
	@$(eval num:=$(word 3,$(list)))
	@cd ./$(username)/$(dir) && \
	 mkdir $(dir)$(num) && \
	 cd $(dir)$(num) && \
	 touch $(dir)$(num).go $(dir)$(num)_test.go && \
	 echo "package $(dir)$(num)" > $(dir)$(num).go && \
	 echo "package $(dir)$(num)" > $(dir)$(num)_test.go
	@python3 gen_readme.py $(username) $(dir) $(num)

## list.<username>.<dir>.<num> : 生成具有链表数据结构的题目目录.
.PHONY: list.%
list.%: new.%
	@sed -i \
	 's/package $(dir)$(num)/package $(dir)$(num)\n\nimport (\n\t. "github.com\/CodeHanHan\/leetcode-go\/base\/LinkList"\n)/g' \
	 ./$(username)/$(dir)/$(dir)$(num)/$(dir)$(num).go
	@sed -i \
	 's/"testing"/"testing"\n\n\t. "github.com\/CodeHanHan\/leetcode-go\/base\/LinkList"/g' \
	 ./$(username)/$(dir)/$(dir)$(num)/$(dir)$(num)_test.go

## tree.<username>.<dir>.<num>: 生成具有二叉树数据结构的题目目录.
.PHONY: tree.%
tree.%: new.%
	@sed -i \
	 's/package $(dir)$(num)/package $(dir)$(num)\n\nimport (\n\t. "github.com\/CodeHanHan\/leetcode-go\/base\/Tree"\n)/g' \
	 ./$(username)/$(dir)/$(dir)$(num)/$(dir)$(num).go
	@sed -i \
	 's/"testing"/"testing"\n\n\t. "github.com\/CodeHanHan\/leetcode-go\/base\/Tree"/g' \
	 ./$(username)/$(dir)/$(dir)$(num)/$(dir)$(num)_test.go

## help: Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"