
## chassix简介

开发chassix的目的并非重复造轮子，而是简单的整合golang现有的框架，在工程实践中形成一致的风格、约定。

chassix基于社区优秀的开源框架整合、二次封装而来，同时支持导出整合框架的原有API，适合大多使用golang开发轻量级应用的场景，因此其定位为轻量级云原生应用开发框架。

## chassix组成

chassix-data 提供常用的sqlite、mysql、postgre、redis等数据库的访问简易封装。

chassix-restful 基于go-restful二次封装 提供更少的API、及灵活配置， 快速开发restful webservice。

chassix-micro 轻量级restful、gRPC 微服务框架。

chassix-logging chassix日志模块。

chassix-panel 提供docker、网关、配置等控制面板。

chassix-logging

``` shell
go get c5x.io/logx
```
chassix-bootstrap

约定配置通过yaml定义，支持从ctrip apollo获取配置，配置文件namespace也仅支持yaml文件。

``` shell
go get c5x.io/bootstrap
```

YAML文件配置
``` yaml
app:
  name: chassix-bootstrap-example
  version: 1.1.0
  env: dev
  debug: no
server:
  port: 8000
  addr: 0.0.0.0
logging:
  level: info  #info,debug,error
  report-caller: true
  no-colors: true
  caller-first: true
```
从apollo获取配置
```yaml
apollo:
  enable: true
  settings:
    appId: <your app>
    cluster: default
    namespaces:
      - app.yml
      - redis.yml
      - custom.yml
    ip: apollo-config-server:8080

```

