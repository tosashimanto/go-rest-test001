package model

// 判定レコード

// 判定レコード作成POSTデータ
type JudgePost struct {
	PileId int64 `json:"pileId"` // 杭ID
	Index  int64 `json:"index"`  // 何本目
}

// 判定レコード作成結果
type JudgeResponse struct {
	JudgementId           int64              `json:"judgementId"`           // 判定ID
	CoreImageArray        []CoreImage        `json:"coreImageArray"`        // コア画像配列
	TrimmedCoreImageArray []TrimmedCoreImage `json:"trimmedCoreImageArray"` // 判定用コア画像配列
}

// 判定用コア画像
type CoreImage struct {
	Index     int64  `json:"index"`     // 何番目
	UploadUrl string `json:"uploadUrl"` // アップロードURL
}

// 判定用コア画像
type TrimmedCoreImage struct {
	Position  int    `json:"position"`  // 位置
	UploadUrl string `json:"uploadUrl"` // アップロードURL
}
