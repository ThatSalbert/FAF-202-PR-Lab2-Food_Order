FROM golang:alpine

RUN mkdir /kitchen

WORKDIR /kitchen

COPY . .

RUN go build -o /go/bin/main

EXPOSE 8000

ENTRYPOINT ["/go/bin/main"]

