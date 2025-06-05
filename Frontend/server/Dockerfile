FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build server.go

EXPOSE 8000

CMD [ "go", "run", "server.go", "--seed" ]