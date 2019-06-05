package push

/*
  视频转码成功事件
*/
type TranscodeCompleteEvent struct {
	TaskId      string               `json:"TaskId"`      //转码任务 ID。
	ErrCode     int64                `json:"ErrCode"`     //错误码 0：成功；其他值：失败。注意：此字段可能返回 null，表示取不到有效值。
	Message     string               `json:"Message"`     //错误信息。注意：此字段可能返回 null，表示取不到有效值。
	FileId      string               `json:"FileId"`      //被转码文件 ID。注意：此字段可能返回 null，表示取不到有效值。
	FileName    string               `json:"FileName"`    //被转码文件名称。注意：此字段可能返回 null，表示取不到有效值。
	Duration    int64                `json:"Duration"`    //视频时长，单位：秒。注意：此字段可能返回 null，表示取不到有效值。
	CoverUrl    string               `json:"CoverUrl"`    //封面地址。注意：此字段可能返回 null，表示取不到有效值。
	PlayInfoSet []*TranscodePlayInfo `json:"PlayInfoSet"` //视频转码后生成的播放信息。注意：此字段可能返回 null，表示取不到有效值。
}

/*
  转码文件播放信息
*/
type TranscodePlayInfo struct {
	Url        string `json:"Url"`        //播放地址。
	Definition int64  `json:"Definition"` //转码规格 ID，参见转码参数模板。
	Bitrate    int64  `json:"Bitrate"`    //视频流码率平均值与音频流码率平均值之和， 单位：bps。
	Height     int64  `json:"Height"`     //视频流高度的最大值，单位：px。
	Width      int64  `json:"Width"`      //视频流宽度的最大值，单位：px。
}
