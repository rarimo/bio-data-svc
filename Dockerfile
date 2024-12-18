FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/rarimo/bio-data-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/bio-data-svc /go/src/github.com/rarimo/bio-data-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/bio-data-svc /usr/local/bin/bio-data-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["bio-data-svc"]
