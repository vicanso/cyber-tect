FROM node:14-alpine as webbuilder

COPY . /cybertect
RUN cd /cybertect/web \
  && npm i \
  && npm run build \
  && rm -rf node_module

FROM golang:1.17-alpine as builder

COPY --from=webbuilder /cybertect /cybertect

RUN apk update \
  && apk add git make curl jq \
  && cd /cybertect \
  && rm -rf asset/dist \
  && cp -rf web/dist asset/ \
  && make install \
  && make generate \
  && ./download-swagger.sh \
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

HEALTHCHECK --timeout=10s --interval=10s CMD [ "wget", "http://127.0.0.1:7001/ping", "-q", "-O", "-"]

CMD ["cybertect"]

ENTRYPOINT ["/entrypoint.sh"]
