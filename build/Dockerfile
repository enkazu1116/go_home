# builder
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# ビルド（バイナリ名 main）
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/main ./cmd

# runtime
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/main /app/main
WORKDIR /app
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/app/main"]