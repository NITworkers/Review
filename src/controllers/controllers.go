//ginを起動するためのアプリケーション
package controllers

//依存パッケージ（他者作）のインポート
import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	//	"github.com/jinzhu/gorm"
	_ "net/http"
)

//依存パッケージ（自作）のインポート
import (
	"applog"
	"models"
)

func Top(ctx *gin.Context) {
	ctx.HTML(200, "top.html", gin.H{})
}

func Kiyopon(ctx *gin.Context) {
	ctx.HTML(200, "kiyopon.html", gin.H{})
}

func Signin(ctx *gin.Context) {
	var db = models.GetDbConnection()
	var userAccount models.UserAccount

	//	g.Request.ParseForm()　⇒　不要では？
	var queryName, isNotEmpty1 = ctx.GetQuery("name")
	var queryPassword, isNotEmpty2 = ctx.GetQuery("pass")

	applog.Logger.Print("name:", queryName)
	applog.Logger.Print("pass:", queryPassword)

	//空データ確認ロジック
	if !isNotEmpty1 || !(len(queryName) > 0) {
		applog.Logger.Print("nameなし or 空")
		goto error
	}
	if !isNotEmpty2 || !(len(queryPassword) > 0) {
		applog.Logger.Print("passなし or 空")
		goto error
	}

	// RecordNotFoundエラーを返すかチェックします
	if db.Where("Name = ?", queryName).First(&userAccount).RecordNotFound() {
		applog.Logger.Print("レコードが見つかりません")
		goto error
	} else {
		if queryPassword == userAccount.Password {
			applog.Logger.Print("ログイン成功")
			goto success
		} else {
			applog.Logger.Print("パスワードが不一致です")
			goto error
		}
	}

success:
	ctx.HTML(200, "kiyopon.html", gin.H{})
	return

error:
	ctx.HTML(200, "error.html", gin.H{})
	return
}

func Signup(ctx *gin.Context) {
	var db = models.GetDbConnection()
	var userAccount models.UserAccount

	//	ctx.Request.ParseForm()　⇒　不要では？
	var queryName, isNotEmpty1 = ctx.GetQuery("name")
	var queryPassword, isNotEmpty2 = ctx.GetQuery("pass")

	applog.Logger.Print("name:", queryName)
	applog.Logger.Print("pass:", queryPassword)

	//空データ確認ロジック
	if !isNotEmpty1 || !(len(queryName) > 0) {
		applog.Logger.Print("nameなし or 空") //テスト用
		goto error
	}
	if !isNotEmpty2 || !(len(queryPassword) > 0) {
		applog.Logger.Print("passなし or 空") //テスト用
		goto error
	}

	//重複Nameの確認用ロジック
	if db.Where("Name = ?", queryName).RecordNotFound() {
		userAccount = models.UserAccount{
			Name:     queryName,
			Password: queryPassword,
		}
		db.Create(&userAccount)

		applog.Logger.Print("登録成功")
		goto success
	} else {
		applog.Logger.Print("インプットユーザ名は使用済")
		goto error
	}

success:
	ctx.HTML(200, "kiyopon.html", gin.H{})
	return

error:
	ctx.HTML(200, "error.html", gin.H{})
	return
}

func GetReviews(ctx *gin.Context) {
	//作成中
}
