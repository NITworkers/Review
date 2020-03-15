//ginを起動するためのアプリケーション
package models

//依存パッケージ（外部）のインポート
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "time"
)

//依存パッケージ（自作）のインポート
import (
	"applog"
)

//構造体（ユーザアカウント）
type UserAccount struct {
	gorm.Model
	Name     string
	Password string
	Session  string
}

//構造体（評価）
type Evaluation struct {
	gorm.Model
	Review         Review `gorm:"foreignkey:ReviewId"` //ReviewIdを外部キーとして使用する
	ReviewId       int
	User           UserAccount `gorm:"foreignkey:EvaluateUserId"` //EvaluateUserId を外部キーとして使用する
	EvaluateUserId int
}

//構造体（レビュー情報）
type Review struct {
	gorm.Model
	SubmitUser   UserAccount `gorm:"foreignkey:SubmitUserId"` // UserId を外部キーとして使用する
	SubmitUserId int
	StoreName    string
	StoreURL     string
	Text         string
}

//DB接続情報の定義
const (
	dbms       = "mysql"
	dbUser     = "web_user"
	dbPassword = "NITWeb_service"
	dbProtocol = "tcp(localhost:3306)"
	dbName     = "web_db"
)

//DB接続情報保持用の変数
var dbConnection *gorm.DB

func Initialize() {
	tmpDbConnection, err := gorm.Open(dbms, dbUser+":"+dbPassword+"@"+dbProtocol+"/"+dbName+"?parseTime=true")

	if err != nil {
		applog.Logger.Print("models.go init error : データベース開けず！")
		panic(err)
	}
	dbConnection = tmpDbConnection
	applog.Logger.Print("データベース接続OK！")

	dbConnection.AutoMigrate(&UserAccount{})
	dbConnection.AutoMigrate(&Evaluation{})
	dbConnection.AutoMigrate(&Review{})
}

func GetDbConnection() *gorm.DB {
	return dbConnection
}

func Finalize() {
	dbConnection.Close()
}
