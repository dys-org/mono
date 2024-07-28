FROM golang:1.22-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . /app

RUN go build -o main .

CMD ["./main"]
