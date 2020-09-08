package channel

import (
	"log"
	"time"
)

func sayHello(name string, channel chan<- string) {
	for {
		str := "Say hello to " + name
		channel <- str

		time.Sleep(time.Second)
	}
}

func listenMe(channel <-chan string) {
	for {
		// 从通道中取出数据, 此处会阻塞直到信道中返回数据
		message := <-channel
		// 打印数据
		log.Println(message)
	}
}
