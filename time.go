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


