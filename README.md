
<div align=center>

</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.24-blue"/>
<img src="https://img.shields.io/badge/gin-1.12.0-lightBlue"/>
<img src="https://img.shields.io/badge/gorm-1.31.1-red"/>
</div>



## 1. 基本介绍

### 1.1 项目介绍


> Gin-G是一个基于gorm和gin开发的后台web框架，源自gin-vue-admin，授权登录，路由，控制器，验证，模型等大不相同，是一个适合RESTful API 开发的web框架



### 1.2 版本列表

- master: 2.0.1

## 2. 使用说明

```

- golang版本 >= v1.24
- IDE推荐：Goland
- 初始化项目： 需要自行生成数据库和表，并在配置文件中添加数据库连接参数
```



- 使用git克隆本项目
    - ```git
        git clone 
        go mod tidy
      ```
    


## 3. 技术选型


- 路由：用`Gin`快速搭建基础restful风格API，`Gin`是一个go语言编写的Web框架。
- 中间件：支持`Gin`中间件
- 验证器：Gin 可以解析并验证请求的输入参数，包括XML，Json, Form等，Gin使用 go-playground/validator/v10 进行验证

- 数据库：使用`gorm`实现对数据库的基本操作,已添加对sqlite数据库的支持。
- 缓存：使用`Redis` 缓存数据。
- 配置文件：Viper支持在运行时让应用程序实时读取配置文件。
- 日志：使用`zerolog`实现日志记录。


## 4. 项目架构

### 4.1 目录结构

```
    ├─gin-g  	        （文件夹）
    │  ├─cmd            （程序入口及配置文件目录）
    │  ├─config         （配置包）
    │  ├─common  	    （通用方法及常量）
    │  ├─middleware     （中间件）
    │  ├─model          （结构体层）
    │  ├─router         （路由）
    │  └─utils	        （公共功能）

```


## 5.. 商用注意事项

如果您将此项目用于商业用途，请遵守Apache2.0协议并保留作者技术支持声明。
