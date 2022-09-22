FROM golang:1.19.1

WORKDIR /app

COPY math* ./

RUN go mod init mate2 && \
    go mod tidy && \
    go build -o /app/mate2

CMD ["./mate2"]
