FROM golang:1.23.3-alpine3.20 AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -trimpath -ldflags="-w -s" -o bin/receipt cmd/receipt/main.go


FROM alpine:3.20
ENV RECEIPT_DEBUG=${RECEIPT_DEBUG:-FALSE}
ENV RECEIPT_PORT=${RECEIPT_PORT:-":5001"}
WORKDIR /app
COPY --from=builder /app/bin/${NAME} .
LABEL maintainer="James Hilton <JamesHilton3310@gmail.com>"
CMD ["./receipt", "server"]