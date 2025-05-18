FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
# Build statically for scratch
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o goapp

FROM scratch
COPY --from=builder /app/goapp /goapp
EXPOSE 8081
ENTRYPOINT ["/goapp"]
