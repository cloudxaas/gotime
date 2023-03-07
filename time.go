package cxtime

import (
	"github.com/cloudxaas/gobytes"
)

func NanoNow() uint64 {
	return uint64(time.Now().UnixNano())
}
func NanoNowBytes() []byte {
	return cxbytes.Uint64ToBytes(uint64(time.Now().UnixNano()))
}


