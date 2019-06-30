# parrot

![avatar](https://cdn.sinaimg.cn.52ecy.cn/large/005BYqpgly1g4eyqwpsj7j306106hq32.jpg)

parrot 是一个快速开发 Go 应用的 微服务 框架，他可以用来快速开发 API、RPC、CRON、MQC、WS、Web 及后端服务等各种应用。




特性


* RESTful 支持、MVC 模型，可以使用 duo 工具快速地开发应用。
* 可以监控 QPS、内存消耗、CPU 使用，以及 goroutine 的运行状况。
* 多集群模式支持
* 内置了强大的模块，包括 cache、日志记录、配置解析、性能监控、上下文操作、ORM。
* 跨平台支持
* 采用了 Go 原生的 http 包来处理请求，goroutine 的并发效率足以应付大流量的 API 应用和 RPC 应用。
* 服务注册与发现,统一的配置中心,支持优雅重启

# 快速入门

## 安装

parrot 包含一些示例应用程序以帮您学习并使用本框架,在这里找到相应的示例 [example](https://github.com/sereiner/example)。

您需要安装 Go 1.11+ 以确保所有功能的正常使用。

你需要安装或者升级 parrot 和 [duo]() 的开发工具:

	$ go get -u github.com/sereiner/parrot
	$ go get -u github.com/sereiner/duo

你还需要安装zookeeper,用来做服务发现和配置管理,你可以利用docker很方便的运行一个zookeeper集群,如果不想安装zookeeper,可以使用本地文件系统代替


为了更加方便的依赖管理操作，请使用 go 1.12 和 go mod。


想要快速建立一个应用来检测安装？

	$ cd 你的工作目录(不需要是GOPATH)
	$ duo new api apiserver
	$ cd apiserver

## 简单示例

下面这个示例程序将会在浏览器中打印 “Hello world”，以此说明使用 parrot 构建 API 应用程序是多么的简单！

	
    package main
    
    import (
        "github.com/sereiner/parrot/context"
        "github.com/sereiner/parrot/parrot"
    )
    
    func main() {

        app := parrot.NewApp(
            parrot.WithPlatName("apiserver"),
            parrot.WithSystemName("apiserver"),
            parrot.WithServerTypes("api"),
            parrot.WithDebug(),
            parrot.WithRegistry("fs://../"))  //使用本地文件系统作为注册中心
        app.API("/hello", helloWorld)
        app.Start()
    }
    func helloWorld(ctx *context.Context) (r interface{}) {
        return "hello world"
    }

把上面的代码保存为 hello.go，然后通过命令行进行编译并执行：

	$ go build -o hello hello.go
	
安装程序

    $ ./hello install

运行程序
    
    $ ./hello run

这个时候你可以打开你的浏览器，通过这个地址浏览 [http://localhost:8090/hello](http://localhost:8090/hello) 返回 {"data":"hello world"}。

那么上面的代码到底做了些什么呢？

1. 首先我们导入了包 `github.com/sereiner/parrot/parrot`。我们知道 Go 语言里面被导入的包会按照深度优先的顺序去执行导入包的初始化，parrot 包中会初始化一个 MicroAPP 的应用和一些参数。
2. 定义 main 函数，所有的 Go 应用程序和 C 语言一样都是 main 函数作为入口，所以我们这里定义了我们应用的入口。
3. NewApp 我们通过初始化函数,初始化了一个应用,并为该应用设置了一些启动参数。
4. API 注册路由，路由就是告诉 parrot，当用户通过http来请求的时候，该如何去调用相应的 API 服务器，这里我们注册了请求 `/hello` 的时候，请求到 `helloWorld` 函数。这里我们需要知道，API 函数的两个参数函数，第一个是路径，第二个是 handlerFunc 或者一个结构体。
6. Start 应用，最后一步就是把在步骤 2 中初始化的 MicroApp 开启起来，其实就是内部监听了 8090 端口。

停止服务的话，请按 `Ctrl+c`。