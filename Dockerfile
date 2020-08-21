FROM golang:1.14

WORKDIR /go/src/adventurer

RUN go get -u github.com/PuerkitoBio/goquery