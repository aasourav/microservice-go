FROM golang:1.23-alpine AS builder
RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o authApp ./cmd/api
RUN chmod +x /app/authApp

FROM alpine
RUN mkdir /app
COPY --from=builder /app/authApp /app
CMD [ "/app/authApp" ]