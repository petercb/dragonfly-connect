FROM alpine:3.18
RUN apk add --no-cache ca-certificates tzdata
COPY dragonfly-connect /
COPY servers.json /
ENTRYPOINT ["/dragonfly-connect"]
