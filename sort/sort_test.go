package sort

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	//会被初始化为0值
	var persons [3]Person

	//匿名变量，只打印元素
	for _, v := range persons {
		log.Println(v)
	}

	//动态创建切片，提前分配空间，提高性能
	ps := make([]Person, 2)
	for _, v := range ps {
		log.Println(v)
	}

	//切片复制，如果切片大小不一致，个数按照小的
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1)
}
