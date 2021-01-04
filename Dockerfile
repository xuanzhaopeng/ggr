FROM alpine:3.12

ENV PASSWORD=test \
    DATA="<qa:browsers xmlns:qa=\"urn:config.gridrouter.qatools.ru\"></qa:browsers>"

RUN apk add -U tzdata ca-certificates apache2-utils curl && rm -Rf /var/cache/apk/*
COPY ggr /usr/bin
COPY scripts/entry.sh /etc/grid-router/entry.sh
RUN chmod +x /etc/grid-router/entry.sh
RUN mkdir -p /etc/grid-router/quota

EXPOSE 4444

ENTRYPOINT ["/etc/grid-router/entry.sh"]
