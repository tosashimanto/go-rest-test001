package model

// 判定実行結果
type JudgeResult struct {
	// 判定結果
	CoreResult int64 `json:"coreResult"` // 0=OK, 1=NG, 2=未確定, 3=条件付きOK
	// 判定用コア画像配列
	TrimmedCoreImageArray []JudgeResultTrimmedCoreImage `json:"trimmedCoreImageArray"`
}

// 判定用コア画像
type JudgeResultTrimmedCoreImage struct {
	Position           int64 `json:"position"`           // 全長コア内位置
	TrimmedCoreImageId int64 `json:"trimmedCoreImageId"` // 判定用コア画像ID
	ImageResult        int64 `json:"imageResult"`        // 0=○, 1=△, 2=×
}
