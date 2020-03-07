//ginを起動するためのアプリケーション
package main

//依存パッケージ（外部）のインポート
import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "net/http"
)

//依存パッケージ（自作）のインポート
import (
	"controllers"
	"models"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("/home/ubuntu/environment/GourmetReview/Pages/html/*.html")

	models.Initialize()
	defer models.Finalize()

	//トップ・ログイン画面
	router.GET("/", controllers.Top)

	//kiyopon画面
	router.GET("/kiyopon", controllers.Kiyopon)

	//サインイン処理
	router.GET("/SignIn", controllers.Signin)

	//サインアップ
	router.GET("/SignUp", controllers.Signup)

	//稼働
	router.Run()
}
