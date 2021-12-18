package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/lazybeaver/xorshift"
)

var seq = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var xorshiftGenerator = xorshift.NewXorShift64Star(uint64(time.Now().Nanosecond()))

func init(){
    
}

func genRandomThreadID() string {
	l := 36
	a := rand.Intn(l * l * l)
	return fmt.Sprintf("%c%c%c", seq[a%l], seq[(a/l)%l], seq[(a/l/l)%l])
}

func genRandomThreadIDV2(length int) string {

	s := make([]byte, length)

	for i := 0; i < length; i++ {
		randInt := xorshiftGenerator.Next()
		s[i] = seq[randInt%36]
	}

    fmt.Println(string(s))

	return string(s)
}


func genRandomThreadIDV4(length int) string {
	u := uuid.NewString()
    s := []rune(u)
    return string(s[(36-length):36])
}

func main() {
	a := genRandomThreadIDV4(3)
	fmt.Printf("%s\n", a)
}
