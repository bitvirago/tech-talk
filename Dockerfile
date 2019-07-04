# Compile stage
FROM golang:1.11.5-alpine3.8 AS build-env
ENV CGO_ENABLED 0

# Allow Go to retreive the dependencies for the build step
RUN apk add --no-cache git

# Get Delve from a GOPATH not from a Go Modules project
WORKDIR /go/src/
RUN GO111MODULE=off go get github.com/go-delve/delve/cmd/dlv

ADD . /tech-talk/
WORKDIR /tech-talk/debuging/remote_debuging

RUN GO111MODULE=on go build -v -gcflags="-N -l" -o main main.go

EXPOSE 9999 2345

CMD ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "./main"]