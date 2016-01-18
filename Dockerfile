# Docker image for the Drone AWS Opsworks plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-opsworks
#     make deps build docker

FROM alpine:3.3

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-opsworks /bin/
ENTRYPOINT ["/bin/drone-opsworks"]
