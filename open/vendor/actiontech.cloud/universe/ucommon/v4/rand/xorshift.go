package rand

import (
	"time"

	"github.com/lazybeaver/xorshift"
)

var xorshiftInstance = xorshift.NewXorShift64Star(uint64(time.Now().Nanosecond()))

func GetRandIntViaXorshift() uint64 {
	return xorshiftInstance.Next()
}
