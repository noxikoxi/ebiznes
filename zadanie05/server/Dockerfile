FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o myapp .

EXPOSE 1323

CMD ["./myapp"]