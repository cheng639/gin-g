
<div align=center>
<img src="http://qmplusimg.henrongyi.top/gvalogo.jpg" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.12-blue"/>
<img src="https://img.shields.io/badge/gin-1.4.0-lightBlue"/>
<img src="https://img.shields.io/badge/gorm-1.9.12-red"/>
</div>



## 1. 基本介绍

### 1.1 项目介绍


> Gin-G是一个基于gorm和gin开发的后台web框架，源自gin-vue-admin，授权登录，路由，控制器，验证，模型等大不相同，是一个适合RESTful API 开发的web框架



### 1.2 版本列表

- master: 1.0.0

## 2. 使用说明

```

- golang版本 >= v1.14
- IDE推荐：Goland
- 初始化项目： 需要自行生成数据库和表，并在配置文件中添加数据库连接参数
```



- 使用git克隆本项目
    - ```git
        git clone 
        go mod tidy
      ```
    


> Zap日志库使用指南&&配置指南

Zap日志库的配置选择在[config.yaml](./server/config.yaml)下的zap

```yaml
# zap logger configuration
zap:
  level: 'debug'
  format: 'console'
  prefix: '[GIN-G]'
  director: 'log'
  link_name: 'latest_log'
  show_line: true
  encode_level: 'LowercaseColorLevelEncoder'
  stacktrace_key: 'stacktrace'
  log_in_console: true
```

| 配置名         | 配置的类型 | 说明                                                         |
| -------------- | ---------- | ------------------------------------------------------------ |
| level          | string     | level的模式的详细说明,请看[zap官方文档](https://pkg.go.dev/go.uber.org/zap?tab=doc#pkg-constants) <br />info: info模式,无错误的堆栈信息,只输出信息<br />debug:debug模式,有错误的堆栈详细信息<br />warn:warn模式<br />error: error模式,有错误的堆栈详细信息<br />dpanic: dpanic模式<br />panic: panic模式<br />fatal: fatal模式<br /> |
| format         | string     | console: 控制台形式输出日志<br />json: json格式输出日志      |
| prefix         | string     | 日志的前缀                                                   |
| director       | string     | 存放日志的文件夹,修改即可,不需要手动创建                     |
| link_name      | string     | 在server目录下会生成一个link_name的[软连接文件](https://baike.baidu.com/item/%E8%BD%AF%E9%93%BE%E6%8E%A5),链接的是director配置项的最新日志文件 |
| show_line      | bool       | 显示行号, 默认为true,不建议修改                              |
| encode_level   | string     | LowercaseLevelEncoder:小写<br /> LowercaseColorLevelEncoder:小写带颜色<br />CapitalLevelEncoder: 大写<br />CapitalColorLevelEncoder: 大写带颜色 |
| stacktrace_key | string     | 堆栈的名称,即在json格式输出日志时的josn的key                 |
| log_in_console | bool       | 是否输出到控制台,默认为true                                  |

- 开发环境 || 调试环境配置建议
	- `level:debug`
	- `format:console`
	- `encode_level:LowercaseColorLevelEncoder`或者`encode_leve:CapitalColorLevelEncoder`
- 部署环境配置建议
	- `level:error`
	- `format:json` 
	- `encode_level: LowercaseLevelEncoder `或者 `encode_level:CapitalLevelEncoder`
	- `log_in_console: false` 
- <font color=red>建议只是建议,按照自己的需求进行即可,给出建议仅供参考</font>



## 3. 技术选型


- 路由：用`Gin`快速搭建基础restful风格API，`Gin`是一个go语言编写的Web框架。
- 中间件：支持`Gin`中间件
- 验证器：Gin 可以解析并验证请求的输入参数，包括XML，Json, Form等，Gin使用 go-playground/validator/v10 进行验证
- 控制器：嵌套的多层控制器，实现方法的继承，用户只需定义少量的方法，减少工作量。
- 数据库：使用`gorm`实现对数据库的基本操作,已添加对sqlite数据库的支持。
- 缓存：使用`Redis` 缓存数据。
- 优雅重启：使用 fvbock/endless优雅地重启或停止 web 服务器
- 配置文件：Viper支持在运行时让应用程序实时读取配置文件。
- 日志：使用`go-logging`实现日志记录。


## 4. 项目架构

### 4.1 目录结构

```
    ├─gin-g  	     （文件夹）
    │  ├─controller     （控制器）
    │  ├─config         （配置包）
    │  ├─core  	        （內核）
    │  ├─global         （全局对象）
    │  ├─initialiaze    （初始化）
    │  ├─middleware     （中间件）
    │  ├─model          （结构体层）
    │  ├─router         （路由）
    │  └─utils	        （公共功能）

```


## 5.. 商用注意事项

如果您将此项目用于商业用途，请遵守Apache2.0协议并保留作者技术支持声明。
