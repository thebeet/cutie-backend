FROM golang:1.22-buster AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go env -w GOPROXY=https://proxy.golang.com.cn,direct
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /app/app .

FROM alpine:3.16
COPY --from=builder /app/app /app
ENV TZ=Asia/Shanghai
ENTRYPOINT ["/app"]