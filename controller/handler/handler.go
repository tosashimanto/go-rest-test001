package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/tosashimanto/go-rest-test001/model"
	"github.com/tosashimanto/go-rest-test001/service/s3"
	"github.com/tosashimanto/go-rest-test001/util"
	"net/http"
	"strconv"
	"time"
)

// 1. トークン生成
func GetToken(c echo.Context) error {

	tokenPost := new(model.TokenJSONPost)
	if err := c.Bind(tokenPost); err != nil {
		return err
	}
	fmt.Println("Token=", tokenPost.Token)
	fmt.Println("Type=", tokenPost.Type)

	// Create token(JWT)
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	//claims["name"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	fmt.Println("token=" + t)
	c.Response().Header().Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", t))

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// 2. 工事一覧取得
func Constructions(c echo.Context) error {
	var constructionArray [5]model.Construction
	constructionArray[0] = model.Construction{
		ConstructionId:    1,
		ExecutionNum:      "100000000012345",
		PropertyName:      "○○様 ○○建設工事その1",
		Address:           "東京都中央区XX-X1",
		ConstManagerPhone: "000-0000-0000",
	}
	constructionArray[1] = model.Construction{
		ConstructionId:    2,
		ExecutionNum:      "100000000012346",
		PropertyName:      "○○様 ○○建設工事その2",
		Address:           "東京都中央区XX-X2",
		ConstManagerPhone: "000-0000-0002",
	}
	constructionArray[2] = model.Construction{
		ConstructionId:    3,
		ExecutionNum:      "100000000012347",
		PropertyName:      "○○様 ○○建設工事その4",
		Address:           "東京都中央区XX-X4",
		ConstManagerPhone: "000-0000-0003",
	}
	constructionArray[3] = model.Construction{
		ConstructionId:    4,
		ExecutionNum:      "100000000012348",
		PropertyName:      "○○様 ○○建設工事その5",
		Address:           "東京都中央区XX-X5",
		ConstManagerPhone: "000-0000-0004",
	}
	constructionArray[4] = model.Construction{
		ConstructionId:    5,
		ExecutionNum:      "100000000012349",
		PropertyName:      "○○様 ○○建設工事その6",
		Address:           "東京都中央区XX-X6",
		ConstManagerPhone: "000-0000-0006",
	}
	constructions := &model.Constructions{
		ConstructionArray: []model.Construction{
			constructionArray[0],
			constructionArray[1],
			constructionArray[2],
			constructionArray[3],
			constructionArray[4],
		},
	}
	jsonString, _ := json.Marshal(constructions)
	util.JSONFormatOut(jsonString)

	return c.JSON(http.StatusOK, constructions)
}

// 3. 図面一覧取得
func GetDrawings(c echo.Context) error {
	constructionId := c.Param("construction_id")
	fmt.Println("constructionId=", constructionId)

	var drawingArray [2]model.Drawing
	drawingArray[0] = model.Drawing{
		DrawingName:  "○○工面図.jpg",
		ReferenceUrl: "https://",
	}

	drawings := &model.Drawings{
		DrawingArray: []model.Drawing{
			drawingArray[0],
		},
	}

	jsonString, _ := json.Marshal(drawings)
	util.JSONFormatOut(jsonString)

	return c.JSON(http.StatusOK, drawings)
}

// 4. Pile一覧取得
func Piles(c echo.Context) error {
	constructionId := c.Param("construction_id")
	fmt.Println("constructionId=", constructionId)
	var constructionId2 int
	constructionId2, _ = strconv.Atoi(c.Param("construction_id"))
	fmt.Println("constructionId=", constructionId2)

	// 全長コア
	var fullCoreArray [4]model.FullCore
	fullCoreArray[0] = model.FullCore{
		Index:        1,
		JudgedNumber: 1,
	}
	fullCoreArray[1] = model.FullCore{
		Index:        2,
		JudgedNumber: 1,
	}
	fullCoreArray[2] = model.FullCore{
		Index:        1,
		JudgedNumber: 0,
	}
	fullCoreArray[3] = model.FullCore{
		Index:        2,
		JudgedNumber: 0,
	}
	// 杭情報
	var pileArray [2]model.Pile
	pileArray[0] = model.Pile{
		PileId:            9999001,
		PileNumber:        101,
		PileDepthUpper:    1.20,
		PileDepthLower:    12.20,
		MachineNumber:     101,
		PlannedCoreNumber: 5,
		Note:              "これはメモです。",
		RegisteredDate:    "2018/12/31",
		CollectionStopped: false,
		SurveyGroupNumber: 1,
		TargetLayerUpper:  0.40,
		TargetLayerLower:  2.40,
		FullCoreArray: []model.FullCore{
			fullCoreArray[0],
			fullCoreArray[1],
		},
	}
	pileArray[1] = model.Pile{
		PileId:            9999002,
		PileNumber:        102,
		PileDepthUpper:    2.20,
		PileDepthLower:    13.20,
		MachineNumber:     102,
		PlannedCoreNumber: 5,
		Note:              "これはメモです。その２",
		RegisteredDate:    "2018/11/02",
		CollectionStopped: false,
		SurveyGroupNumber: 2,
		TargetLayerUpper:  0.40,
		TargetLayerLower:  2.40,
		FullCoreArray: []model.FullCore{
			fullCoreArray[2],
			fullCoreArray[3],
		},
	}

	piles := &model.Piles{
		PileArray: []model.Pile{
			pileArray[0],
			pileArray[1],
		},
	}

	jsonString, _ := json.Marshal(piles)
	util.JSONFormatOut(jsonString)

	return c.JSON(http.StatusOK, piles)
}

// 5. 杭変更依頼

// 6. 判定レコード作成
func PostJudge(c echo.Context) error {

	judgePost := new(model.JudgePost)
	if err := c.Bind(judgePost); err != nil {
		return err
	}
	fmt.Println("PileId=", judgePost.PileId)
	fmt.Println("Index=", judgePost.Index)

	judgeResponse := setTrimImageArray(judgePost.PileId, judgePost.Index)
	jsonString, _ := json.Marshal(judgeResponse)
	util.JSONFormatOut(jsonString)

	return c.JSON(http.StatusOK, judgeResponse)
}

// 6.1 判定レコード作成応答のコア情報作成
func setTrimImageArray(pileId int64, index int64) *model.JudgeResponse {
	const trimImageNum = 11

	var trimmedCoreImageArray [trimImageNum]model.TrimmedCoreImage
	for i := 0; i < len(trimmedCoreImageArray); i++ {
		// 画像Upload URL
		objectKey := fmt.Sprintf("PileId_%d_Index_%d_Position_%d", pileId, index, i)
		fmt.Println("objectKey=", objectKey)
		uploadUrl, uploadUrlErr := s3.NewPutPreSignedS3URL(objectKey)
		if uploadUrlErr != nil {
			// return c.JSON(http.StatusInternalServerError, nil)
		}
		fmt.Println("uploadUrl=", uploadUrl)
		referenceUrl, referenceUrlErr := s3.NewGetPreSignedS3URL(objectKey)
		if referenceUrlErr != nil {
			// return c.JSON(http.StatusInternalServerError, nil)
		}
		fmt.Println("referenceUrl=", referenceUrl)

		trimmedCoreImageArray[i] = model.TrimmedCoreImage{
			Position:  i + 1,
			UploadUrl: uploadUrl,
		}
	}
	judgeResponse := &model.JudgeResponse{
		JudgementId:           999999,
		TrimmedCoreImageArray: trimmedCoreImageArray[:],
	}
	return judgeResponse
}

// 7. 判定履歴取得

// 8. 判定実行
func PutJudge(c echo.Context) error {

	judgementId := c.Param("judgement_id")
	fmt.Println("judgement_id=", judgementId)

	var trimmedCoreImageArray [11]model.JudgeResultTrimmedCoreImage
	trimmedCoreImageArray[0] = model.JudgeResultTrimmedCoreImage{
		Position:           1,
		TrimmedCoreImageId: 999901,
		ImageResult:        0, // 0=○, 1=△, 2=×
	}
	trimmedCoreImageArray[1] = model.JudgeResultTrimmedCoreImage{
		Position:           2,
		TrimmedCoreImageId: 999902,
		ImageResult:        0, // 0=○, 1=△, 2=×
	}
	trimmedCoreImageArray[2] = model.JudgeResultTrimmedCoreImage{
		Position:           3,
		TrimmedCoreImageId: 999903,
		ImageResult:        1, // 0=○, 1=△, 2=×
	}
	trimmedCoreImageArray[3] = model.JudgeResultTrimmedCoreImage{
		Position:           4,
		TrimmedCoreImageId: 999904,
		ImageResult:        1, // 0=○, 1=△, 2=×
	}
	trimmedCoreImageArray[4] = model.JudgeResultTrimmedCoreImage{
		Position:           5,
		TrimmedCoreImageId: 999905,
		ImageResult:        0, // 0=○, 1=△, 2=×
	}
	trimmedCoreImageArray[5] = model.JudgeResultTrimmedCoreImage{
		Position:           6,
		TrimmedCoreImageId: 999905,
		ImageResult:        0, // 0=○, 1=△, 2=×
	}
	trimmedCoreImageArray[6] = model.JudgeResultTrimmedCoreImage{
		Position:           7,
		TrimmedCoreImageId: 999905,
		ImageResult:        0, // 0=○, 1=△, 2=×
	}
	trimmedCoreImageArray[7] = model.JudgeResultTrimmedCoreImage{
		Position:           8,
		TrimmedCoreImageId: 999905,
		ImageResult:        0, // 0=○, 1=△, 2=×
	}
	trimmedCoreImageArray[8] = model.JudgeResultTrimmedCoreImage{
		Position:           9,
		TrimmedCoreImageId: 999905,
		ImageResult:        0, // 0=○, 1=△, 2=×
	}
	trimmedCoreImageArray[9] = model.JudgeResultTrimmedCoreImage{
		Position:           10,
		TrimmedCoreImageId: 999905,
		ImageResult:        0, // 0=○, 1=△, 2=×
	}
	trimmedCoreImageArray[10] = model.JudgeResultTrimmedCoreImage{
		Position:           11,
		TrimmedCoreImageId: 999905,
		ImageResult:        0, // 0=○, 1=△, 2=×
	}

	judgeResult := &model.JudgeResult{
		CoreResult:            2, // 0=OK, 1=NG, 2=未確定, 3=条件付きOK
		TrimmedCoreImageArray: trimmedCoreImageArray[:],
	}
	jsonString, _ := json.Marshal(judgeResult)

	util.JSONFormatOut(jsonString)

	return c.JSON(http.StatusOK, judgeResult)
}

// 9. チェック結果登録
func PostReJudge(c echo.Context) error {

	judgementId := c.Param("judgement_id")
	fmt.Println("judgement_id=", judgementId)

	postReJudgeResult := new(model.PostReJudgeResult)
	if err := c.Bind(postReJudgeResult); err != nil {
		return err
	}
	fmt.Println("CheckResultArray len=", len(postReJudgeResult.CheckResultArray))
	for i := 0; i < len(postReJudgeResult.CheckResultArray); i++ {
		fmt.Printf("No.%d CheckItemId=%d\n", i+1, postReJudgeResult.CheckResultArray[i].CheckItemId)
		fmt.Printf("No.%d TrimmedCoreImageId=%d\n", i+1, postReJudgeResult.CheckResultArray[i].TrimmedCoreImageId)
		fmt.Printf("No.%d CheckStatus=%t\n", i+1, postReJudgeResult.CheckResultArray[i].CheckStatus)
	}

	var dividedCoreArray [11]model.ReJudgeResultDividedCore
	dividedCoreArray[0] = model.ReJudgeResultDividedCore{
		Position: 1,
		Result:   0, // 0=○, 1=△, 2=×
	}
	dividedCoreArray[1] = model.ReJudgeResultDividedCore{
		Position: 2,
		Result:   0, // 0=○, 1=△, 2=×
	}
	dividedCoreArray[2] = model.ReJudgeResultDividedCore{
		Position: 3,
		Result:   2, // 0=○, 1=△, 2=×
	}
	dividedCoreArray[3] = model.ReJudgeResultDividedCore{
		Position: 4,
		Result:   0, // 0=○, 1=△, 2=×
	}
	dividedCoreArray[4] = model.ReJudgeResultDividedCore{
		Position: 5,
		Result:   0, // 0=○, 1=△, 2=×
	}
	dividedCoreArray[5] = model.ReJudgeResultDividedCore{
		Position: 6,
		Result:   0, // 0=○, 1=△, 2=×
	}
	dividedCoreArray[6] = model.ReJudgeResultDividedCore{
		Position: 7,
		Result:   0, // 0=○, 1=△, 2=×
	}
	dividedCoreArray[7] = model.ReJudgeResultDividedCore{
		Position: 8,
		Result:   0, // 0=○, 1=△, 2=×
	}
	dividedCoreArray[8] = model.ReJudgeResultDividedCore{
		Position: 9,
		Result:   0, // 0=○, 1=△, 2=×
	}
	dividedCoreArray[9] = model.ReJudgeResultDividedCore{
		Position: 10,
		Result:   0, // 0=○, 1=△, 2=×
	}
	dividedCoreArray[10] = model.ReJudgeResultDividedCore{
		Position: 11,
		Result:   0, // 0=○, 1=△, 2=×
	}

	reJudgeResultResponse := &model.ReJudgeResultResponse{
		Result:           0, // 0=OK, 1=NG, 2=未確定, 3=条件付きOK
		DividedCoreArray: dividedCoreArray[:],
	}
	jsonString, _ := json.Marshal(reJudgeResultResponse)
	util.JSONFormatOut(jsonString)

	return c.JSON(http.StatusOK, reJudgeResultResponse)
}

// 10. チェックシート取得
func GetCheckSheet(c echo.Context) error {
	const checkItemNo = 4
	var checkItemrray [checkItemNo]model.CheckItem
	checkItemrray[0] = model.CheckItem{
		CheckItemId:  8881,
		CheckContent: "あいうえお順",
	}
	checkItemrray[1] = model.CheckItem{
		CheckItemId:  8882,
		CheckContent: "アイウエオ順",
	}
	checkItemrray[2] = model.CheckItem{
		CheckItemId:  8883,
		CheckContent: "1111111111111111111",
	}
	checkItemrray[3] = model.CheckItem{
		CheckItemId:  8884,
		CheckContent: "2222222222222222222",
	}

	checkSheet := &model.CheckSheet{
		CheckItemArray: checkItemrray[:],
	}
	jsonString, _ := json.Marshal(checkSheet)
	util.JSONFormatOut(jsonString)

	return c.JSON(http.StatusOK, checkSheet)
}
