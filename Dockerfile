FROM golang:1.4
MAINTAINER tobe tobeg3oolge@gmail.com

RUN apt-get update -y

RUN apt-get install -y git

ADD . /go/kubernetes-worker
WORKDIR /go/kubernetes-worker

RUN go get
RUN go build

CMD /bin/bash