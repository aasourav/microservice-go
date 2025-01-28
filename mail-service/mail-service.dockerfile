FROM golang:1.23-alpine AS builder
RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o mailApp ./cmd/api
RUN chmod +x /app/mailApp

FROM alpine
RUN mkdir /app
COPY --from=builder /app/mailApp /app
COPY --from=builder /app/templates /app/templates

CMD [ "/app/mailApp" ]
