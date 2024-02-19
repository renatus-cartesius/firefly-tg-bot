FROM golang:1.20-alpine as builder

WORKDIR /builder
COPY . .
RUN go build -o bin/

FROM alpine:3.19.1 as executor

WORKDIR /app
COPY --from=builder /builder/bin/firebot .

CMD ["/app/firebot"]