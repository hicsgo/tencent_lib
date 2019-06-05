package push

/*
  文件上传成功回调事件
*/
type FileUploadEvent struct {
	FileId          string          `json:"FileId"`          //文件唯一 ID。
	ProcedureTaskId string          `json:"ProcedureTaskId"` //若视频上传时指定了视频处理流程，则该字段为流程任务 ID。注意：此字段可能返回 null，表示取不到有效值。
	MediaBasicInfo  *MediaBasicInfo `json:"MediaBasicInfo"`  //上传完成后生成的媒体文件基础信息。
}

/*
  媒体基础信息
*/
type MediaBasicInfo struct {
	Name          string           `json:"Name"`          //媒体文件名称。注意：此字段可能返回 null，表示取不到有效值。
	Description   string           `json:"Description"`   //媒体文件描述。注意：此字段可能返回 null，表示取不到有效值。
	CreateTime    string           `json:"CreateTime"`    //媒体文件的创建时间，使用 ISO 日期格式。注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime    string           `json:"UpdateTime"`    //媒体文件的最近更新时间（如修改视频属性、发起视频处理等会触发更新媒体文件信息的操作），使用 ISO 日期格式。注意：此字段可能返回 null，表示取不到有效值
	ExpireTime    string           `json:"ExpireTime"`    //媒体文件的过期时间，使用 ISO 日期格式。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。“9999-12-31T23:59:59Z”表示永不过期。注意：此字段可能返回 null，表示取不到有效值。
	ClassId       int64            `json:"ClassId"`       //媒体文件的分类 ID。注意：此字段可能返回 null，表示取不到有效值。
	ClassName     string           `json:"ClassName"`     //媒体文件的分类名称。注意：此字段可能返回 null，表示取不到有效值。
	ClassPath     string           `json:"ClassPath"`     //媒体文件的分类路径，分类间以“-”分隔，如“新的一级分类 - 新的二级分类”。注意：此字段可能返回 null，表示取不到有效值。
	CoverUrl      string           `json:"CoverUrl"`      //媒体文件的封面图片地址	注意：此字段可能返回 null，表示取不到有效值。
	Type          string           `json:"Type"`          //媒体文件的封装格式，例如 mp4、flv 等。注意：此字段可能返回 null，表示取不到有效值。
	MediaUrl      string           `json:"MediaUrl"`      //原始媒体文件的 URL 地址。注意：此字段可能返回 null，表示取不到有效值。
	SourceInfo    *MediaSourceData `json:"SourceInfo"`    //该媒体文件的来源信息。注意：此字段可能返回 null，表示取不到有效值。
	StorageRegion string           `json:"StorageRegion"` //媒体文件存储地区，如 ap-guangzhou，参见地域列表。注意：此字段可能返回 null，表示取不到有效值。
	TagSet        []string         `json:"TagSet"`        //媒体文件的标签信息。	注意：此字段可能返回 null，表示取不到有效值。
	Vid           string           `json:"Vid"`           //直播录制文件的唯一标识.注意：此字段可能返回 null，表示取不到有效值。
}

/*
  来源文件信息
*/
type MediaSourceData struct {
	// 媒体文件的来源类别：
	// Record：来自录制。如直播录制、直播时移录制等。
	// Upload：来自上传。如拉取上传、服务端上传、客户端 UGC 上传等。
	// VideoProcessing：来自视频处理。如视频拼接、视频剪辑等。
	// Unknown：未知来源。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceType string `json:"SourceType"`
	// 用户创建文件时透传的字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceContext string `json:"SourceContext"`
}
