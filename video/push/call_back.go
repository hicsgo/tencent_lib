package push

/*
  腾讯云点播回调
*/
type VideoCallBack struct {
	EventType                         string                  `json:"EventType"`
	FileUploadEvent                   *FileUploadEvent        `json:"FileUploadEvent"`
	ProcedureStateChangeEvent         string                  `json:"ProcedureStateChangeEvent"`
	FileDeleteEvent                   string                  `json:"FileDeleteEvent"`
	PullCompleteEvent                 string                  `json:"PullCompleteEvent"`
	EditMediaComplete                 string                  `json:"EditMediaComplete"`
	WechatPublishComplete             string                  `json:"WechatPublishComplete"`
	TranscodeCompleteEvent            *TranscodeCompleteEvent `json:"TranscodeCompleteEvent"`
	ConcatCompleteEvent               string                  `json:"ConcatCompleteEvent"`
	ClipCompleteEvent                 string                  `json:"ClipCompleteEvent"`
	CreateImageSpriteCompleteEvent    string                  `json:"CreateImageSpriteCompleteEvent"`
	SnapshotByTimeOffsetCompleteEvent string                  `json:"SnapshotByTimeOffsetCompleteEvent"`
}
