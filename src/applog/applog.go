//ginを起動するためのアプリケーション
package applog

//依存パッケージ（外部）のインポート
import (
	"log"
	"os"
)

//Logger情報保持用の変数
var Logger *log.Logger = log.New(os.Stdout, "[GourmetReview App LOG]", log.LstdFlags|log.Lshortfile)
