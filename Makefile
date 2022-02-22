ROOT_PACKAGE=github.com/Codehanhan/leetcode-go

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR) && pwd -P))
endif

USERS ?= yangchnet HT-CHEN520 MyOwnBoss9808

include yangchnet/yangchnet.mk
include yangchnet/offer/offer.mk
include yangchnet/topic/topic.mk
include yangchnet/goden/goden.mk

# TODO content 

.PHONY: test
test: 
	@$(MAKE) yangchnet.test
	@cd HT-CHEN520 && make test

.PHONY: cht-new-%
cht-new-%:
	@cd ./HT-CHEN520 && make new-$*
