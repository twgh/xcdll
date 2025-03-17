<h1 align="center">xcdll</h1>
<p align="center">
    <a href="https://github.com/twgh/xcdll/releases"><img src="https://img.shields.io/badge/release-1.3.391-blue" alt="release"></a>
    <a href="http://www.xcgui.com"><img src="https://img.shields.io/badge/XCGUI-3.3.9.1-blue" alt="XCGUI"></a>
   <a href="https://golang.org"> <img src="https://img.shields.io/badge/golang-≥1.16-blue" alt="golang"></a>
    <a href="https://pkg.go.dev/github.com/twgh/xcdll"><img src="https://img.shields.io/badge/go.dev-reference-brightgreen" alt="GoDoc"></a>
    <a href="https://github.com/twgh/xcdll/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-brightgreen" alt="License"></a>
</p>



## 介绍

存放 xcgui.dll, 本仓库版本号可能和 [xcgui](https://github.com/twgh/xcgui) 仓库的不一样, 不用在意, 这里肯定会是最新的dll

## 获取

```bash
go get github.com/twgh/xcdll
```

## 使用方式

#### 1. 写出 dll 到程序运行目录, 在 import 中添加以下语句即可

```go
_ "github.com/twgh/xcdll/init"
```

#### 2. 写出 dll 到系统临时目录, 在 app.New() 之前调用

```go
xc.WriteDll(xcdll.DLL)
```

