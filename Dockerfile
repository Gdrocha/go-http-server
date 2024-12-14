FROM golang:1.15 as build

WORKDIR /usr/src/app

COPY go.mod .

RUN go mod download

COPY . .

CMD ["go", "run", "src/main.go"]

EXPOSE 8081