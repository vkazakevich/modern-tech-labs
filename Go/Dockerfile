FROM golang:1.24

WORKDIR /home/app

RUN go install github.com/air-verse/air@latest

EXPOSE 1323

CMD ["air", "-c", ".air.toml"]