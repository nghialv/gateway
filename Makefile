.PHONY: build
build:
	go build -o cmd/gateway/gateway cmd/gateway/main.go

.PHONY: test
test:
	go test ./...

.PHONY: gen
gen:
	rm -rf docs
	buf generate
	mv docs/pkg/* docs/
	rm -rf docs/pkg

.PHONY: dep
dep:
	go mod tidy
	buf mod update

.PHONY: lint
lint:
	buf lint

.PHONY: run/docs
run/docs:
	docker run -t -i --rm -p 1313:8080 \
    -e SWAGGER_JSON=/docs/service/greeting/v2/service.swagger.yaml \
    -v ${PWD}/docs/:/docs \
    swaggerapi/swagger-ui
