.PHONY: gen/code
gen/code:
	buf generate

.PHONY: update/dep
update/dep:
	go mod tidy
