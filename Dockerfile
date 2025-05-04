FROM golang:latest


WORKDIR /app

COPY . .

RUN go mod init learnDocker || true
RUN go mod tidy

RUN go build -o main queueArray.go
CMD ["./main"]
