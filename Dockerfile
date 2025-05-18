# Builder stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o app

# Final stage
FROM scratch
COPY --from=builder /app/app /app
EXPOSE 8081
ENTRYPOINT ["/app"]
