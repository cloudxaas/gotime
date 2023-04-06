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
func YearNow(buf []byte) []byte {
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

func YearNowString() string {
	t := time.Now()
	return string([]byte{
		byte(t.Year()/1000) + '0',
		byte(t.Year()/100%10) + '0',
		byte(t.Year()/10%10) + '0',
		byte(t.Year()%10) + '0',
	})
}
