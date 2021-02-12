# cybertect

[![Build Status](https://github.com/vicanso/cybertect/workflows/Test/badge.svg)](https://github.com/vicanso/cybertect/actions)

提供常用的HTTP接口、TCP端口、DNS域名解析以及Ping的定时检测告警。

## HTTP检测

HTTP检测通过指定检测URL，定时调用判断返回的HTTP状态码是否>=200且<400，如果是则认为成功，否则失败（对于https还检测期证书是否差不多过期，如果要过期则认为检测失败），失败时通过email发送告警邮箱。

- `名称` 检测配置名称
- `URL` 检测地址，配置检测的http(s)访问地址则可
- `IP列表` 指定URL中域名对应的解析，如果域名解析的IP为多个，可以配置多个IP地址，以`,`分隔。如果不需要指定（配置的检测地址为IP形式或直接通过DNS解析），则配置为`0.0.0.0`
- `状态` 是否启用状态
- `超时` 设置超时时长，单位为秒
- `接收者` 选择接收告警邮件的用户
- `描述` 检测配置描述

![](./images/http-setting.jpg)

完成配置之后，系统会定时执行检测配置，相关检测结果可在列表中查询并可查询每次检测的详情，包括HTTP请求的完成链路时间（tcp连接、tls连接等）。

![](./images/http-detect-result.jpg)

![](./images/http-detect-result-detail.jpg)

## DNS检测

DNS检测域名在指定DNS服务器的解析记录是否与期望的IP列表一致，主要用于检测是否有DNS劫持，支持IPV4与IPV6的DNS解析。

![](./images/dns-setting.jpg)

![](./images/dns-detect-result.jpg)

![](./images/dns-detect-result-detail.jpg)

## TCP

TCP检测指定的一堆地址的端口监听状态，如redis集群等，主要用于简单的服务是否可用的检测。

![](./images/tcp-setting.jpg)

![](./images/tcp-detect-result.jpg)

![](./images/tcp-detect-result-detail.jpg)

## Ping

Ping检测用于检测网络的连通性，主要用于测试简单的网络连通、机器是否在线等最基本的检测。

![](./images/ping-setting.jpg)

![](./images/ping-detect-result.jpg)

![](./images/ping-detect-result-detail.jpg)

### postgres

```
docker pull postgres:13-alpine

docker run -d --restart=always \
  -v $PWD/data:/var/lib/postgresql/data \
  -e POSTGRES_PASSWORD=A123456 \
  -p 5432:5432 \
  --name=cybertect-data \
  postgres:13-alpine

docker exec -it cybertect-data sh

psql -c "CREATE DATABASE cybertect;" -U postgres
psql -c "CREATE USER vicanso WITH PASSWORD 'A123456';" -U postgres
psql -c "GRANT ALL PRIVILEGES ON DATABASE cybertect to vicanso;" -U postgres
```

### 启动程序

```bash
docker run -d --restart=always \
  -p 7630:7001 \
  -e GO_ENV=production \
  -e POSTGRES_URI=postgresql://vicanso:A123456@127.0.0.1:5432/cybertect \
  -e SECRET=xxxx \
  -e MAIL_HOST=smtp.office365.com \
  -e MAIL_PORT=587 \
  -e MAIL_USER=tree.xie@outlook.com \
  -e MAIL_PASS=pass \
  -e DETECTOR_INTERVAL=1m \
  --name=cybertect \
  vicanso/cybertect
```

- `GO_ENV` 设置为正式环境
- `POSTGRES_URI` 数据库连接地址
- `SECRET` jwt的加密串，如果不指定每次启动时动态生成（每次重启程序则登录失效）
- `MAIL_HOST` 告警发送邮箱域名
- `MAIL_PORT` SMTP端口
- `MAIL_USER` 邮箱账号
- `MAIL_PASS` 邮箱密码
- `DETECTOR_INTERVAL` 检测间隔，默认为1m（1分钟一次)

## 规范

- 所有自定义的error都必须为hes.Error
