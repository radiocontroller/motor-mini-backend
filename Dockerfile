FROM golang:1.13 as mod
ENV GOPROXY="https://goproxy.cn"
WORKDIR /motor
COPY go.mod ./
RUN go mod download
COPY . /motor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

#FROM alpine:latest
#WORKDIR /motor
#COPY . /motor
