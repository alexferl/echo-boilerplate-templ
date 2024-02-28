ARG GOLANG_VERSION=1.21-alpine
FROM golang:${GOLANG_VERSION} AS builder
MAINTAINER Alexandre Ferland <me@alexferl.com>

WORKDIR /build

RUN apk add --no-cache git

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" ./cmd/app

FROM scratch
COPY --from=builder /build/app /app
COPY --from=builder /build/configs /configs
COPY --from=builder /build/static/dist /static/dist

ENTRYPOINT ["/app"]

EXPOSE 1323