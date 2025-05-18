# Builder stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod ./app
RUN go mod download
COPY . .
RUN go build -o goapp

# Final stage
FROM scratch
COPY --from=builder /app/goapp /goapp
EXPOSE 8081
ENTRYPOINT ["/goapp"]
