FROM alpine:3.21

COPY ./openfeature-cli usr/local/bin/openfeature-cli

RUN chmod +x /usr/local/bin/openfeature-cli

ENTRYPOINT ["/usr/local/bin/openfeature-cli"]
