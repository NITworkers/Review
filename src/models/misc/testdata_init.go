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

	//全データ削除
	db.Unscoped().Delete(models.UserAccount{})
	db.Unscoped().Delete(models.Evaluation{})
	db.Unscoped().Delete(models.Review{})

	//テストデータ作成
	db.Create(&models.UserAccount{
		Name:     "User1",
		Password: "User1",
	})
	db.Create(&models.UserAccount{
		Name:     "User2",
		Password: "User2",
	})
	db.Create(&models.UserAccount{
		Name:     "User3",
		Password: "User3",
	})
}
