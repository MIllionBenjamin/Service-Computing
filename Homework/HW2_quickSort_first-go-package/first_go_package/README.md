# 第一个Go语言程序包
在 `testPackage` 文件夹下创建 `fpk.go` 文件，在该文件中定义 `fpk` 包


而后在 `test.go` 文件中写main函数。只需 `import "./testPackage"` 便可在 `test.go` 中使用 `fpk` 包中的函数。
<br>
例如：`fpk.Reverse("!oG ,olleH")`