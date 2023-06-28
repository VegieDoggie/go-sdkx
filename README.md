# Go-SDKx

提供断言式异常或断言式日志记录的接口

## 背景

Go语言的错误判断太烦人了，于是专门写个工具来减少panic样板式代码，此外，额外提供了断言式log函数(便于编写简单测试脚本)。

## 预览

- 断言式恐慌
- 断言式日志记录

说明: 同时传递多个参数时，函数会从最后一个参数开始检查，如果期间有符合断言的情况，则立即执行指令并返回。
## 快速开始

```cmd
go get github.com/VegetableDoggies/go-sdkx@v1.0.0
```

```go
// 说明: 同时传递多个参数时，函数会从最后一个参数开始检查，如果期间有符合断言的情况，则立即执行指令并返回。

// err!=nil 时抛出异常
panicIf.Err(err)
panicIf.Errf("有个错误: %v", err)
panicIf.Err(gcron.AddSingleton(ctx, pattern, f)) // 更加偷懒的用法，比如在使用 goframe 定时任务时，你压根不想关心它有几个返回值，你只想包含错误时终止程序启动

// val==false 时抛出异常
numLike := "1"
val, b := new(big.Int).SetString(numLike, 10)
panicIf.Flase(b)
panicIf.Falsef("参数异常: %v %v", numLike, b)

...
...
```
