FROM golang:latest

WORKDIR /app

COPY . .

CMD ["tail", "-f", "/dev/null"]