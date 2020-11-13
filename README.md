<div align="center">

# gin-admin-template

 基于 GIN + GORM/MONGO + CASBIN + WIRE 实现的 RBAC 权限管理系统模块
 

[![License](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)

</div>

**Live demo:** [https://vue-iview-admin-template.ywbang.icu](https://vue-iview-admin-template.ywbang.icu)

**Swagger documentation:** [https://admin.ywbang.icu/swagger/index.html](https://admin.ywbang.icu/swagger/index.html)

## 特性

- 遵循 `RESTful API` 设计规范 & 基于接口的编程规范
- 基于 `GIN` 框架，提供了丰富的中间件支持（JWTAuth、CORS、RequestLogger、RequestRateLimiter、TraceID、CasbinEnforce、Recover、GZIP）
- 基于 `Casbin` 的 RBAC 访问控制模型 -- **权限控制可以细粒度到按钮 & 接口**
- 基于 `Gorm/Mongo` 的数据库存储 -- 存储层抽象了标准的外部业务层调用接口，内部采用封闭式实现（为后续切换数据存储提供了较大的便利）
- 基于 `WIRE` 的依赖注入 -- 依赖注入本身的作用是解决了各个模块间层级依赖繁琐的初始化过程
- 基于 `Logrus & Context` 实现了日志输出，通过结合 Context 实现了统一的 TraceID/UserID 等关键字段的输出(同时支持日志钩子写入到`Gorm/Mongo`)
- 基于 `JWT` 的用户认证 -- 基于 JWT 的黑名单验证机制
- 基于 `Swaggo` 自动生成 `Swagger` 文档 -- 独立于接口的 mock 实现
- 基于 `net/http/httptest` 标准包实现了 API 的单元测试
- 基于 `go mod` 的依赖管理

## 依赖工具

```bash
go get -u github.com/cosmtrek/air
go get -u github.com/google/wire/cmd/wire
go get -u github.com/swaggo/swag/cmd/swag
```

- [air](https://github.com/cosmtrek/air) -- Live reload for Go apps
- [wire](https://github.com/google/wire) -- Compile-time Dependency Injection for Go
- [swag](https://github.com/swaggo/swag) -- Automatically generate RESTful API documentation with Swagger 2.0 for Go.

## 依赖框架

- [Gin](https://gin-gonic.com/) -- The fastest full-featured web framework for Go.
- [GORM](http://gorm.io/) -- The fantastic ORM library for Golang
- [Mongo](https://github.com/mongodb/mongo-go-driver) -- The Go driver for MongoDB
- [Casbin](https://casbin.org/) -- An authorization library that supports access control models like ACL, RBAC, ABAC in Golang
- [Wire](https://github.com/google/wire) -- Compile-time Dependency Injection for Go

## 快速开始

```bash
# clone project
git clone -b master https://github.com/liuvigongzuoshi/gin-admin-template
# switch to the project directory
cd gin-admin-template
# use air to run
air
# or use Makefile to run
make start
# or use the go command to run
go run cmd/gin-admin/main.go web -c ./configs/config.toml
```

启动成功之后，可在浏览器中输入地址进行访问：`http://localhost:10088/swagger/index.html`

## 生成 `swagger` 文档

```bash
# 基于 Makefile
make swagger
# 或者使用 swag 命令
swag init --parseDependency --generalInfo ./internal/app/swagger.go --output ./internal/app/swagger
```

## 重新生成依赖注入文件

```bash
# 基于 Makefile
make wire
# 或者使用 wire 命令
wire gen ./internal/app
```

## 前端工程

[vue-iview-admin-template](https://github.com/liuvigongzuoshi/vue-iview-admin-template) - 基于 View UI 组件库参考 Ant Design Pro 的 vue 2.0 后台管理系统模板

## 工具

- [gin-admin-cli](https://github.com/gin-admin/gin-admin-cli) - [GinAdmin](https://github.com/LyricTian/gin-admin) 辅助工具，提供创建项目、快速生成功能模块的功能

## License

[MIT](https://github.com/liuvigongzuoshi/gin-admin-template/blob/master/LICENSE)

Copyright (c) 2020 liuvigongzuoshi
