FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y --no-install-recommends postgresql-client && rm -rf /var/lib/apt/lists/*
RUN chmod +x wait-for-postgres.sh && mkdir -p files
RUN go build -o main ./cmd/avito

CMD ["./main"]
