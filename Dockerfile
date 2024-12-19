FROM golang:1.23.3-alpine as buildbase

ARG CI_JOB_TOKEN

RUN apk add git build-base ca-certificates

WORKDIR /go/src/github.com/rarimo/bio-data-svc
COPY . .

RUN go mod tidy && go mod vendor
RUN GOOS=linux go build -o /usr/local/bin/bio-data-svc /go/src/github.com/rarimo/bio-data-svc

FROM scratch
COPY --from=alpine:3.9 /bin/sh /bin/sh
COPY --from=alpine:3.9 /usr /usr
COPY --from=alpine:3.9 /lib /lib

COPY --from=buildbase /usr/local/bin/bio-data-svc /usr/local/bin/bio-data-svc
COPY --from=buildbase /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


ENTRYPOINT ["bio-data-svc"]
