FROM golang

MAINTAINER mini-alem

RUN mkdir /ascii-docker

ADD . /ascii-docker
WORKDIR /ascii-docker/server
RUN go build main.go

CMD ["/ascii-docker/server/main"]
