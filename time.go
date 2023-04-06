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

func GetCurrentYear() []byte {
    now := time.Now()
    year := now.Year()
    if year < 1000 {
        return []byte{'0', '0', '0', '0'}
    }
    var buf [4]byte
    for i := 3; i >= 0; i-- {
        buf[i] = byte('0' + year%10)
        year /= 10
    }
    return buf[:]
}

