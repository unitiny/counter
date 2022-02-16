package counter2

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var conn redis.Conn

// Init 计数器初始化
func Init() {
	r, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis Dial", err)
	}
	conn = r
	FuncCbFlush()
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
	data, err := redis.Values(conn.Do("hkeys", "counter"))
	if err != nil {
		fmt.Println("redis hkeys err:", err)
		return
	}
	for _, v := range data {
		_, _ = conn.Do("hdel", "counter", v)
	}
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
	_, err := conn.Do("hset", "counter", key, val)
	if err != nil {
		fmt.Println("redis hset err:", err)
	}
}

// Read 读取计数器
func Read() {
	//fmt.Println("当前计数器存储：")
	data, err := redis.Values(conn.Do("hgetall", "counter"))
	if err != nil {
		fmt.Println("redis values err:", err)
		return
	}
	for k, v := range data {
		fmt.Printf("%s ", v.([]byte))
		if k%2==1{
			fmt.Println("")
		}
	}
}