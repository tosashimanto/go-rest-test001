package test

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tosashimanto/go-rest-test001/db/gorm_test/db_connect"
	"github.com/tosashimanto/go-rest-test001/db/gorm_test/model"
)

func main() {
	db := db_connect.GormConnect()
	fmt.Println(db.Value)
	// 実行完了後DB接続を閉じる
	defer db.Close()

	// テーブル名が複数系でない場合、これを指定すること
	db.SingularTable(true)
	db.LogMode(true)

	var test002 model.Test002
	var test002a model.Test002

	var count = 0
	db.Debug().Find(&test002).Count(&count)
	db.Debug().Find(&test002a)
	if count == 0 {
		fmt.Println("該当レコードなし")
	} else {
		fmt.Println("Name:" + test002.Name)
		fmt.Println("RegisteredDate:" + (test002.CreatedAt).String())
		s0 := fmt.Sprintf("%d", test002.ConstructionID)
		fmt.Println("ConstructionID: " + s0)
		s1 := fmt.Sprintf("%d", test002.CreatedBy)
		fmt.Println("CreatedBy: " + s1)
		s2 := fmt.Sprintf("%d", test002.Value)
		fmt.Println("Value: " + s2)

		fmt.Println("UpdatedAt:" + (test002.UpdatedAt).String())
		fmt.Println("RegisteredDate:" + (test002.CreatedAt).String())

		fmt.Println("aa__Name:" + test002a.Name)
		fmt.Println("aa__CreatedAt:" + test002a.CreatedAt.String())
		s1a := fmt.Sprintf("%d", test002a.CreatedBy)
		fmt.Println("CreatedBy: " + s1a)
		s2a := fmt.Sprintf("%d", test002a.UpdatedBy)
		fmt.Println("UpdatedBy: " + s2a)
	}

}
