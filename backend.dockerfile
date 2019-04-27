FROM golang:latest
RUN go get -u github.com/golang/dep/cmd/dep
RUN mkdir -p /go/src/github.com/hiroaki-yamamoto/go-gql-sample
WORKDIR /go/src/github.com/hiroaki-yamamoto/go-gql-sample/backend
