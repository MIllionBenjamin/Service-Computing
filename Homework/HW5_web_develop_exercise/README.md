# 简单的Go语言Web应用

> 第一次尝试使用Go进行Web开发，实现了Hello World服务端和客户端

- [简单的Go语言Web应用](#%e7%ae%80%e5%8d%95%e7%9a%84go%e8%af%ad%e8%a8%80web%e5%ba%94%e7%94%a8)
	- [一、服务端](#%e4%b8%80%e6%9c%8d%e5%8a%a1%e7%ab%af)
		- [1. `http.HandleFunc`](#1-httphandlefunc)
		- [2. `http.ListenAndServe`](#2-httplistenandserve)
	- [二、客户端](#%e4%ba%8c%e5%ae%a2%e6%88%b7%e7%ab%af)
		- [1. `http.Get`](#1-httpget)
	- [三、代码运行说明](#%e4%b8%89%e4%bb%a3%e7%a0%81%e8%bf%90%e8%a1%8c%e8%af%b4%e6%98%8e)

## 一、服务端
### 1. `http.HandleFunc`
```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```
该方法接收两个参数，一个是路由匹配的字符串（ url 路径），另外一个是 `func(ResponseWriter, *Request)` 类型的函数（实际的处理对象）。

`HandleFunc` 实质上是对更底层结构的封装。`HandleFunc` 在 `DefaultServeMux` 中注册一个处理给定模式的处理者函数。`ServeMux` 文档中解释了模式之间是如何关联的。

示例（来自于Go语言官方英文文档，下同）
```go
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 2. `http.ListenAndServe`
```go
func ListenAndServe(addr string, handler Handler) error
```
`ListenAndServe` 监听TCP网络地址（addr）并调用带有处理程序的服务器来处理传入的连接。已接受的连接会被配置，使得TCP保持工作。

参数中的处理程序handler通常被设为 `nil`，这种情况下会使用默认的底层控件 `DefaultServeMux`。

`ListenAndServe` 函数总是返回非 `nil` 的错误error。

示例
```go
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## 二、客户端
### 1. `http.Get`
```go
func Get(url string) (resp *Response, err error)
```
`Get` 向指定url提出一个GET请求，如果收到的回应是以下重定向码，`Get` 会遵循重定向，最多进行十次重定向。
```
301 (Moved Permanently)
302 (Found)
303 (See Other)
307 (Temporary Redirect)
308 (Permanent Redirect)
```
如果进行过多重定向或者出现HTTP协议错误，`Get` 会返回一个错误。非2xx回应不会产生错误。所有返回的错误都是 `*url.Error` 类型。如果请求超时或被中断，`url.Error` 中的Timeout方法会返回true。

如果err是nil值，resp总会包含一个非空的resp.Body。在完成对resp.Body的读取后，应当关闭它。

`Get` 是对 `DefaultClient.Get` 的包装。如果要用自定义头部进行请求，请使用 `NewRequest` 和 `DefaultClient.Do` 。

示例
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
```

## 三、代码运行说明
先运行 `hello_server.go`。成功运行后在浏览器访问 `localhost:8080/hello` 可以看到"Hello, Go Web!"的字样。<br>
这时运行 `hello_client.go`。命令行会打印出GET得到的内容，也是"Hello, Go Web!"。<br>
可以在不同目录下运行这两份代码。