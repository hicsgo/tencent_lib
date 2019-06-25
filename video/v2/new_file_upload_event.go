package v2

/*
上传回调回来的json示例
{
	"version": "4.0",
	"eventType": "NewFileUpload",
	"data": {
		"status": 0,
		"message": "",
		"vodTaskId": "",
		"fileId": "5285890790714697960",
		"fileUrl": "http://1257018105.vod2.myqcloud.com/c7596f40vodcq1257018105/1eb4c1fb5285890790714697960/TL916EePAC8A.mp4",
		"fileName": "1561433720",
		"continued": 0,
		"author": "",
		"sourceType": "ClientUpload",
		"sourceContext": "{\"sign_type\":\"post_video\"}",
		"metaData": {
			"audioDuration": 4.5279998779296875,
			"audioStreamList": [{
				"bitrate": 64386,
				"codec": "aac",
				"samplingRate": 44100
			}],
			"bitrate": 2189506,
			"container": "mov,mp4,m4a,3gp,3g2,mj2",
			"duration": 4,
			"floatDuration": 4.605999946594238,
			"height": 960,
			"md5": "",
			"rotate": 0,
			"size": 1266129,
			"totalSize": 1266129,
			"videoDuration": 4.605999946594238,
			"videoStreamList": [{
				"bitrate": 2125120,
				"codec": "h264",
				"fps": 24,
				"height": 960,
				"width": 540
			}],
			"width": 540
		}
	}
}
*/

/*
  上传文件回调事件
*/
type NewFileUploadEvent struct {
	FileId          string    `json:"fileId"`          //文件唯一 ID。
	FileName        string    `json:"fileName"`        //文件展示名称。
	CoverUrl        string    `json:"coverUrl"`        //文件封面地址。
	FileUrl         string    `json:"fileUrl"`         //文件播放地址。
	Author          string    `json:"author"`          //作者信息。
	SourceType      string    `json:"sourceType"`      //文件的上传来源。目前有 Record：录制，ClientUpload：客户端上传，ServerUpload：服务端上传。
	SourceContext   string    `json:"sourceContext"`   //上传时指定透传的字段，该字段目前最多256字节。
	StreamID        string    `json:"streamId"`        //推流 ID，录制上传特有。
	ProceduceTaskID string    `json:"procedureTaskId"` //该视频上传之后进行了指定流程，则该参数为流程任务 ID。
	TranscodeTaskID string    `json:"transcodeTaskId"` //如果该视频上传之后发起了转码，则该参数为转码任务 ID。
	MetaData        *MetaData `json:"metaData"`        //视频元信息
}

/*
  视频元信息
*/
type MetaData struct {
	Size            int64              `json:"size"`            //视频大小。单位：字节。
	Container       string             `json:"container"`       //容器类型，例如 m4a 和 mp4 等。
	Bitrate         int64              `json:"bitrate"`         //视频流码率平均值与音频流码率平均值之和。 单位：kbps。
	Height          int64              `json:"height"`          //视频流高度的最大值。单位：px。
	Width           int64              `json:"width"`           //视频流宽度的最大值。单位：px。
	Md5             string             `json:"md5"`             //视频的 MD5 值。
	Duration        int64              `json:"duration"`        //视频时长。单位：秒。
	Totate          int64              `json:"rotate"`          //视频拍摄时的选择角度。单位：度。
	AudioDuration   float64            `json:"audioDuration"`   //音频时长
	AudioStreamList []*AudioStreamInfo `json:"audioStreamList"` //音频集合
	VideoDuration   float64            `json:"videoDuration"`   //视频时长
	VideoStreamList []*VideoStreamInfo `json:"videoStreamList"` //视频流集合
}

//视频流信息
type VideoStreamInfo struct {
	Bitrate int64  `json:"bitrate"` //视频流的码率，单位：kbps。
	Height  int64  `json:"height"`  //视频流的高度，单位：px。
	Width   int64  `json:"width"`   //视频流的宽度，单位：px。
	Codec   string `json:"codec"`   //视频流的编码格式，例如 h264。
	Fps     int64  `json:"fps"`     //帧率，单位：hz。
}

//音频流信息
type AudioStreamInfo struct {
	Bitrate      int64  `json:"bitrate"`      //音频流的码率。 单位：kbps。
	SamplingRate int64  `json:"samplingRate"` //音频流的采样率。 单位：hz。
	Codec        string `json:"codec"`        //音频流的编码格式，例如 aac。
}
