项目布局文档：https://github.com/golang-standards/project-layout/blob/master/README_zh-CN.md

### KnowFood Server

基于gin的API服务端，采用清洁架构设计。本项目采用传统的MVC模式，适用于业务API服务端开发。项目布局参考[project-layout](https://github.com/golang-standards/project-layout)。

## 环境要求

### 数据库配置
1. PostgreSQL数据库安装
```shell
podman pull docker.io/library/postgres:latest
```

2. 运行数据库容器
```shell
podman run --name postgres -e POSTGRES_USER=ethan -e POSTGRES_PASSWORD=AiRead@2024. -p 5432:5432 -v /home/ethan/code/airead_server/data -d [image_id/name]
```

3. 数据库初始化
```shell
# Linux环境
PGPASSWORD=AiRead@2024. psql -h 127.0.0.1 -p 5432 -U ethan
# macOS环境
PGPASSWORD=Airead@2024. psql -h 127.0.0.1 -p 5432 -U ethan
# 导入初始数据
PGPASSWORD=AiRead@2024. psql -h 127.0.0.1 -p 5432 -d ethan -U ethan -f init_pg.sql
```

### Go环境配置
1. 安装Go SDK (推荐Go 1.21+)
   - 从[Go官网](https://go.dev/dl/)下载并安装

2. 安装依赖注入工具wire
```shell
go install github.com/google/wire/cmd/wire@latest
```

3. 配置Go模块代理
```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

### 代理设置（如需要）
```shell
export https_proxy=http://192.168.8.162:7890 http_proxy=http://192.168.8.162:7890 all_proxy=socks5://192.168.8.162:7890
```

## 项目运行

1. 安装依赖
```shell
go mod tidy
```

2. 运行项目
```shell
# 开发环境运行
go run . -c ../conf/config.yml
# macOS环境运行
go run . -c ../conf/config_mac.yml

# 使用make构建
make        # 构建项目，输出在target目录
make clean  # 清理构建文件
```

## 项目特性

- 基于Gin框架的Web API服务
- 清洁架构设计，遵循依赖注入原则
- 使用Wire进行依赖注入
- 集成Viper进行配置管理
- 集成Gorm进行数据库操作
- Redis集成
- JWT认证支持
- Session管理
- 统一的API响应格式
- 全局错误处理
- 日志管理（Zap + Lumberjack）
- 优雅停机支持
- 使用Make进行项目构建
- 版本管理支持

## 项目结构

遵循清洁架构原则，项目分为以下层次：
- Handler层：处理HTTP请求
- Service层：业务逻辑处理
- Repository层：数据访问

## 文档

### 理论基础

#### 清洁架构 [(Robert C. Martin)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

![image](https://user-images.githubusercontent.com/8643542/159397149-17f58fba-a3c0-4874-b49a-ae724989af59.png)

按照依赖注入，面向对象编程思想，开闭原则，可拓展，可测性等原则来规划项目。

### 目前整合组件及实现功能

- 加入viper使用yml配置文件来配置项目信息，启动时指定不同的配置文件
- 优雅停机实现，停机时清理资源。
- 集成gorm 并自定义JsonTime 解决Json序列化和反序列化只支持UTC时间的问题（可以自定义时间格式）
  提供了部分demo，可以按照demo在项目中直接使用。
- 整合redis，开箱即用，通过yml文件对redis进行配置
- 整合zap，lumerjack 完善日志输出，日志分割。
- 集成jwt，提供demo代码，自定义授权失败成功等的响应格式，跟全局api响应格式统一
- 实现session管理
- md5, bcrypt和uuid生成工具包
- 应用统一封装响应格式，参照笔者参与的大型项目经验和规范。
- 项目全局错误码封装，go的error封装。
- 应用统一入口日志记录中间件实现，日志log_id透传。
- 添加makefile，可以使用make 命令进行编译，打包。
- 完善了项目版本管理，使用make命令编译后的项目可以方便跟踪线上发布版本
- 更多功能会根据个人实际开发中的经验添加，不过度封装，保持简单。
