FROM golang:1.14

WORKDIR /go/src/github.com/ughvj/adventurer

RUN go get -u github.com/PuerkitoBio/goquery && \
    go get -u google.golang.org/grpc