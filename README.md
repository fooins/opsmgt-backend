# 运营管理系统-后端服务（opsmgt-backend）

福保成立初期设计开发了一套运营管理系统，支撑内部人员日常的运营管理工作，主要功能有：保单管理、理赔管理、产品管理、渠道管理、统计分析等等。本工程是其后端服务。[前端界面>>](../../../opsmgt-frontend)

- 整体[业务梳理](../../../.github/tree/main/profile/成立初期/成立初期业务梳理.md)和[系统设计](../../../.github/tree/main/profile/成立初期/成立初期系统设计.md)
- 相关[数据库表结构](../../../.github/tree/main/profile/成立初期/sql)
- [REST API 参考文档](./REST-API-reference-latest.md)
- [版本发布记录](./releases)

## 目录结构

```
├─ releases  // 发布信息目录
│
├─ src  // 源代码目录
│  ├─ components  // 业务组件目录
│  │
│  ├─ app.js  // 程序主应用实现
│  ├─ main.js  // 程序启动入口
│  ├─ router.js  // HTTP 路由实现
│  └─ server.js  // HTTP 服务实现
│
└─ REST-API-reference-latest.md  // REST API 参考文档
```

## 使用说明

1. 准备工作：安装 Go(1.19)、MySQL(8.x)、Redis(7.x) 和 Git。
2. 克隆代码：`git clone https://github.com/fooins/opsmgt-backend.git`。
3. 安装依赖：`go mod tidy`。
4. 安装 Air：`go install github.com/cosmtrek/air@latest`。
5. 启动程序：`air`。

## 环境变量

本项目根据 `GO_ENV` 环境变量来识别当前所处的运行环境类型，用于指导某些程序作出相应的不同的动作，比如日志组件在不同环境下会记录不同级别的日志。启动服务时请务必设置正确的环境变量，特别是生产环境。目前支持以下值：

| 环境变量值  | 说明     |
| ----------- | -------- |
| production  | 生产环境 |
| development | 开发环境 |
