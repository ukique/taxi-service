FROM golang:1.25-alpine AS build
WORKDIR /server
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o service ./cmd/taxi-simulation

FROM alpine:latest
COPY --from=build /server/service .
COPY --from=build /server/config ./config
CMD ["./service"]