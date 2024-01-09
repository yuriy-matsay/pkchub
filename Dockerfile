FROM golang:1.18 AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build -o main .

FROM alpine:latest
COPY --from=builder /app ./

EXPOSE 8000
ENTRYPOINT ["./main"]
