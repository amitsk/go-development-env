# syntax=docker/dockerfile:1

FROM golang:1.26-alpine AS builder
WORKDIR /src
RUN apk add --no-cache ca-certificates
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/heroes ./cmd/heroes

FROM alpine:3.21
RUN adduser -D -H -u 65532 app
COPY --from=builder /out/heroes /heroes
USER app
EXPOSE 8080
ENTRYPOINT ["/heroes"]
