package img

import (
	"encoding/base64"
	"fmt"
	"tencent_lib/common"
)

/*
   腾讯数据万象签名库
*/
type Signature struct {
	SecretId         string //云 API 密钥中的 SecretId，获取方式参考 客户端上传指引 - 获取云 API 密钥。
	SecretKey        string //云 API 密钥中的 SecretKey
	AppID            string //APPID
	Bucket           string //Bucket
	CurrentTimeStamp int64  //当前 Unix 时间戳。
	ExpireTime       int64  //签名到期 Unix 时间戳。expireTime = currentTimeStamp + 签名有效时长,签名有效时长最大取值为7776000 ，即90天。
	Random           string //构造签名明文串的参数，无符号32位随机数。
	FilePath         string
}

/*
返回签名字符串
*/
func (s *Signature) Sign() string {
	signStr := ""
	if s != nil {
		//待签名字符串
		original := fmt.Sprintf("a=%s&b=%s&k=%s&e=%d&t=%d&r=%s&f=%s", s.AppID, s.Bucket, s.SecretId, s.ExpireTime, s.CurrentTimeStamp, s.Random, s.FilePath)

		//签名
		signatureTmp := common.HmacSha1(s.SecretKey, original)

		//拼接
		temp := fmt.Sprintf("%s%s", signatureTmp, original)

		//结果字符串
		signStr = base64.StdEncoding.EncodeToString([]byte(temp))
	}
	return signStr
}
