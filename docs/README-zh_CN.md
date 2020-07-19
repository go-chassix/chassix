# chassis-go 
[![Build Status](https://cloud.drone.io/api/badges/chassisx/chassis-go/status.svg)](https://cloud.drone.io/chassisx/chassis-go)

chassis-go 集成优秀的go框架以使用golang快速开发web服务

使用[go-restful](https://github.com/emicklei/go-restful)开发rest api
使用[gorm](https://gorm.io)操作关系数据库
支持从yaml文件、apollo配置中心读取配置
开箱即用的内存、Redis缓存
使用```github.com/go-playground/validator```验证struct属性

Features:

- [x] gorm
- [x] go-restful
- [x] YAML config files
- [x] Validate
- [x] DB migrate
- [x] API errors
- [x] Redis cache
- [x] Memory cache

usage:
```
go get c6x.io/chassis
```
