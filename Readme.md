# Balancer

>is a layer 7 load balancer that supports http and https, and it is also a go library that implements load balancing algorithms.

大佬的 Github 主页: [zehuamama](https://github.com/zehuamama/) 大佬的 Tiny Golang 教学系列, 知乎地址如下:
[马丸子](https://www.zhihu.com/people/ao-sha-xi-pan-37)

创建这个仓库的目的是用来学习 `Git` `Github` `Golang` `Balancer` 以及 负载均衡算法


# 学习进度
目前在做售后，时间比较少，每天回家只能看一小会儿；
基础也比较薄弱；

## 第二天
今天看到了 structure HTTPProxy 和 function NewHTTPProxy 这两部分；
主要了解了 http 标准库的一些结构体和接口

比如
`url.Parse(rawUrl string)(*url.URL, error)` 用来解析 url, 返回 `url.URL` 的指针对象和可能引发的错误;

`url.URL` 是标准库 `net/http/url` 中的一个结构体, 主要包括 url Scheme/Host/Path/RawPath/Opaque/Fragment/RawQuery 字段

`http.Request` http 包的标准库中的 Structure, 主要包括 Request 的方法/请求头/请求体/Host(请求的主机名和端口号)/Form/PostForm/RemoteAddr/RequestURL 等

`http.CanonicalHeaderKey` 方法用来返回标准请求头参数, 避免请求头大小写带来的问题;

