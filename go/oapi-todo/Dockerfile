FROM golang:1.22-alpine
RUN apk add --no-cache bash
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . /app

RUN go build -o main .

CMD ["./wait-for-it.sh", "db:3306", "--", "./main"]
