FROM golang:1.10 as application
WORKDIR /go/src/scheduler
COPY . .
RUN go get -d -v ./... && go install -v ./...

FROM node:carbon as documentation
WORKDIR /root/
COPY docker-entrypoint.sh /entrypoint.sh
COPY --from=application /go/bin/scheduler .
COPY --from=application /go/src/scheduler/doc ./doc
RUN chmod a+x /entrypoint.sh && chown root:root /entrypoint.sh && npm install -g api-console-cli

EXPOSE 80
EXPOSE 8000

ENTRYPOINT ["/entrypoint.sh"]

CMD ["./scheduler"]
