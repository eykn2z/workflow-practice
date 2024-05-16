FROM golang:alpine as builder
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR $GOPATH/bin/
COPY . ./api-server

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/api-server /go/bin/api-server

ENTRYPOINT ["/go/bin/api-server"]