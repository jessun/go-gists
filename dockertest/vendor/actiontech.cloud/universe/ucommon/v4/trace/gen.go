package trace

import (
	"fmt"
	"math/rand"

	_rand "actiontech.cloud/universe/ucommon/v4/rand"
)

const seq = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const seqLen = 36 // const seq length
const TraceIDLength = 9

func NewTraceID() string {
	return genRandomThreadIdViaXorshift(TraceIDLength)
}

func genRandomThreadIdViaXorshift(length uint64) string {

	threadId := make([]byte, length, length)

	for i := 0; uint64(i) < length; i++ {
		randInt := _rand.GetRandIntViaXorshift()
		threadId[i] = seq[randInt%seqLen]
	}

	return string(threadId)
}

// Deprecated:  performance issue
func genRandomThreadId() string {
	seq := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	l := len(seq)
	a := rand.Intn(l * l * l)
	return fmt.Sprintf("%c%c%c", seq[a%l], seq[(a/l)%l], seq[(a/l/l)%l])
}
