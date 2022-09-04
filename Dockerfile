FROM golang:1.18-bullseye AS builder
WORKDIR /app
COPY . .
RUN go mod download && go mod tidy && go mod verify
RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 go build

FROM gcr.io/distroless/base-debian10 AS Go-Server
COPY --from=builder /app/go-books-api ./server
EXPOSE 8000
ENTRYPOINT ["./server"]
