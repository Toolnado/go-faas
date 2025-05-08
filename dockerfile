FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /app/faas ./cmd/faas

FROM alpine:3.19
COPY --from=builder /app/faas /faas
EXPOSE 8080
ENTRYPOINT ["/faas"]