package v2

type ProcedureStateChanged struct {
	// 	data.status	String	任务流状态，有 PROCESSING 和 FINISH 两种。
	// data.errCode	Integer	错误码。0：成功。其他值：失败。
	// data.message	String	错误信息。
	// data.fileId	String	文件 ID。
	// data.metaData	Object	视频元信息，该字段一定存在，字段信息参见 metaData（视频元信息） 。
	// data.contentReviewList	Array	内容审核结果列表，字段信息参见 contentReviewList（内容审核列表） 。
	// data.aIAnalysisList	Array	智能分析结果列表，字段信息参见aIAnalysisList（智能分析列表）。
	// data.drm	Object	文件加密信息，用户在发起任务流时在 转码控制参数 指定了加密，该字段才存在。 字段信息参见 drm（视频加密信息） 。
	// data.processTaskList	Array	任务流包含的任务列表，字段信息参见 processTaskList（任务列表） 。
}
