FROM golang:1.17.11

WORKDIR /app

COPY *.go ./

RUN go mod init mate && \
    go mod tidy && \
    go build -o /app/mate

CMD ["./mate"]
