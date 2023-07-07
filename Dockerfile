FROM golang:latest AS builder

WORKDIR /app

COPY . .

#CMD ["tail", "-f", "/dev/null"]
RUN GOOS=linux go build -o server .

ENTRYPOINT ["./server"]