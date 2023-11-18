FROM golang:latest 

WORKDIR /app

RUN go env -w GO111MODULE=on
RUN go env -w GOPRIVATE=https://goproxy.cn,direct

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/go-nunu/nunu@latest 

