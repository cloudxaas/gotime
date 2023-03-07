package cxtime

import (
	"github.com/cloudxaas/gobytes"
)

func GetTimeNanosecondNow() uint64 {
	return uint64(time.Now().UnixNano())
}
func GetTimeNanosecondNowBytes() []byte {
	return cxbytes.Uint64ToBytes(uint64(time.Now().UnixNano()))
}


