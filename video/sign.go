package video

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"tencent_lib/common"
)

/*
   腾讯云点播签名库
*/
type Signature struct {
	SecretId         string
	SecretKey        string
	CurrentTimeStamp int64
	ExpireTime       int64
	Random           string
}

/*
返回签名字符串
*/
func (s *Signature) Sign() string {
	signStr := ""
	if s != nil {
		//待签名字符串
		original := fmt.Sprintf("secretId=%s&currentTimeStamp=%d&expireTime=%d&random=%s", s.SecretId, s.CurrentTimeStamp, s.ExpireTime, s.Random)

		//encode
		encodeOriginal := url.PathEscape(original)

		//签名
		signatureTmp := common.HmacSha1(s.SecretKey, original)

		//拼接
		temp := fmt.Sprintf("%s%s", signatureTmp, encodeOriginal)

		//结果字符串
		signStr = base64.StdEncoding.EncodeToString([]byte(temp))
	}
	return signStr
}
