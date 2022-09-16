.PHONY: build
build:
	go build -o cmd/gateway/gateway cmd/gateway/main.go

.PHONY: test
test:
	go test ./...

.PHONY: gen
gen:
	go mod tidy
	buf mod update
	buf generate
	mv docs/pkg/* docs/
	rm -rf docs/pkg

.PHONY: lint
lint:
	buf lint
