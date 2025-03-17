# xcdll
存放 xcgui.dll

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

