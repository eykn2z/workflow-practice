FROM golang:alpine as builder
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
COPY . /app/api-server
RUN GOOS=linux GOARCH=amd64 go build -o api-server main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/api-server/api-server /go/bin/api-server

ENTRYPOINT ["/go/bin/api-server"]