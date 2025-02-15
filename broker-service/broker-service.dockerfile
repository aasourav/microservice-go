FROM golang:1.23-alpine AS builder
RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api
RUN chmod +x /app/brokerApp

FROM alpine
RUN mkdir /app
COPY --from=builder /app/brokerApp /app
CMD [ "/app/brokerApp" ]