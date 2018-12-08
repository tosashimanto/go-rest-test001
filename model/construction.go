package model

// list
type Constructions struct {
	// 工事配列
	ConstructionArray []Construction `json:"constructionArray"`
}

// 工事
type Construction struct {
	ConstructionId    int64  `json:"constructionId"`    // 工事ID
	ExecutionNum      string `json:"executionNum"`      // 工事番号
	PropertyName      string `json:"propertyName"`      // 物件名
	Address           string `json:"address"`           // 住所
	ConstManagerPhone string `json:"constManagerPhone"` // 管理者電話番号
}
