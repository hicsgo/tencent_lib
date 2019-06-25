package test

import (
	"encoding/json"
	"fmt"
	"tencent_lib/video"
	"testing"
	"time"
)

/*
 测试签名
 签名对比地址：https://video.qcloud.com/signature/ugcgenerate.html?_ga=1.205021971.2008746626.1558955159
*/
func TestSignStr(t *testing.T) {
	classID := int64(0)
	procedure := "QCVB_SimpleProcessFile({10,20},150,10,10)"
	taskNotifyMode := "None"
	taskPriority := int64(0)
	isOpen := false
	signature := &video.Signature{
		SecretId:         "AKIDh7ITUGxIyOnxu",
		SecretKey:        "mLhohMmxseBSUGn57tUv62kE",
		CurrentTimeStamp: time.Now().Unix(),
		ExpireTime:       time.Now().AddDate(0, 0, 5).Unix(),
		Random:           "1559620466456788899488994889941",
	}
	signature.SetClassID(&classID)
	a := struct {
		Type string `json:"type"`
	}{
		Type: "1",
	}
	aBytes, _ := json.Marshal(a)
	signature.SetSourceContext(string(aBytes))
	signature.SetPtt(procedure, taskNotifyMode, &taskPriority)
	signature.SetOneTimeValid(&isOpen)
	// secretId=xx&currentTimeStamp=1559620466&expireTime=1559620466&random=1559620466&classId=0&sourceContext={"type":"1"}&procedure=x&taskNotifyMode=x&taskPriority=0&oneTimeValid=0
	// 4Z6mmC2CVUciH5tXv9sSAp+PQLBzZWNyZXRJZD14eCZjdXJyZW50VGltZVN0YW1wPTE1NTk2MjA0NjYmZXhwaXJlVGltZT0xNTU5NjIwNDY2JnJhbmRvbT0xNTU5NjIwNDY2JmNsYXNzSWQ9MCZzb3VyY2VDb250ZXh0PSU3QiUyMnR5cGUlMjI6JTIyMSUyMiU3RCZwcm9jZWR1cmU9eCZ0YXNrTm90aWZ5TW9kZT14JnRhc2tQcmlvcml0eT0wJm9uZVRpbWVWYWxpZD0w

	result := signature.Sign()

	fmt.Println(result)
}
