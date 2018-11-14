# CloudGo

CloudGo 是一个基于 Gin 框架开发的简单的 WEB 服务程序，仅用于学习 WEB 服务器的工作原理，以及 Go 语言下 WEB 框架的使用。

## 使用指南

    go install github.com/siskonemilia/CloudGo
    CloudGo -h hostname -p port

## 为什么使用框架？

>不要重复造轮子

上面这句话完美地诠释了众多框架存在的重要性。因为在 WEB 课程上已经学习过服务器相关的知识，我们不需要再在轮子上耗费时间，直接使用现成的框架将会给我们带来开发速度和运行速度上的双重利好。同时，合理的使用框架可以显著地增强应用的安全性和可读性，为真正产品级的应用打下良好的基础。

## 为什么使用 Gin？

[Go WEB Framework Stars](https://github.com/mingrammer/go-web-framework-stars) 是 Github 上整理流行 Go WEB 框架的仓库。在这个仓库给出的列表中，[Gin](https://github.com/gin-gonic/gin) 雄踞榜首。作为一个易于使用的框架，Gin 还拥有其他框架无法比拟的超强性能，这一切都造就了它作为 Go WEB 框架王者的身份。顺带一提，Gin 的名字来源于鸡尾酒的基酒「Gin 金酒」，因为他的 API 是类似于另一个有名的 Go WEB 框架：「Martini 马天尼」（一种著名的，以金酒为唯一基酒的古典鸡尾酒饮品）。

## Curl 测试

首先，在本地创建两个终端。其中一个运行以下命令来启动服务器：

    go run main.go

另一个运行如下命令来通过 Curl 访问服务器：

    curl -v localhost:8080/ginTest

在服务器的终端，你可以看到如下信息：

    [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:   export GIN_MODE=release
    - using code:  gin.SetMode(gin.ReleaseMode)

    [GIN-debug] GET    /ginTest                  --> github.com/siskonemilia/CloudGo/routes.Router.func1 (3 handlers)
    [GIN-debug] Listening and serving HTTP on localhost:8080
    [GIN] 2018/11/14 - 09:26:03 | 200 |     200.348µs |       127.0.0.1 | GET      /ginTest

在 Curl 的终端，你可以看到如下信息：

    *   Trying ::1...
    * TCP_NODELAY set
    * Connection failed
    * connect to ::1 port 8080 failed: Connection refused
    *   Trying 127.0.0.1...
    * TCP_NODELAY set
    * Connected to localhost (127.0.0.1) port 8080 (#0)
    > GET /ginTest HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.54.0
    > Accept: */*
    >
    < HTTP/1.1 200 OK
    < Content-Type: application/json; charset=utf-8
    < Date: Wed, 14 Nov 2018 01:26:03 GMT
    < Content-Length: 71
    <
    * Connection #0 to host localhost left intact
    {"message":"You've successfully received a message from a gin server."}

Curl 测试完成，我们的 Gin 服务器完成了任务。值得注意的是，Gin 的响应耗时仅为微秒级别，性能非常出色。

## AB 测试

注意：macOS 的终端已经自带了 Apache Benchmark，我们无需下载和安装即可使用。

通过以下命令开始进行压力测试：

    ab -n 10000 -c 100 http://127.0.0.1:8080/ginTest

注意在 macOS 的 Apache Benchmark 不可以使用 localhost，否则会导致 `apr_socket_connect(): Invalid argument (22)` 的报错发生。测试结果如下：

    This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/

    Benchmarking 127.0.0.1 (be patient)
    Completed 100 requests
    Completed 200 requests
    Completed 300 requests
    Completed 400 requests
    Completed 500 requests
    Completed 600 requests
    Completed 700 requests
    Completed 800 requests
    Completed 900 requests
    Completed 1000 requests
    Finished 1000 requests


    Server Software:
    Server Hostname:        127.0.0.1
    Server Port:            8080

    Document Path:          /ginTest
    Document Length:        71 bytes

    Concurrency Level:      100
    Time taken for tests:   0.226 seconds
    Complete requests:      1000
    Failed requests:        0
    Total transferred:      194000 bytes
    HTML transferred:       71000 bytes
    Requests per second:    4429.91 [#/sec] (mean)
    Time per request:       22.574 [ms] (mean)
    Time per request:       0.226 [ms] (mean, across all concurrent requests)
    Transfer rate:          839.26 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0    1   1.0      0       5
    Processing:     0   21   8.5     21      60
    Waiting:        0   20   8.6     20      60
    Total:          1   21   8.3     21      61
    WARNING: The median and mean for the initial connection time are not within a normal deviation
            These results are probably not that reliable.

    Percentage of the requests served within a certain time (ms)
      50%     21
      66%     24
      75%     27
      80%     28
      90%     30
      95%     33
      98%     40
      99%     48
    100%     61 (longest request)