FROM golang:1.18-buster AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080

FROM base AS builder
RUN go build -o api main.go

FROM alpine:3.15 as runner
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["api"]