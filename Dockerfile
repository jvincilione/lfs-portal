FROM golang:alpine

MAINTAINER Jacques Vincilione <@jvincilione>

WORKDIR /go/src/lfs-portal

# Copy the local package files to the container's workspace.
COPY . /go/src/lfs-portal

# Add git support, install dependencies, build binary.
# Afterwards, remove all unneeded stuff to reduce layer size
RUN apk add --update --no-cache --virtual build-dependencies git \
  && /go/src/lfs-portal/go-get.sh \
  && apk del --no-cache build-dependencies \
  && go install lfs-portal \
  && rm -R /go/src/github.com /go/src/go.opencensus.io /go/src/google.golang.org /go/src/gopkg.in /go/src/cloud.google.com \
  && rm -R /go/pkg

# Run the lfs-portal command by default when the container starts.
ENTRYPOINT /go/bin/lfs-portal

# Document that the service listens on port 8080.
EXPOSE 8080
