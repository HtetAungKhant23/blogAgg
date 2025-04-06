FROM golang:1.24-alpine AS builder

WORKDIR /build
COPY . .
RUN go build -o ./blogapi

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/blogapi ./blogapi
CMD ["/app/blogapi"]
