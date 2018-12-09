FROM golang:alpine AS builder
MAINTAINER Stephane Busso <stephane.busso@gmail.com>
RUN apk update && apk add git && apk add ca-certificates
RUN adduser -D -g '' appuser
WORKDIR /src
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o server

FROM alpine AS production
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /src/server /usr/local/bin/server
USER appuser
CMD server