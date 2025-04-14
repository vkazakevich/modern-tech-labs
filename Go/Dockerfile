FROM golang:1.24

WORKDIR /home/app

RUN go install github.com/air-verse/air@latest

EXPOSE 1323

COPY . .

RUN go build server.go

CMD ["air", "-c", ".air.toml"]