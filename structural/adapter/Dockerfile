# 1 stage: buidler
FROM golang:alpine AS builder
WORKDIR /sample
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -ldflags="-s -w" -o adapter ./cmd/main.go

# 2 stage: release
FROM scratch
USER 1000:1000
COPY --from=builder /sample/adapter /adapter
ENTRYPOINT ["/adapter"]