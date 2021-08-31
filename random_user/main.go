package main

import (
	"fmt"
	"time"
)

func main() {
	userList := []string{"zhangqiang", "huangweicen", "wujiabang", "wuyang", "bishaohua", "shibeibei"}

	// i := rand.Intn(len(userList))
	//
	// user := userList[i]
	//
	// fmt.Printf("Congratulations to %v", user)

	for {
		for _, u := range userList {
			fmt.Println(u)
			time.Sleep(time.Millisecond * 100)
		}
	}
}
