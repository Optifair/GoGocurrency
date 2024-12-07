FROM golang:1.23
WORKDIR /currency
COPY currency/ .
RUN go build -o currency ./cmd/main.go
EXPOSE 8087
CMD ["./currency"]