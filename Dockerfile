FROM golang:1.10

WORKDIR /go/src/github.com/efrenfuentes/imageproxy
COPY . .

RUN go get -d -v ./...
RUN go build

EXPOSE 4567
CMD ["./imageproxy"]

