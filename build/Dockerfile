# go version
FROM golang:1.19.2-alpine3.16

# Update and install git
RUN apk upgrade --update && \
	apk --no-cache add git curl	vim	neovim

# make directory for app
RUN mkdir -p /go/src/github.com/go-todo

# set up working directory
WORKDIR /go/src/github.com/go-todo

# Migrate host files to container working directory
ADD . /go/src/github.com/go-todo

# import packages
RUN go get -u golang.org/x/tools/cmd/goimports
