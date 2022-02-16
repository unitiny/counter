## 使用

counter里为计数器模块，counter_test.go为测试用例。
打开终端，输入命令即可看到效果
```go
go test -bench="." -benchmem
```

![img.png](img.png)

## 介绍

counter包用的是管道

counter1包用的是锁

counter2包用的是redis
