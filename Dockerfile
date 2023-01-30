FROM alpine:3.17.1
ENTRYPOINT ["/meeting-summary"]
COPY meeting-summary /