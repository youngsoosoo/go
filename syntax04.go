package main

import (
	"fmt"
	"math/rand"
	"time" // seed 생성용 패키지
)

func main() {
	//seed설정
	seed := time.Now().Unix()
	rand.Seed(seed)

	for i := 1; i < 6; i++ {
		dice := rand.Intn(6) + 1
		fmt.Println(dice)
	}
}
