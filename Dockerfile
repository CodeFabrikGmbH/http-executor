FROM golang:1.11.10-alpine3.9 as build

WORKDIR /workdir

RUN apk add --no-cache curl build-base git ca-certificates

COPY . /workdir

RUN go build -o /app

FROM alpine:3.9

RUN apk add --no-cache docker

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]