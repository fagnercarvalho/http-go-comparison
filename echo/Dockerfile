FROM golang:1.16.3-alpine

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["/go/bin/echo"]

#ENTRYPOINT ["tail"]
#CMD ["-f","/dev/null"]