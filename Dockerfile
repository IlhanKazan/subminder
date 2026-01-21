FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o subminder cmd/api/main.go


FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add tzdata

COPY --from=builder /app/subminder .

COPY --from=builder /app/configs ./configs
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./subminder"]