ROOT_PACKAGE=github.com/Codehanhan/leetcode-go

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

GO := go

ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR) && pwd -P))
endif

USERS ?= yangchnet HT-CHEN520

include yangchnet/yangchnet.mk
include HT-CHEN520/HT-CHEN520.mk
include yangchnet/offer/offer.mk
include yangchnet/topic/topic.mk
include yangchnet/goden/goden.mk

.PHONY: content
content: $(addsuffix .content,$(USERS))

.PHONY: test
test: $(addsuffix .test,$(USERS))

.PHONY: cht-new-%
cht-new-%:
	@cd ./HT-CHEN520 && make new-$*
