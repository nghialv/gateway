.PHONY: build
build:
	go build -o cmd/gateway/gateway cmd/gateway/main.go

.PHONY: gen
gen:
	go mod tidy
	buf mod update
	buf generate
