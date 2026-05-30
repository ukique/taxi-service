FROM golang:1.25-alpine AS build
WORKDIR /server
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY . .
RUN go build -o service ./cmd/taxi-service

FROM alpine:latest
COPY --from=build /server/service .
COPY --from=build /go/bin/goose /usr/local/bin/goose
COPY migrations ./migrations
COPY ../../entrypoint.sh .
RUN chmod +x entrypoint.sh
CMD ["./entrypoint.sh"]