FROM golang:alpine

MAINTAINER Jacques Vincilione <@jvincilione>

# Add git support
RUN apk add --update --no-cache git

# Copy the local package files to the container's workspace.
ADD . /go/src/lfs-portal

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN /go/src/lfs-portal/go-get.sh
RUN go install lfs-portal

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/lfs-portal

# Document that the service listens on port 80.
EXPOSE 8080
