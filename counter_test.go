package main

import (
	"MyItem/counter"
	"MyItem/counter1"
	"MyItem/counter2"
	"fmt"
	"testing"
)

func Counter() {
	counter.Flush2broker(5000, counter.FuncCbFlush)

	counter.Incr("get.called", 123)
	counter.Incr("get.called", 456)
}

func Counter1() {
	counter1.Flush2broker(5000, counter1.FuncCbFlush)

	counter1.Incr("get.called", 123)
	counter1.Incr("get.called", 456)
}

func Counter2() {
	counter2.Flush2broker(5000, counter2.FuncCbFlush)

	counter2.Incr("get.called", 123)
	counter2.Incr("get.called", 456)
}

func BenchmarkCounter(b *testing.B) {
	counter.Init()
	for i := 0; i < b.N; i++ {
		Counter()
	}
}

func BenchmarkCounter1(b *testing.B) {
	counter1.Init()
	for i := 0; i < b.N; i++ {
		Counter1()
	}
}

func BenchmarkCounter2(b *testing.B) {
	counter2.Init()
	for i := 0; i < b.N; i++ {
		Counter2()
	}
}

func main() {
	Counter1()
	// 防止线程退出
	var wait string
	_, _ = fmt.Scan(&wait)
}
