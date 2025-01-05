FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env.docker .env

RUN go build -o myserv .

EXPOSE 8181

CMD ["./myserv"]