FROM golang:1.20.1

WORKDIR /usr/src/kubernetes-redeploy-workload

COPY go.mod go.sum ./
RUN go mod download

COPY src src

RUN CGO_ENABLED=0 go build -C src -a -o ../bin/kubernetes-redeploy-workload .

FROM alpine:3.17

RUN apk update && apk add bash

COPY --from=0 /usr/src/kubernetes-redeploy-workload/bin/kubernetes-redeploy-workload /usr/local/bin/kubernetes-redeploy-workload

COPY docker-entrypoint.sh /usr/local/bin/

CMD ["/usr/local/bin/docker-entrypoint.sh"]
