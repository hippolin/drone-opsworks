# Docker image for the Drone Opsworks plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-opsworks
#     make deps build docker

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-opsworks /bin/
ENTRYPOINT ["/bin/drone-opsworks"]
