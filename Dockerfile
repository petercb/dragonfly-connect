FROM alpine:3.23

ARG TARGETPLATFORM

# hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates tzdata

COPY $TARGETPLATFORM/dragonfly-connect /
COPY servers.json /

ENTRYPOINT ["/dragonfly-connect"]
