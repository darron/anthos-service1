FROM cgr.dev/chainguard/alpine-base:latest
LABEL org.label-schema.vcs-url="https://github.com/darron/anthos-service1"
RUN apk add --update --no-cache \
  ca-certificates && \
  apk add curl && \
  rm -vf /var/cache/apk/*
WORKDIR /
COPY curl.sh /
CMD ["/bin/sh", "/curl.sh"]