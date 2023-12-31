<p align="center"><a href="https://vuejs.org" target="_blank" rel="noopener noreferrer"><img width="400" src="https://youxiatest.oss-cn-beijing.aliyuncs.com/fungo_logo.png" alt="Vue logo"></a></p>

<p align="center">
  <a href="https://circleci.com/gh/vuejs/vue/tree/dev"><img src="https://img.shields.io/circleci/project/github/vuejs/vue/dev.svg?sanitize=true" alt="Build Status"></a>
  <a href="https://codecov.io/github/vuejs/vue?branch=dev"><img src="https://img.shields.io/codecov/c/github/vuejs/vue/dev.svg?sanitize=true" alt="Coverage Status"></a>
  <a href="https://npmcharts.com/compare/vue?minimal=true"><img src="https://img.shields.io/npm/dm/vue.svg?sanitize=true" alt="Downloads"></a>
  <a href="https://www.npmjs.com/package/vue"><img src="https://img.shields.io/npm/v/vue.svg?sanitize=true" alt="Version"></a>
  <a href="https://www.npmjs.com/package/vue"><img src="https://img.shields.io/npm/l/vue.svg?sanitize=true" alt="License"></a>
  <a href="https://chat.vuejs.org/"><img src="https://img.shields.io/badge/chat-on%20discord-7289da.svg?sanitize=true" alt="Chat"></a>
</p>

## 关于

`go-gin-api` 是基于 [Gin](https://github.com/gin-gonic/gin) 进行模块化设计的 API 框架，封装了常用的功能，使用简单，致力于进行快速的业务研发，同时增加了更多限制，约束项目组开发成员，规避混乱无序及自由随意的编码。

供参考学习，线上使用请谨慎！

集成组件：

1. 支持 [rate](https://golang.org/x/time/rate) 接口限流 
1. 支持 panic 异常时邮件通知 
1. 支持 [cors](https://github.com/rs/cors) 接口跨域 
1. 支持 [Prometheus](https://github.com/prometheus/client_golang) 指标记录 
1. 支持 [Swagger](https://github.com/swaggo/gin-swagger) 接口文档生成 
1. 支持 [GraphQL](https://github.com/99designs/gqlgen) 查询语言 
1. 支持 trace 项目内部链路追踪 
1. 支持 [pprof](https://github.com/gin-contrib/pprof) 性能剖析
1. 支持 errno 统一定义错误码 
1. 支持 [zap](https://go.uber.org/zap) 日志收集 
1. 支持 [viper](https://github.com/spf13/viper) 配置文件解析
1. 支持 [gorm](https://gorm.io/gorm) 数据库组件
1. 支持 [go-redis](https://github.com/go-redis/redis/v7) 组件
1. 支持 RESTful API 返回值规范
1. 支持 生成数据表 CURD、控制器方法 等代码生成器
1. 支持 [cron](https://github.com/jakecoffman/cron) 定时任务，在后台可界面配置
1. 支持 [websocket](https://github.com/gorilla/websocket) 实时通讯，在后台有界面演示
1. 支持 web 界面，使用的 [Light Year Admin 模板](https://gitee.com/yinqi/Light-Year-Admin-Using-Iframe)



