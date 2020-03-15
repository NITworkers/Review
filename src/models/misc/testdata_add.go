//ginを起動するためのアプリケーション
package main

//依存パッケージ（外部）のインポート
import ()

//依存パッケージ（自作）のインポート
import (
	"models"
)

func main() {
	models.Initialize()
	defer models.Finalize()

	var db = models.GetDbConnection()

	//現行のテストデータに付け足す際にはこちらを使用
	//（作成時間などに影響を受けるものがあり、自力作成はめんどそう）

	//テストデータ作成
	db.Create(&models.Review{
		SubmitUserId: 3,
		StoreName:    "ほげほげ2",
		StoreURL:     "HOGE",
		Text:         "ほげげほげ2",
	})
}
