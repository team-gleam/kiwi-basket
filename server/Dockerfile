FROM golang:1.14.3

WORKDIR /kiwi-basket
ADD ./server/src/ .
ADD ./config.yaml .

RUN go build -o kiwi-basket .
CMD ["./kiwi-basket"]
