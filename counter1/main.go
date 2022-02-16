package counter1

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	lock  sync.Mutex
	store map[string]int
}

var counter *Counter

// Init 计数器初始化
func Init() {
	counter = new(Counter)
	counter.store=make(map[string]int)
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
	//fmt.Println("重置计数器...")
	Init()
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

	counter.lock.Lock()
	defer counter.lock.Unlock()

	counter.store[key] += val
}

// Read 读取计数器
func Read() {
	counter.lock.Lock()
	defer counter.lock.Unlock()

	//fmt.Println("当前计数器存储：")
	for k, v := range counter.store {
		fmt.Println(k, ":", v)
	}
}
