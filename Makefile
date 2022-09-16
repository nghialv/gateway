.PHONY: build
build:
	go build

.PHONY: gen/code
gen/code:
	buf generate

.PHONY: update/dep
update/dep:
	go mod tidy
	buf mod update
