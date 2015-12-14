# Docker image for the Drone AWS Opsworks plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-aws-opsworks
#     make deps build docker

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-aws-opsworks /bin/
ENTRYPOINT ["/bin/drone-aws-opsworks"]
