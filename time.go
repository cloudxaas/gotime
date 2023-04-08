package cxtime

import (
	"time"
	"unsafe"

	cxbytes "github.com/cloudxaas/gobytes"	
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
	year := t.Year()

	buf := [4]byte{
		byte(year/1000) + '0',
		byte((year/100)%10) + '0',
		byte((year/10)%10) + '0',
		byte(year%10) + '0',
	}

	return *(*string)(unsafe.Pointer(&buf))
}
