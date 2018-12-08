package model

// チェック結果登録 POST JSONデータ
type PostReJudgeResult struct {
	// チェック結果配列
	CheckResultArray []PostJudgeCheckItem `json:"checkResultArray"`
}

// チェック結果
type PostJudgeCheckItem struct {
	CheckItemId        int64 `json:"checkItemId"`        // チェック項目ID
	TrimmedCoreImageId int64 `json:"trimmedCoreImageId"` // 判定用コア画像ID
	CheckStatus        bool  `json:"checkStatus"`        // チェック状態 YES: true, NO: false
}

// チェック結果登録応答
type ReJudgeResultResponse struct {
	// 判定結果
	Result int64 `json:"result"` // 0=OK, 1=NG, 2=未確定, 3=条件付きOK
	// 分割コア配列
	DividedCoreArray []ReJudgeResultDividedCore `json:"dividedCoreArray"`
}

// 分割コア
type ReJudgeResultDividedCore struct {
	Position int64 `json:"position"` // 全長コア内位置
	Result   int64 `json:"result"`   // 画像判定結果 0=○, 1=△, 2=×
}
