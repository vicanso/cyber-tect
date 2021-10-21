# Cyber Tect

[![Build Status](https://github.com/vicanso/cyber-tect/workflows/Test/badge.svg)](https://github.com/vicanso/cyber-tect/actions)

提供常用的HTTP接口、TCP端口、DNS域名解析、Ping以及各常用数据库的定时检测告警。

## database

数据库暂仅支持postgres与mysql，仅需要在配置数据库连接串时指定则可，连接串格式如下：

- `postgres`: postgres://root:pass@127.0.0.1:5432/cybertect?maxIdleConns=5&maxIdleTime=30m&maxOpenConns=100
- `mysql`: mysql://root:pass@tcp(127.0.0.1:3306)/cybertect?timeout=30s&parseTime=true&maxIdleConns=5&maxIdleTime=30m&maxOpenConns=100

### postgres

用户信息及检测配置、结果等数据保存在postgres中，若无现成的postgres可使用以下脚本启动实例：

```
docker pull postgres:14-alpine

docker run -d --restart=always \
  -v $PWD/data:/var/lib/postgresql/data \
  -e POSTGRES_PASSWORD=A123456 \
  -p 5432:5432 \
  --name=postgres \
  postgres:14-alpine

docker exec -it postgres sh

psql -c "CREATE DATABASE cybertect;" -U postgres
psql -c "CREATE USER vicanso WITH PASSWORD 'A123456';" -U postgres
psql -c "GRANT ALL PRIVILEGES ON DATABASE cybertect to vicanso;" -U postgres
```

## 项目启动

项目连接数据库使用ent框架，相关代码动态生成，因此使用前需要先执行：

```bash
make install && make generate
```

启动程序：

```bash
go run main.go 
```

## HTTP检测

HTTP检测通过指定检测URL，定时调用判断返回的HTTP状态码是否>=200且<400，如果是则认为成功，否则失败（对于https还检测期证书是否差不多过期，如果要过期则认为检测失败），失败时通过email发送告警邮箱。配置如下：

- `名称` 检测配置名称
- `状态` 是否启用状态
- `超时` 设置超时时长，单位为秒
- `用户列表` 可以编辑该配置的用户
- `告警接收` 选择接收告警邮件的用户
- `IP列表` 指定URL中域名对应的解析，如果域名解析的IP为多个，可以配置多个IP地址，以`,`分隔。如果不需要指定（配置的检测地址为IP形式或直接通过DNS解析），则配置为`0.0.0.0`
- `检测地址` 配置检测的http(s)访问地址则可
- `配置描述` 检测配置描述

![](./images/http-setting.jpg)

完成配置之后，系统会定时执行检测配置，相关检测结果可在列表中查询并可查询每次检测的详情，包括HTTP(s)请求完整链路的时间（tcp连接、tls连接等）。

![](./images/http-detect-result.jpg)

![](./images/http-detect-result-detail.jpg)

## DNS检测

DNS检测域名在指定DNS服务器的解析记录是否与期望的IP列表一致，主要用于检测是否有DNS劫持，支持IPV4与IPV6的DNS解析。配置如下：

- `名称` 检测配置名称
- `状态` 是否启用状态
- `超时` 设置超时时长，单位为秒
- `用户列表` 可以编辑该配置的用户
- `告警接收` 选择接收告警邮件的用户
- `域名` 检测域名
- `IP列表` 域名对应的IP地址列表，如果DNS解析的IP不在此列表中，则失败
- `DNS` DNS服务器列表
- `配置描述` 检测配置描述

![](./images/dns-setting.jpg)

![](./images/dns-detect-result.jpg)

![](./images/dns-detect-result-detail.jpg)

## TCP检测

TCP检测指定的多个地址的端口监听状态(相关服务)，如redis集群等，主要用于简单的服务是否可用的检测。配置如下：

- `名称` 检测配置名称
- `状态` 是否启用状态
- `超时` 设置超时时长，单位为秒
- `用户列表` 可以编辑该配置的用户
- `告警接收` 选择接收告警邮件的用户
- `地址列表` 检测的地址列表
- `超时` 设置超时时长，单位为秒
- `配置描述` 检测配置描述

![](./images/tcp-setting.jpg)

![](./images/tcp-detect-result.jpg)

![](./images/tcp-detect-result-detail.jpg)

## Ping检测

Ping检测用于检测网络的连通性，主要用于测试简单的网络连通、机器是否在线等最基本的检测。配置如下：

- `名称` 检测配置名称
- `状态` 是否启用状态
- `超时` 设置超时时长，单位为秒
- `用户列表` 可以编辑该配置的用户
- `告警接收` 选择接收告警邮件的用户
- `IP列表` 检测的IP列表
- `配置描述` 检测配置描述


![](./images/ping-setting.jpg)

![](./images/ping-detect-result.jpg)

![](./images/ping-detect-result-detail.jpg)

## Database检测

Database检测用于测试数据库连通性，主要用于简单的测试数据库是否可正常连接，现支持以下数据库：`redis`，`postgres`，`mysql`以及`mongodb`

### Redis

Redis数据库支持三种模式，数据库驱动使用[go-redis](https://github.com/go-redis/redis)，数据库连接串格式如下：

- `单实例`: `redis://[:pass@]host:port/`，密码选项根据数据库是否有设置密码而添加。
- `Sentinel`: `redis://[:pass@]host1,port1,host2:port2/?master=master[&sentinelPassword=sentinelPassword]`，密码选项根据数据库是否有设置密码而添加，sentinel必须指定master，若不指定master则会判断为cluster模式。
- `Cluster`: `redis://[:pass@]host1,port1,host2:port2,host3:port3/`，密码选项根据数据库是否有设置密码而添加。

### Postgres

postgres连接串格式如下：`postgres://[jack:secret@]foo.example.com:5432[,...bar.example.com:5432]/mydb`，数据库驱动使用[pgx](https://github.com/jackc/pgx)模块。

### Mysql

mysql连接串格式如下：`mysql://[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]`，数据库驱动使用[mysql](https://github.com/go-sql-driver/mysql)模块。

### Mongodb

mongodb连接串格式如下：`mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[defaultauthdb][?options]]`，数据库驱动使用[mongodb](https://github.com/mongodb/mongo-go-driver)模块。

![](./images/database-setting.jpg)

![](./images/database-detect-result.jpg)

![](./images/database-detect-result-detail.jpg)

## 个人信息设置

告警信息使用Email发送，因此需要设置个人邮箱后才可接收到告警信息。

![](./images/profile.jpg)


## 启动程序

建议直接使用已打包好的docker镜像启动项目，启动脚本如下：

```bash
docker run -d --restart=always \
  -p 7001:7001 \
  -e GO_ENV=production \
  -e DATABASE_URI=postgres://vicanso:A123456@127.0.0.1:5432/cybertect \
  -e MAIL_SMTP=smtp://tree.xie@outlook.com:pass@smtp.office365.com:587 \
  -e DETECTOR_INTERVAL=1m \
  -e DETECTOR_RESULT_EXPIRED=30d \
  --name=cybertect \
  vicanso/cybertect
```

- `GO_ENV` 设置为正式环境
- `POSTGRES_URI` 数据库连接地址
- `MAIL_SMTP` 用于发送告警邮件的SMTP设置 
- `DETECTOR_INTERVAL` 检测间隔，默认为1m（1分钟一次)
- `DETECTOR_RESULT_EXPIRED` 检测结果过期时间，默认为30天(30d)
