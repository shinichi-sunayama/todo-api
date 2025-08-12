# ====== ビルドステージ ======
FROM golang:1.22.4-bullseye AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
COPY .env .env

RUN go build -o main .

# ====== 実行ステージ ======
FROM golang:1.22.4-bullseye

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

EXPOSE 8080

CMD ["./main"]
