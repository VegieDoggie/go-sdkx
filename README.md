# Go-SDKx

提供断言式异常或断言式日志记录的接口

## 背景

减少Go语言的panic样板式代码

## 预览

- 断言式恐慌
- 断言式日志记录

说明: 同时传递多个参数时，函数会从最后一个参数开始检查，如果期间有符合断言的情况，则立即执行指令并返回。
## 快速开始

```cmd
go get github.com/VegetableDoggies/go-sdkx@v1.0.2
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
