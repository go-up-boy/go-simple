## go-simple

A framework to make golang MVC simple

## 0. go-simple 是什么？

go-simple是用gin搭建的mvc结构框架，是一个web api 快速开发工具，集成了开发中常用的模块，一键生成增删改查api代码，拥有很好的性能。<br>
普通mvc结构基本一致，不过将业务代码分层到模块中心，对多人协作开发较友好。
<br>
集成功能：
* JWT Auth认证中间件
* 限流中间件
* zap日志系统
* gorm
* gorm数据库迁移
* 图形验证码
* viper配置
* cobra一键生成api代码

## 1. 目录结构

开发者无需关心过多目录文件，按照目录约定开发即可，所有业务逻辑均在 模块中心 实现。
* app
  * cmd 命令行
  * http 控制器
  * middlewares 全局中间件
  * models 
  * modules 业务模块中心
    * user_module
      * user gorm模型
      * user_logics 业务层
      * user_policies 授权策略
      * user_services ……
      * user_utils ……
    * shop_module
    * order_module
    * ……
  * request 表单验证
* bootstrap
  * database.go
  * logger.go
  * redis.go
  * route.go
* config
  * app.go
  * captcha.go
  * config.go
  * database.go
  * jwt.go
  * log.go
  * redis.go
* database
* pkg 公共包
* routes 路由文件
* storage
  * —— logs 日志

## 2. 快速开始
    git clone -b dev https://github.com/zzzphp/go-simple.git test

    go mod download

    go mod tidy

    // 使用migrate生成数据库结构
    go run mian.go make migration user
    // 执行迁移
    go run mian.go migrate up
    // 一键生成模块 注意：控制器根目录为 http/controllers
    go run main.go make module [模块名] [控制器路径.../Name]
    // 修改生成后gorm模型结构即可

生成的代码包含增删改查等基本代码，只需要在 routes/api.go 添加相应的路由即可

## 3. 适合什么项目使用？

适合中小型项目无需复杂的架构设计、功能简单，增删改查较多，可以减少编写重复代码。<br>

## 4. 未完成的功能
1. 定时任务
2. 多数据库支持
3. 断点续传、秒传公共包等
4. rbac api
5. ……
