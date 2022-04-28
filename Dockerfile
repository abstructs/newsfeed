FROM golang:1.18.1-alpine3.15

ADD . /app
WORKDIR /app

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]