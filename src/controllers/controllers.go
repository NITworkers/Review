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

func SignupView(ctx *gin.Context) {
	ctx.HTML(200, "signup.html", gin.H{})
}

func SigninView(ctx *gin.Context) {
	ctx.HTML(200, "signin.html", gin.H{})
}

func Signin(ctx *gin.Context) {
	var db = models.GetDbConnection()
	var userAccount models.UserAccount

	//	g.Request.ParseForm()　⇒　不要では？
	var queryName, isNotEmpty1 = ctx.GetQuery("_userName")
	var queryPassword, isNotEmpty2 = ctx.GetQuery("_userPass")

	applog.Logger.Print("userName:", queryName)
	applog.Logger.Print("userPass:", queryPassword)

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

func SearchReviews(ctx *gin.Context) {
	//作成中
	//「レビューから探す」で画面遷移時に実行
	//　　話題のレビューと新着レビューを3件ずつ取得
	var db = models.GetDbConnection()
	//var hotReviews [3]models.Review
	var newReviews []models.Review
	var Reviews []models.Review

	//話題のレビュー取得（何をもって話題とするか考える）

	//新着レビュー取得（作成日時が最新の３件）
	db.Order("created_at desc").Limit(3).Find(&Reviews)
	newReviews = Reviews

	applog.Logger.Print(len(newReviews))
	applog.Logger.Print(newReviews[0].StoreName)
	applog.Logger.Print(newReviews[1].StoreName)
	applog.Logger.Print(newReviews[2].StoreName)
	applog.Logger.Print(newReviews[0].CreatedAt)
	applog.Logger.Print(newReviews[1].CreatedAt)
	applog.Logger.Print(newReviews[2].CreatedAt)
	ctx.HTML(200, "SearchReviews.html", gin.H{})
}

/*
func SubmitReview(ctx *gin.Context) {
	/*
	 * PEND事項いくつかあり
	 * １．ユーザーID取得方法
	 * ２．WEB側ページ構成（投稿成功時、失敗時）

	var db = models.GetDbConnection()
	var Review models.Review
	/* ユーザーIDの取得方法がPEND
	 * ・View側から連携あり？
	 * ・ゲストは投稿できないとして、どこで制限かけるか（WEB側？）
	 * ※いったんページ側で制限かけて、UserId=NULLで入ってこない想定で作成
	var querySubmitUserId = ctx.GetQuery("userId")
	var queryStoreName, isNotEmpty1 = ctx.GetQuery("storename")
	var queryStoreURL = ctx.GetQuery("storeurl")
	var queryText = ctx.GetQuery("text")

	applog.Logger.Print("SubmitUserId:", querySubmitUserId)
	applog.Logger.Print("StoreName:", queryStoreName)
	applog.Logger.Print("StoreURL:", queryStoreURL)
	applog.Logger.Print("Text:", queryText)

	//空データ確認ロジック
	if !isNotEmpty1 || !(len(queryStoreName) > 0) {
		applog.Logger.Print("店名が入力なし") //テスト用
		goto error
	}

	Review = models.Review{
		SubmitUserId: querySubmitUserId,
		StoreName:    queryStoreName,
		StoreURL:     queryStoreURL,
		Text:         queryText,
	}
	db.Create(&Review)

	applog.Logger.Print("登録成功")
	goto success

success:
	ctx.HTML(200, "submited.html", gin.H{}) //ファイル名は仮
	return

error:
	ctx.HTML(200, "error.html", gin.H{}) //ファイル名は仮（このままでもよいか？）
	return

}
*/
/*
func EditReview(ctx *gin.Context) {
	//
	// 作成中（勝手に作ってくれてもいいよ）
	// すでに投稿されたレビューについて修正を行う


	//DBから対象のレビューを取得（必要情報はView側から連携あり想定）

	//対象の投稿についてアップデートをかける
}
*/
/*
func DeleteReview(ctx *gin.Context) {
	//
	 * 作成中（勝手に作ってくれてもいいよ）
	 * すでに投稿されたレビューについて削除を行う
	 *
	//

	//DBから対象のレビューを取得（必要情報はView側から連携あり想定）

	//対象の投稿についてデリートをかける
}

*/
/*
func EvaluateUp(ctx *gin.Context) {
	/*
	 * 作成中（勝手に作ってくれてもいいよ）
	 * 投稿に対し、評価を実施

	var db = models.GetDbConnection()
	var Evaluation models.Evaluation

	//評価対象のビュー、評価者の情報を設定（必要情報はView側から連携あり想定）
	/* ユーザーIDの取得方法がPEND
	 * ・View側から連携あり？
	 * ・ゲストの場合にどのような値が入ってくる？
	var queryReviewId, isNotNull = ctx.GetQuery("reviewId")
	var queryEvaluateUserId = ctx.GetQuery("evaluateuserid")

	applog.Logger.Print("ReviewId:", queryReviewId)
	applog.Logger.Print("EvaluateUserId:", queryEvaluateUserId)

	//空データ確認ロジック
	if !isNotNull || !queryReviewId == NULL {
		applog.Logger.Print("評価対象のレビュー情報なし") //テスト用
		goto error
	}

	/*PENDあり
	 * 評価者がゲストユーザーか、登録ユーザーかで分岐が必要
	 * ※現在は登録ユーザーのみを想定したロジック

	Evaluation = models.Evaluation{
		ReviewId:       queryReviewId,
		EvaluateUserId: queryEvaluateUserId,
	}
	db.Create(&Evaluation)

	applog.Logger.Print("登録成功")
	goto success

success:
	ctx.HTML(200, "submited.html", gin.H{}) //ファイル名は仮
	return

error:
	ctx.HTML(200, "error.html", gin.H{}) //ファイル名は仮（このままでもよいか？）
	return

}
*/
/*
func EvaluateDelete(ctx *gin.Context) {
	/*
	 * 作成中（勝手に作ってくれてもいいよ）
	 * 投稿に対し、一度入れた評価を取り消す。
	 * PEND事項：ゲストユーザーは取り消し不可にする？（ゲストユーザーの考え方整理で決定）
	var db = models.GetDbConnection()
	var Evaluation models.Evaluation

	//対象の評価情報を取得（必要情報はView側から連携あり想定）
	/* ユーザーIDの取得方法がPEND
	 * ・View側から連携あり？
	 * ・ゲストが取り消すことはなし想定でいいか？
	var queryReviewId, isNotNull1 = ctx.GetQuery("reviewId")
	var queryEvaluateUserId, isNotNull2 = ctx.GetQuery("evaluateuserid")

	applog.Logger.Print("ReviewId:", queryReviewId)
	applog.Logger.Print("EvaluateUserId:", queryEvaluateUserId)

	//空データ確認ロジック
	if !isNotNull1 || !queryReviewId == NULL {
		applog.Logger.Print("評価取り消し対象のレビュー情報なし") //テスト用
		goto error
	}
	if !isNotNull2 || !queryEvaluateUserId == NULL {
		applog.Logger.Print("評価取り消し対象のユーザー情報なし") //テスト用
		goto error
	}

	//DBから評価情報取得

	//対象データをDBから削除

	applog.Logger.Print("削除成功")
	goto success

success:
	ctx.HTML(200, "submited.html", gin.H{}) //ファイル名は仮
	return

error:
	ctx.HTML(200, "error.html", gin.H{}) //ファイル名は仮（このままでもよいか？）
	return

}
*/
