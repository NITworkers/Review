//ginを起動するためのアプリケーション
package models

//依存パッケージ（他者作）のインポート
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type UserAccount struct {
	gorm.Model
	Name     string
	Password string
}

//構造体（レビュー情報）
type Review struct {
	gorm.Model
	submitUserId int
	storeName    string
	storeURL     string
	text         string
	evaluations  []Evaluation
}

//構造体（評価）
type Evaluation struct {
	gorm.Model
	reviewId       int
	evaluateUserId int
}

//DB接続情報の定義
const (
	dbms       = "mysql"
	dbUser     = "web_user"
	dbPassword = "NITWeb_service"
	dbProtocol = "tcp(localhost:3306)"
	dbName     = "user_info"
)

//DB接続情報保持用の変数
var dbConnection gorm.DB

func Initialize() {
	fmt.Println(dbConnection)
	dbConnection, err := gorm.Open(dbms, dbUser+":"+dbPassword+"@"+dbProtocol+"/"+dbName+"?parseTime=true")

	if err != nil {
		fmt.Println("models.go init error : データベース開けず！")
		panic(err)
	}
	fmt.Println("database connected")

	dbConnection.AutoMigrate(&UserAccount{})
}

func GetDbConnection() gorm.DB {
	return dbConnection
}

func Finalize() {
	dbConnection.Close()
}
