package common

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func HmacSha1(key, original string) (sign string) {
	key1 := []byte(key)
	mac := hmac.New(sha1.New, key1)
	mac.Write([]byte(original))
	sign = fmt.Sprintf("%s", mac.Sum(nil))
	//notes:这个地方需要用%s
	return
}
