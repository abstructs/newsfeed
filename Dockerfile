FROM golang:1.18.1-alpine3.15

# WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
# COPY *.go ./

ADD . /app
WORKDIR /app

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

# RUN go build -o ./build

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]