FROM golang:1.14

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8081

CMD ["go", "run", "main.go"]
