FROM golang:1.24

WORKDIR /home/app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go build server.go

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]