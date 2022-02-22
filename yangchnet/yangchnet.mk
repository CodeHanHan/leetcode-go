MODULE ?= offer topic goden

.PHONY: yangchnet.test
yangchnet.test: $(addsuffix .test, $(addprefix yangchnet., $(MODULE)))

