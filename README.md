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
// err!=nil 时抛出异常
panicIf.Err(err)
panicIf.Errf("有个错误: %v", err)
panicIf.Err(gcron.AddSingleton(ctx, pattern, f))

// val==false 时抛出异常
numLike := "1"
val, b := new(big.Int).SetString(numLike, 10)
panicIf.Flase(b)
panicIf.Falsef("参数异常: %v %v", numLike, b)
...
...
```
