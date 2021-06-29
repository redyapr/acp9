# Stage 1 (Build Golang Project)
FROM golang:alpine AS builder
ENV GO111MODULE=on
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod ./
RUN go mod download
RUN go clean --modcache
COPY . .
RUN go build -o main

# Stage 2 (Reduce Size Without Golang Image)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
EXPOSE $APP_PORT
CMD ["./main"]