FROM golang:1.23.3

WORKDIR /app

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/ .

RUN go build -o tinycache .

EXPOSE 8080

CMD ["./tinycache"]