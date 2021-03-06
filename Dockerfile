### Multi-stage build
FROM golang:1.10-alpine3.7 as build

RUN apk --no-cache add git ca-certificates

RUN go get -u -v github.com/goadesign/goa/... && \
    go get -u -v github.com/afex/hystrix-go/hystrix && \
    go get -u -v github.com/Microkubes/microservice-tools/... && \
    go get -u -v github.com/satori/go.uuid

COPY . /go/src/github.com/Microkubes/microservice-registration

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/Microkubes/microservice-registration


### Main
FROM scratch

ENV API_GATEWAY_URL="http://localhost:8001"

COPY --from=build /go/src/github.com/Microkubes/microservice-registration/config.json /config.json
COPY --from=build /go/bin/microservice-registration /usr/local/bin/microservice-registration
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8080

CMD ["/usr/local/bin/microservice-registration"]
