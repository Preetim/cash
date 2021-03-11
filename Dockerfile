FROM golang:1.14-alpine3.11 as baseimage

RUN mkdir /app

FROM baseimage

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

RUN go test

CMD []
