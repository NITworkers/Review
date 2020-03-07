### gin（アプリケーションサーバ）の起動方法

1. goコマンドで実行
    - ターミナル（シェル）で`go run src/main.go` を実行

2. cloud9で実行
    - cloud9のエディターでsrc/main.goを開き、上部の「RUN」ボタンをクリック
    - 続いて上部の「Preview」ボタンをクリックすると、簡易ブラウザで画面確認ができます



### JavaScriptのビルド方法

Reactを使う上で、手で書いたJavaScriptの一部はHTMLから呼び出す前に事前ビルドが必要。
yarnコマンド（webpackでのビルド）準備済み
- ターミナル（シェル）で`yarn build`を実行
- ビルドの設定ファイルはwebpack.config.js


### データベース初期化 / テストデータの準備方法

1. データベースの初期化
  - ターミナル（シェル）で`sudo mysql < src/models/misc/db_init.sql` を実行

2. テストデータの準備
  - src/models/misc/testdata_init.goを実行



### TODO
- git登録　⇒　過去にgithub上にレポジトリ作成しているが
  cloud9上のワークスペースがgit管理されなくなっている
  （.gitディレクトリを誰か削除した？）
  - →犯人はNITworker1
  - →とりあえずgitレポジトリにしました。GitHub(remote登録)は未実施

- ページ遷移を要整理。HTMLとGO側で整合性が必要。各ページも内容適当

- react対応。HTMLべた書きから、reactでのコンテンツ作成・モジュール化に移行したい
  - →NITworker1がやるよ

- DB設定を要整理。
    - パスワードがプレーンで保存されている



### 参考：ファイル/フォルダ体系
整理中…
- _old　いらないファイル置き場。整理後削除予定

- Pages　クライアントサイドのソース等置き場
- Pages/js/build　ビルドされたJavaScriptの置き場
- Pages/js/handmade　ビルド前やビルドしないJavaScriptの置き場

- src　サーバサイドのソース等置き場
- src/controllers　MVCのC（controller）処理のソース置き場
- src/models　MVCのM（modes）= DB周り処理のソース置き場

- node_modules　nodejs（サーバサイドJavaScript、yarnで使用）の稼働用
- package-lock.json　nodejs（サーバサイドJavaScript、yarnで使用）の稼働用
- package.json　nodejs（サーバサイドJavaScript、yarnで使用）の稼働用
- webpack.config.js　nodejs（サーバサイドJavaScript、yarnで使用）の稼働用
