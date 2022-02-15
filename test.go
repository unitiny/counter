package main

import (
	"MyItem/counter"
	"fmt"
	"time"
)

func main() {
	counter.Init()
	counter.Flush2broker(5000, counter.FuncCbFlush)

	// 测试计数器运作
	counter.Incr("get.called", 123)
	counter.Incr("get.called", 456)
	counter.Incr("A", 0)
	counter.Incr("", 456)
	counter.Incr("B", 789)

	// 测试是否重置计数器
	counter.Read()
	time.Sleep(time.Second * 6)
	counter.Incr("C", 123)
	counter.Read()

	// 防止线程退出
	var wait string
	_, _ = fmt.Scan(&wait)
}
