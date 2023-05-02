FROM golang:1.20 as build
RUN mkdir /service1
WORKDIR /service1
ADD . .
RUN make linux
RUN chmod a+x /service1/bin/service1

FROM cgr.dev/chainguard/alpine-base:latest
LABEL org.label-schema.vcs-url="https://github.com/darron/anthos-service1"
RUN apk add --update --no-cache \
  ca-certificates && \
  rm -vf /var/cache/apk/*
WORKDIR /
COPY --from=build /service1/bin/service1 .
ENTRYPOINT ["/service1"]