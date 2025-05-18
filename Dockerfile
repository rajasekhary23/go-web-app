FROM golang:1.19 AS builder
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .

RUN go build

FROM scratch
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
