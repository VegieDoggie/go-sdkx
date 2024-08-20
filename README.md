# Go-SDKx

提供断言式异常接口

## 背景

减少Go语言的panic样板式代码

## 预览

- 断言式恐慌

## 快速开始

```cmd
go get github.com/VegieDoggie/go-sdkx@v1.0.5
```

```go
panicIf.Err(err)
panicIf.True(x == "x")
panicIf.False(false)
panicIf.Nil(nil)
panicIf.Empty(make(map[string]any))

logIf.Err(err)
logIf.True(x == "x")
logIf.False(false)
logIf.Nil(nil)
logIf.Empty(make(map[string]any))

```
