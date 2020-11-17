# STEP 1 
FROM golang:alpine as builder
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR $GOPATH/bin/
COPY ./exe/client-service_linux ./client-service

# STEP 2 build a small image
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/client-service /go/bin/client-service

ENV API_SERVICE_URL 172.17.0.2
ENV API_SERVICE_PORT 50051

ENTRYPOINT ["/go/bin/client-service"]