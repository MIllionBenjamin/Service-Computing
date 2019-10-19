# Cobra Generator -- Cobra生成器（Cobra使用说明）译文 - 2019.7.25版（当前最新版）
> 关键词：Golang、Cobra

> _Translated by MillionBenjamin(me)_
> <br>
> _Here is the [Original English Version](https://github.com/spf13/cobra/blob/master/cobra/README.md)_
> <br>
> 由MillionBenjamin翻译
> <br>
> 原文见[此链接](https://github.com/spf13/cobra/blob/master/cobra/README.md)
<br>


<script async src="//busuanzi.ibruce.info/busuanzi/2.3/busuanzi.pure.mini.js"></script>
<span id="busuanzi_container_page_pv">本文总阅读量<span id="busuanzi_value_page_pv"></span>次</span>

- [Cobra Generator -- Cobra生成器（Cobra使用说明）译文 - 2019.7.25版（当前最新版）](#cobra-generator----cobra%e7%94%9f%e6%88%90%e5%99%a8cobra%e4%bd%bf%e7%94%a8%e8%af%b4%e6%98%8e%e8%af%91%e6%96%87---2019725%e7%89%88%e5%bd%93%e5%89%8d%e6%9c%80%e6%96%b0%e7%89%88)
- [Cobra Generator](#cobra-generator)
    - [cobra init](#cobra-init)
    - [cobra add](#cobra-add)
    - [配置 cobra generator](#%e9%85%8d%e7%bd%ae-cobra-generator)


# Cobra Generator

Cobra提供了帮助你创建应用并添加命令的程序。这是把Cobra合并到你的应用的最简单方式。

为了使用Cobra命令，要用以下命令编译它。


    go get github.com/spf13/cobra/cobra


（译者注：在国内以此方式安装Cobra会遇到问题。解决方案可见[此博客](https://studygolang.com/articles/21268)）

这将会在你的 `$GOPATH/bin` 目录下创建Cobra的可执行程序。


### cobra init

`cobra init [app]` 命令会为你的应用创建最初的代码。这是一个非常有用的命令，它会让你的应用程序拥有正确的结构，这样你就可以体验到Cobra的所有优点。它还会自动应用你的许可到你的应用程序上。

Cobra init 非常的智能。你可以在当前目录运行这个命令，也可以指定一个指向已存在项目的相对路径。如果目录不存在，它会为你创建这些目录。

Cobra generator的更新使其与GOPATH解耦了。现在 `--pkg-name` 是必需的。

**注:** 现在，在非空目录中初始化不再会失败。

```
mkdir -p newApp && cd newApp
cobra init --pkg-name github.com/spf13/newApp
```

或

```
cobra init --pkg-name github.com/spf13/newApp path/to/newApp
```



### cobra add

一旦应用被初始化，Cobra就可以为你的应用创建额外的命令。不妨假设你创建了一个应用，并且你想要它有以下命令：

* app serve
* app config
* app config create

在你项目目录（即你的main.go文件所在目录）下运行以下命令：

```
cobra add serve
cobra add config
cobra add create -p 'configCmd'
```

*注意： 命名你的命令时，使用驼峰命名（camelCase），不要用蛇形命名（snake_case/snake-case）。否则你会遇到报错。例如 `cobra add add-user` 是错误的，而 `cobra add addUser` 可用的*

一旦你运行这三条命令，你会得到类似于此的目录结构：


```
  ▾ app/
    ▾ cmd/
        serve.go
        config.go
        create.go
      main.go
```

这时你可以运行 `go run main.go` 来运行你的应用。`go run
main.go serve`, `go run main.go config`, `go run main.go config create` 以及 `go run main.go help serve` 等等都会正常工作。

显然你还没有向这些命令假如你自己的代码。这些命令已经准备就绪，只待你给他们任务。编程愉快！


### 配置 cobra generator

如果你提供一些简单的配置文件，Cobra generator会变得更加易用。这些配置文件可以帮助你消除繁多flags里重复的信息。

例如 ~/.cobra.yaml 文件：

```yaml
author: Steve Francia <spf@spf13.com>
license: MIT
```

通过把 `license` 设为 `none` 你可以不添加许可。当然你也可以指定一个自定义的许可。

```yaml
license:
  header: This file is part of {{ .appName }}.
  text: |
    {{ .copyright }}

    This is my license. There are many like it, but this one is mine.
    My license is my best friend. It is my life. I must master it as I must
    master my life.
```

你也可以使用内建的许可。比如**GPLv2**, **GPLv3**, **LGPL**,
**AGPL**, **MIT**, **2-Clause BSD** 和 **3-Clause BSD**。

__[Support Me](https://millionbenjamin.github.io/Service-Computing/SupportMe)__