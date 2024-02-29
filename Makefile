.PHONY: dev gen run vite build test cover fmt docker-build docker-run

.DEFAULT: help
help:
	@echo "make dev"
	@echo "	setup development environment"
	@echo "make gen"
	@echo "	run templ generate"
	@echo "make run"
	@echo "	build and run go app"
	@echo "make vite"
	@echo "	run vite"
	@echo "make build"
	@echo "	run go build and npm run build"
	@echo "make test"
	@echo "	run go test"
	@echo "make cover"
	@echo "	run go test with -cover"
	@echo "make cover-html"
	@echo "	run go test with -cover and show HTML"
	@echo "make tidy"
	@echo "	run go mod tidy"
	@echo "make fmt"
	@echo "	run gofumpt"
	@echo "make docker-build"
	@echo "	build docker image"
	@echo "make docker-run"
	@echo "	run docker image"

check-air:
ifeq (, $(shell which air))
	$(error "air not in $(PATH), air (https://github.com/cosmtrek/air) is required")
endif

check-gofumpt:
ifeq (, $(shell which gofumpt))
	$(error "gofumpt not in $(PATH), gofumpt (https://pkg.go.dev/mvdan.cc/gofumpt) is required")
endif

check-templ:
ifeq (, $(shell which templ))
	$(error "templ not in $(PATH), templ (https://templ.guide/quick-start/installation) is required")
endif

dev:
	npm i

gen: check-templ
	templ generate

run: check-air
	air

vite:
	npm run dev

build: gen
	go build -o app-bin ./cmd/app && npm run build

test: gen
	go test -v ./...

cover: gen
	go test -cover -v ./...

cover-html: gen
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

tidy:
	go mod tidy

fmt: check-gofumpt
	gofumpt -l -w .

docker-build: build
	docker build -t app .

docker-run:
	docker run -p 3000:3000 --rm app --env-name prod
