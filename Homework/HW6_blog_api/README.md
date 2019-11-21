# 模仿Github设计一个博客网站的API

> 第一次尝试使用Go进行api开发，使用新手友好的Beego框架，并且还可以结合 Swagger 自动化生成文档。

- [模仿Github设计一个博客网站的API](#%e6%a8%a1%e4%bb%bfgithub%e8%ae%be%e8%ae%a1%e4%b8%80%e4%b8%aa%e5%8d%9a%e5%ae%a2%e7%bd%91%e7%ab%99%e7%9a%84api)
  - [一、数据库配置](#%e4%b8%80%e6%95%b0%e6%8d%ae%e5%ba%93%e9%85%8d%e7%bd%ae)
  - [二、创建API项目](#%e4%ba%8c%e5%88%9b%e5%bb%baapi%e9%a1%b9%e7%9b%ae)

## 一、数据库配置
创建数据库blog_api并选择使用
```sql
CREATE DATABASE blog_api;
USE blog_api;
```
Github的API[https://api.github.com/users/MIllionBenjamin](https://api.github.com/users/MIllionBenjamin)返回的用户信息内容如下：
```json
{
  "login": "MIllionBenjamin",
  "id": 37035090,
  "node_id": "MDQ6VXNlcjM3MDM1MDkw",
  "avatar_url": "https://avatars2.githubusercontent.com/u/37035090?v=4",
  "gravatar_id": "",
  "url": "https://api.github.com/users/MIllionBenjamin",
  "html_url": "https://github.com/MIllionBenjamin",
  "followers_url": "https://api.github.com/users/MIllionBenjamin/followers",
  "following_url": "https://api.github.com/users/MIllionBenjamin/following{/other_user}",
  "gists_url": "https://api.github.com/users/MIllionBenjamin/gists{/gist_id}",
  "starred_url": "https://api.github.com/users/MIllionBenjamin/starred{/owner}{/repo}",
  "subscriptions_url": "https://api.github.com/users/MIllionBenjamin/subscriptions",
  "organizations_url": "https://api.github.com/users/MIllionBenjamin/orgs",
  "repos_url": "https://api.github.com/users/MIllionBenjamin/repos",
  "events_url": "https://api.github.com/users/MIllionBenjamin/events{/privacy}",
  "received_events_url": "https://api.github.com/users/MIllionBenjamin/received_events",
  "type": "User",
  "site_admin": false,
  "name": "MillionBenjamin",
  "company": null,
  "blog": "",
  "location": "Guangzhou, Guangdong, China or Shenzhen, Guangdong, China",
  "email": null,
  "hireable": null,
  "bio": null,
  "public_repos": 15,
  "public_gists": 0,
  "followers": 3,
  "following": 2,
  "created_at": "2018-03-04T06:35:33Z",
  "updated_at": "2019-11-18T20:49:09Z"
}
```

参考Github API的内容，在数据库中创建表

```sql
CREATE TABLE `user` (
    login VARCHAR(20) NOT NULL,
    id INT(8) NOT NULL AUTO_INCREMENT,
    url VARCHAR(100) DEFAULT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT NOW(),
    other_information VARCHAR(100) DEFAULT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
```

向建好的表中插入数据：
```sql
INSERT INTO user
(login, id, url)
VALUES
('MillionBenjamin', 10000000, 'https://github.com/MIllionBenjamin');
```


## 二、创建API项目
安装Beego和Bee开发工具，具体可见[Beego官方文档安装教程](https://beego.me/quickstart)
```shell
go get -u github.com/astaxie/beego
go get -u github.com/beego/bee
```

安装go语言mysql操作库
```shell
go get -u github.com/go-sql-driver/mysql
```

然后使用bee api命令行创建项目。go-blog-api是项目名称，root是数据库用户，blog-api是数据库名称。注意这样创建的项目目录是在 `$GOPATH/src` 下的。
```shell
bee api go-blog-api -conn="root:YOURPASSWORD,@tcp(127.0.0.1:3306)/blog_api"
```

进入项目目录，输入运行命令
```shell
bee run -downdoc=true -gendoc=true
```

此时访问
```url
http://localhost:8080/v1/user
```
可以GET所有用户信息（user表中所有信息）
```json
[
  {
    "Login": "MillionBenjamin",
    "Id": 10000000,
    "Url": "https://github.com/MIllionBenjamin",
    "CreateAt": "2019-11-21T15:07:51+08:00",
    "OtherInformation": ""
  }
]
```

访问
```url
http://localhost:8080/v1/user/10000000
```
可以GET用户id为10000000的用户的信息，因为在数据库中建立user表时设定id为主码
```json
{
  "Login": "MillionBenjamin",
  "Id": 10000000,
  "Url": "https://github.com/MIllionBenjamin",
  "CreateAt": "2019-11-21T15:07:51+08:00",
  "OtherInformation": ""
}
```

访问
```url
http://localhost:8080/swagger/
```
进入Swagger，这是一个API生成文档和测试工具。
