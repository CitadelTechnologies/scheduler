FROM golang:1.10

WORKDIR /go/src/scheduler

COPY . .

RUN go get -d -v ./... && go install -v ./...

EXPOSE 80

CMD ["/go/bin/scheduler"]
