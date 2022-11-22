FROM golang:1.18-alpine
RUN apk add --no-cache make
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./Makefile .
COPY ./planetocd.go .

COPY ./articles ./articles
COPY ./server/ ./server
COPY ./translate/ ./translate
COPY ./utils/ ./utils

# CGO_ENABLED=0 required when using scratch image
RUN CGO_ENABLED=0 make build

FROM scratch
WORKDIR /app
COPY --from=0 /app/bin/planetocd .
ENTRYPOINT ["/app/planetocd"]
