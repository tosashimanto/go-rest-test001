package model

// 図面一覧
type Drawings struct {
	// 図面配列
	DrawingArray []Drawing `json:"drawingArray"`
}

type Drawing struct {
	DrawingName  string `json:"drawingName"`  // 図面名
	ReferenceUrl string `json:"referenceUrl"` // 参照URL
}
