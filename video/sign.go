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
	SecretId         string //云 API 密钥中的 SecretId，获取方式参考 客户端上传指引 - 获取云 API 密钥。
	SecretKey        string //云 API 密钥中的 SecretKey
	CurrentTimeStamp int64  //当前 Unix 时间戳。
	ExpireTime       int64  //签名到期 Unix 时间戳。expireTime = currentTimeStamp + 签名有效时长,签名有效时长最大取值为7776000 ，即90天。
	Random           string //构造签名明文串的参数，无符号32位随机数。
	sourceContext    string //客户端上传附带信息，在 事件通知 - 上传完成通知 中可以根据该字段识别一次上传行为，参见 客户端上传指引 - 事件通知。
	classID          *int64 //视频文件分类，默认为0。
	procedure        string //视频后续任务操作，详见 任务流综述。
	taskPriority     *int64 //视频后续任务优先级（仅当指定了 procedure 时才有效），取值范围为[-10, 10]，默认为0。
	taskNotifyMode   string /*任务流状态变更通知模式（仅当指定了 procedure 时才有效）。
		                            Finish：只有当任务流全部执行完毕时，才发起一次事件通知。
		                            Change：只要任务流中每个子任务的状态发生变化，都进行事件通知。
		                            None：不接受该任务流回调。
	                             	默认为 Finish。*/
	oneTimeValid *int64 //签名是否单次有效，参见 客户端上传指引 - 单次有效签名，默认为0，表示不启用；1表示签名单次有效，相关错误码详见下文的单次有效签名相关部分。
}

/*
返回签名字符串
*/
func (s *Signature) Sign() string {
	signStr := ""
	if s != nil && s.SecretId != "" && s.SecretKey != "" && s.CurrentTimeStamp != 0 && s.ExpireTime != 0 && s.Random != "" {
		//待签名字符串
		original := fmt.Sprintf("secretId=%s&currentTimeStamp=%d&expireTime=%d&random=%s", s.SecretId, s.CurrentTimeStamp, s.ExpireTime, s.Random)

		if s.classID != nil {
			original = fmt.Sprintf("%s&classId=%d", original, *s.classID)
		}

		if s.sourceContext != "" {
			original = fmt.Sprintf("%s&sourceContext=%s", original, s.sourceContext)
		}

		if s.procedure != "" && s.taskNotifyMode != "" && s.taskPriority != nil {
			original = fmt.Sprintf("%s&procedure=%s&taskNotifyMode=%s&taskPriority=%d", original, s.procedure, s.taskNotifyMode, *s.taskPriority)
		}

		if s.oneTimeValid != nil {
			original = fmt.Sprintf("%s&oneTimeValid=%d", original, *s.oneTimeValid)
		}
		fmt.Println(original)

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

/*
  视频文件分类，默认为0。
*/
func (s *Signature) SetClassID(classID *int64) {
	s.classID = classID
}

/*
  设置procedure,taskNotifyMode,taskPriority
*/
func (s *Signature) SetPtt(procedure, taskNotifyMode string, taskPriority *int64) {
	if procedure != "" && taskNotifyMode != "" && taskPriority != nil {
		s.procedure = procedure
		s.taskNotifyMode = taskNotifyMode
		s.taskPriority = taskPriority
	}
}

/*
  设置是否单次有效
*/
func (s *Signature) SetOneTimeValid(isOpen *bool) {
	if isOpen != nil {
		oneTimeValid := int64(0)
		if *isOpen {
			oneTimeValid = 1
		}
		s.oneTimeValid = &oneTimeValid
	}
}

/*
  设置SourceContext
*/
func (s *Signature) SetSourceContext(sourceContext string) {
	if sourceContext != "" {
		s.sourceContext = sourceContext
	}
}
