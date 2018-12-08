package model

// チェックシート
type CheckSheet struct {
	// チェック項目配列
	CheckItemArray []CheckItem `json:"checkItemArray"`
}

// チェック項目
type CheckItem struct {
	CheckItemId  int64  `json:"checkItemId"` // チェック項目ID
	CheckContent string `json:"content"`     // チェック項目文
}
