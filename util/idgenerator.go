package util

import (
	"crypto/md5"
	"time"
)

// IdGen 32b timestamp + 18b counter + 6b countspace + 2b reserved + 6b serverid
// 18b counter =  10b millisecond(0-999) + 8b counter(自增) 此处用8b 随机值代替
func IdGen() int64 {
	var result int64

	timeNowNano := time.Now().UnixNano()

	timeNowSec := (timeNowNano / 1e9) << 32

	timeNowMilli := ((timeNowNano / 1e6) % 1e3) << 22

	rand8Bit := int64(GetRand(256)) << 14 // 不包含256

	counterSpace := "counterSpace"
	counterSpaceBit := (int64(md5.Sum([]byte(counterSpace))[0]) % 64) << 8 //  保证6bit

	var serverId int64 = 63
	serverIdBit := serverId % 64

	result = timeNowSec + timeNowMilli + rand8Bit + counterSpaceBit + serverIdBit

	return result
}
