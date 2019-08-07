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
WORKDIR /
COPY --from=builder /usr/src/build .
CMD ["./gorest"]
