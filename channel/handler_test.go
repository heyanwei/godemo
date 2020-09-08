package channel

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// 创建一个字符串类型的通道
	channel := make(chan string)

	go sayHello("John", channel)
	go sayHello("Mike", channel)

	listenMe(channel)
}
