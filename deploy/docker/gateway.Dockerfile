FROM golang:1.23
WORKDIR /app
COPY gateway/ .
RUN go build -o gateway ./cmd/main.go
EXPOSE 8086
CMD ["./gateway"]