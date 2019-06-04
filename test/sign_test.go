package test

import (
	"fmt"
	"tencent_lib/video"
	"testing"
)

/*
 测试签名
 签名对比地址：https://video.qcloud.com/signature/ugcgenerate.html?_ga=1.205021971.2008746626.1558955159
*/
func TestSignStr(t *testing.T) {
	signature := &video.Signature{
		SecretId:         "xx",
		SecretKey:        "xx",
		CurrentTimeStamp: 1559620466,
		ExpireTime:       1559620466,
		Random:           "1559620466",
	}
	result := signature.Sign()

	fmt.Println(result)
	//result:TsWr05NxibILiwkFmVkp3E7Pr/FzZWNyZXRJZD14eCZjdXJyZW50VGltZVN0YW1wPTE1NTk2MjA0NjYmZXhwaXJlVGltZT0xNTU5NjIwNDY2JnJhbmRvbT0xNTU5NjIwNDY2
}
