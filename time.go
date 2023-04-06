package cxtime

import (
	"time"
	"github.com/cloudxaas/gobytes"
)

func NanoNow() uint64 {
	return uint64(time.Now().UnixNano())
}
func NanoNowBytes() []byte {
	return cxbytes.Uint64ToBytes(NanoNow())
}
func GetYearNowInBytes(buf []byte) []byte {
    now := time.Now().Year()
    for i := len(buf) - 1; i >= 0; i-- {
        buf[i] = byte(now%10) + '0'
        now /= 10
        if now == 0 {
            return buf[i:]
        }
    }
    return nil
}
