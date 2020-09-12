FROM alpine:latest

WORKDIR /opt/cohesity-hello-world
COPY cohesity-hello-world /opt/cohesity-hello-world
COPY static /opt/cohesity-hello-world/static
CMD ["/opt/cohesity-hello-world/cohesity-hello-world"]
