FROM golang:1.14

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir /app

ADD . /app

WORKDIR  /app

RUN go mod download

RUN go build -o main .

EXPOSE 5000

CMD ["app/main"]



