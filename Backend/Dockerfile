# build stage of go binary
FROM golang:1.26.4-bookworm as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/server


FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/server /app

EXPOSE 80

CMD ["/app/server"]