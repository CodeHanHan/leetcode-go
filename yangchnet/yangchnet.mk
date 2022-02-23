MODULE ?= offer topic goden

.PHONY: yangchnet.test
yangchnet.test: $(addsuffix .test, $(addprefix yangchnet., $(MODULE)))

.PHONY: yangchnet.content
yangchnet.content:
	@$(GO) run content.go -d ./yangchnet -o ./yangchnet

