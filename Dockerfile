FROM docker.io/golang:alpine AS builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o main .

FROM docker.io/alpine:latest
RUN mkdir /app
COPY --from=builder /app/main /app/main
WORKDIR /app

ENTRYPOINT [ "/app/main" ]
