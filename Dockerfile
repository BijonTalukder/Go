FROM golang:latest


WORKDIR /app

COPY . .

RUN go mod init learnDocker || true
RUN go mod tidy

RUN go build -o main CIRCULARRAY.go
CMD ["./main"]
# docker build -t learn-go-docker .
# docker run --rm learn-go-docker

