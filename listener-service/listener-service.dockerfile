FROM golang:1.23-alpine AS builder
RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o loggerApp ./cmd/api
RUN chmod +x /app/loggerApp

FROM alpine
RUN mkdir /app
COPY --from=builder /app/loggerApp /app
CMD [ "/app/loggerApp" ]

# FROM alpine:latest

# RUN mkdir /app

# COPY listenerApp /app

# CMD [ "/app/listenerApp"]