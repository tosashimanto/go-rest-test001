package model

// 杭一覧
type Piles struct {
	// 杭配列
	PileArray []Pile `json:"pileArray"`
}

// 杭
type Pile struct {
	PileId            int64      `json:"pileId"`            // 杭ID
	PileNumber        int64      `json:"pileNumber"`        // 杭番号
	PileDepthUpper    float32    `json:"pileDepthUpper"`    // 杭天端深さ
	PileDepthLower    float32    `json:"pileDepthLower"`    // 杭下端深さ
	MachineNumber     int64      `json:"machineNumber"`     // 施工機番号
	PlannedCoreNumber int64      `json:"plannedCoreNumber"` // 採取予定本数
	Note              string     `json:"note"`              // 備考
	CollectionStopped bool       `json:"collectionStopped"` // 採取中止 trueならステータス：中止
	SurveyGroupNumber int64      `json:"surveyGroupNumber"` // 検査対象群番号
	TargetLayerUpper  float32    `json:"targetLayerUpper"`  // 設計対象層上端
	TargetLayerLower  float32    `json:"targetLayerLower"`  // 設計対象層下端
	RegisteredDate    string     `json:"registeredDate"`    // 登録日
	FullCoreArray     []FullCore `json:"fullCoreArray"`     // 全長コア配列
}

type FullCore struct {
	Index        int64 `json:"index"`        // 何本目
	JudgedNumber int64 `json:"judgedNumber"` // 判定回数
}
