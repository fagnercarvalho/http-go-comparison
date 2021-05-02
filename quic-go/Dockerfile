FROM golang:1.16.3-buster

ENV GO111MODULE=on
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["/go/bin/quic-go"]