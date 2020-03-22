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

	//静的コンテンツの配信
	router.Static("/static", "/home/ubuntu/environment/GourmetReview/Pages/js/build")

	//画像イメージの配信
	router.Static("/image", "/home/ubuntu/environment/GourmetReview/Pages/img")

	//トップ画面
	router.GET("/", controllers.Top)

	//kiyopon画面
	router.GET("/kiyopon", controllers.Kiyopon)

	//サインイン処理
	router.GET("/SignIn", controllers.Signin)

	//サインアップ処理
	router.GET("/SignUp", controllers.Signup)

	//サインイン画面
	router.GET("/SignInView", controllers.SigninView)

	//サインアップ画面
	router.GET("/SignUpView", controllers.SignupView)

	//ユーザー情報修正処理（ロジックあとで書く）

	//ユーザー情報削除処理（ロジックあとで書く）

	//レビューから探す
	router.GET("/SearchReviews", controllers.SearchReviews)

	//レビュー投稿（ロジックあとで書く）

	//レビュー編集（ロジックあとで書く）

	//レビュー削除（ロジックあとで書く）

	//レビュー評価（ロジックあとで書く）

	//レビュー評価取り消し（ロジックあとで書く）

	//稼働
	router.Run()
}
