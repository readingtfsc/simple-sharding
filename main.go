package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	ids := make([]int64, 0)
	for i := 0; i < 1000000; i++ {
		j := rand.Int63n(10000000000)
		ids = append(ids, j)
	}

	for _, id := range ids {
		fmt.Println("当前用户 ", id)
		shard := id % 10
		switch shard {
		case 0:
			fmt.Printf("用户 = %d 是 %d 号库 \n", id, shard)
		case 1:
			fmt.Printf("用户 = %d 是 %d 号库 \n", id, shard)
		case 2:
			fmt.Printf("用户 = %d 是 %d 号库 \n", id, shard)
		case 3:
			fmt.Printf("用户 = %d 是 %d 号库 \n", id, shard)
		case 4:
			fmt.Printf("用户 = %d 是 %d 号库 \n", id, shard)
		case 5:
			fmt.Printf("用户 = %d 是 %d 号库 \n", id, shard)
		case 6:
			fmt.Printf("用户 = %d 是 %d 号库 \n", id, shard)
		case 7:
			fmt.Printf("用户 = %d 是 %d 号库 \n", id, shard)
		case 8:
			fmt.Printf("用户 = %d 是 %d 号库 \n", id, shard)
		case 9:
			fmt.Printf("用户 = %d 是 %d 号库 \n", id, shard)
		}
	}
}