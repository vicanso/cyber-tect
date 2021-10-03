# Cyber Tect

[![Build Status](https://github.com/vicanso/cyber-tect/workflows/Test/badge.svg)](https://github.com/vicanso/cyber-tect/actions)

提供常用的HTTP接口、TCP端口、DNS域名解析以及Ping的定时检测告警。

## 项目启动

项目连接数据库使用ent框架，相关代码动态生成，因此使用前需要先执行：

```bash
make install && make generate
```

## HTTP检测

HTTP检测通过指定检测URL，定时调用判断返回的HTTP状态码是否>=200且<400，如果是则认为成功，否则失败（对于https还检测期证书是否差不多过期，如果要过期则认为检测失败），失败时通过email发送告警邮箱。配置如下：

- `名称` 检测配置名称
- `URL` 检测地址，配置检测的http(s)访问地址则可
- `IP列表` 指定URL中域名对应的解析，如果域名解析的IP为多个，可以配置多个IP地址，以`,`分隔。如果不需要指定（配置的检测地址为IP形式或直接通过DNS解析），则配置为`0.0.0.0`
- `状态` 是否启用状态
- `超时` 设置超时时长，单位为秒
- `接收者` 选择接收告警邮件的用户
- `描述` 检测配置描述

![](./images/http-setting.jpg)

完成配置之后，系统会定时执行检测配置，相关检测结果可在列表中查询并可查询每次检测的详情，包括HTTP(s)请求完整链路的时间（tcp连接、tls连接等）。

![](./images/http-detect-result.jpg)

![](./images/http-detect-result-detail.jpg)

## DNS检测

DNS检测域名在指定DNS服务器的解析记录是否与期望的IP列表一致，主要用于检测是否有DNS劫持，支持IPV4与IPV6的DNS解析。配置如下：

- `名称` 检测配置名称
- `域名地址` 检测域名
- `DNS` DNS服务器，如果多个则以`,`分隔
- `状态` 是否启用状态
- `超时` 设置超时时长，单位为秒
- `接收者` 选择接收告警邮件的用户
- `IP列表` 域名对应的IP地址列表，多个IP以`,`分隔，如果DNS解析的IP不在此列表中，则失败
- `描述` 检测配置描述

![](./images/dns-setting.jpg)

![](./images/dns-detect-result.jpg)

![](./images/dns-detect-result-detail.jpg)

## TCP检测

TCP检测指定的多个地址的端口监听状态(相关服务)，如redis集群等，主要用于简单的服务是否可用的检测。配置如下：

- `名称` 检测配置名称
- `状态` 是否启用状态
- `检测地址` 检测的地址列表，如果多个以`,`分隔
- `超时` 设置超时时长，单位为秒
- `接收者` 选择接收告警邮件的用户
- `描述` 检测配置描述

![](./images/tcp-setting.jpg)

![](./images/tcp-detect-result.jpg)

![](./images/tcp-detect-result-detail.jpg)

## Ping检测

Ping检测用于检测网络的连通性，主要用于测试简单的网络连通、机器是否在线等最基本的检测。配置如下：

- `名称` 检测配置名称
- `状态` 是否启用状态
- `接收者` 选择接收告警邮件的用户
- `超时` 设置超时时长，单位为秒
- `检测IP` 检测的IP列表，如果多个以`,`分隔
- `描述` 检测配置描述


![](./images/ping-setting.jpg)

![](./images/ping-detect-result.jpg)

![](./images/ping-detect-result-detail.jpg)

## 个人信息设置

告警信息使用Email发送，因此需要设置个人邮箱后才可接收到告警信息。

![](./images/profile.jpg)

## 首页设置

首页展示最近的检测结果，可配置展示的数量及定时刷新。

![](./images/main-setting.jpg)

## postgres

用户信息及检测配置、结果等数据保存在postgres中，若无现成的postgres则可使用以下脚本启动实例：

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

## 启动程序

```bash
docker run -d --restart=always \
  -p 7001:7001 \
  -e GO_ENV=production \
  -e POSTGRES_URI=postgresql://vicanso:A123456@127.0.0.1:5432/cybertect \
  -e MAIL_HOST=smtp.office365.com \
  -e MAIL_PORT=587 \
  -e MAIL_USER=tree.xie@outlook.com \
  -e MAIL_PASS=pass \
  -e DETECTOR_INTERVAL=1m \
  -e DETECTOR_RESULT_EXPIRED=30d \
  --name=cybertect \
  vicanso/cybertect
```

- `GO_ENV` 设置为正式环境
- `POSTGRES_URI` 数据库连接地址
- `MAIL_HOST` 告警发送邮箱域名
- `MAIL_PORT` SMTP端口
- `MAIL_USER` 邮箱账号
- `MAIL_PASS` 邮箱密码
- `DETECTOR_INTERVAL` 检测间隔，默认为1m（1分钟一次)
- `DETECTOR_RESULT_EXPIRED` 检测结果过期时间，默认为30天(30d)
