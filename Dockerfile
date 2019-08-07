FROM golang:1.12 AS builder

WORKDIR /usr/src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make compile




FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache libc6-compat
COPY --from=builder /usr/src/build .
COPY --from=builder /usr/src/config.json.prod config.json
CMD ["./gorest"]
