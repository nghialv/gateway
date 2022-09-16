.PHONY: build
build:
	go build cmd/gateway/gateway.go

.PHONY: gen
gen:
	go mod tidy
	buf mod update
	buf generate
