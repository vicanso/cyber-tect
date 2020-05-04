FROM node:12-alpine as webbuilder

COPY . /cybertect
RUN cd /cybertect/web \
  && yarn \
  && yarn build \
  && rm -rf node_module

FROM golang:1.14-alpine as builder

COPY --from=webbuilder /cybertect /cybertect

RUN apk update \
  && apk add git make \
  && go get -u github.com/gobuffalo/packr/v2/packr2 \
  && cd /cybertect \
  && make build

FROM alpine 

EXPOSE 7001

# tzdata 安装所有时区配置或可根据需要只添加所需时区

RUN addgroup -g 1000 go \
  && adduser -u 1000 -G go -s /bin/sh -D go \
  && apk add --no-cache ca-certificates tzdata

COPY --from=builder /cybertect/cybertect /usr/local/bin/cybertect
COPY --from=builder /cybertect/entrypoint.sh /entrypoint.sh

USER go

WORKDIR /home/go

HEALTHCHECK --timeout=10s CMD [ "wget", "http://127.0.0.1:7001/ping", "-q", "-O", "-"]

CMD ["cybertect"]

ENTRYPOINT ["/entrypoint.sh"]
