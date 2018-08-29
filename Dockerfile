FROM golang:1.11.0

ENV ROOT_PATH /go/src/app
RUN mkdir $ROOT_PATH

WORKDIR $ROOT_PATH

RUN apt-get update && apt-get install -y mysql-client

# for dep
COPY . $ROOT_PATH

RUN go get -u github.com/golang/dep/cmd/dep && dep ensure
