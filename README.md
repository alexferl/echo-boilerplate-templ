# echo-boilerplate-templ [![Go Report Card](https://goreportcard.com/badge/github.com/alexferl/echo-boilerplate-templ)](https://goreportcard.com/report/github.com/alexferl/echo-boilerplate-templ) [![codecov](https://codecov.io/gh/alexferl/echo-boilerplate-templ/branch/master/graph/badge.svg)](https://codecov.io/gh/alexferl/echo-boilerplate-templ)

A Go boilerplate for web development using:
- [echo](https://github.com/labstack/echo)
- [templ](https://templ.guide/)
- [htmx](https://htmx.org/)
- [Alpine.js](https://alpinejs.dev/)
- [Flowbite](https://flowbite.com/)
- [TailwindCSS](https://tailwindcss.com/)

## Requirements
Before getting started, install the following:

Required:
- [Air](https://github.com/cosmtrek/air?tab=readme-ov-file#installation)
- [Node.js](https://nodejs.org/en/download)
- [templ](https://templ.guide/quick-start/installation)

Optional:
- [gofumpt](https://pkg.go.dev/mvdan.cc/gofumpt) (needed to run `make fmt`)

## Using
Setup the dev environment first:
```shell
make dev
```

### Building & Running locally
```shell
make run
```
In another terminal:
```shell
make vite
```

Navigate to `http://localhost:3000` in your browser.

### Docker
#### Build
```shell
make docker-build
```

#### Run
```shell
make docker-run
```

#### Passing args
CLI:
```shell
docker run -p 3000:3000 --rm app --env-name prod
```

Environment variables:
```shell
docker run -p 3000:3000 -e "APP_ENV_NAME=prod" --rm app
```
