FROM golang:alpine

RUN mkdir /food-order

WORKDIR /food-order

COPY . .

RUN go build -o /go/bin/main

EXPOSE 7000

ENTRYPOINT ["/go/bin/main"]

