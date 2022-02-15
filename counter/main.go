package counter

import (
	"fmt"
	"time"
)

type counterWriter struct {
	key string
	val int
}

var counter map[string]int    // 储存计数
var writer chan counterWriter // 用管道写入计数

// Init 计数器初始化
func Init() {
	counter = make(map[string]int)
	writer = make(chan counterWriter, 1000)
	go monitorCounter()
}

// Flush2broker 间隔t秒调用一次fn函数
func Flush2broker(t int, fn func()) {
	ticker := time.NewTicker(time.Duration(t) * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				fn()
			}
		}
	}()
}

// FuncCbFlush 重置计数器
func FuncCbFlush() {
	fmt.Println("重置计数器...")
	newCounter := make(map[string]int)
	counter = newCounter
}

// Incr 计数器根据key-value自增
func Incr(key string, val int) {
	if len(key) == 0 {
		fmt.Println("Incr的参数key为空，无效")
		return
	}
	if val == 0 {
		fmt.Println("Incr的参数val为0,无效")
		return
	}
	obj := counterWriter{key: key, val: val}
	writer <- obj
}

// Read 读取计数器
func Read() {
	fmt.Println("当前计数器存储：")
	for k,v:=range counter{
		fmt.Println(k,":",v)
	}
}

// 后台监听计数器
func monitorCounter() {
	for {
		select {
		case c := <-writer:
			counter[c.key] += c.val
		}
	}
}
