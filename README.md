### 期間目標

5月末までに、一段落つけたい
 - ユーザー投稿機能
 - ユーザー参照機能
とか

### デザインイメージ

https://xd.adobe.com/view/c72a41d6-ddee-4f9c-4da4-dee3ba705b4d-b11b/


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
  - 初期化せず、データ追加する際にはsrc/models/misc/testdata_add.goで対応します


### TODO

- git登録　⇒　過去にgithub上にレポジトリ作成しているが
  cloud9上のワークスペースがgit管理されなくなっている
  （.gitディレクトリを誰か削除した？）
  - →犯人はNITworker1
  　 →とりあえずgitレポジトリにしました。GitHub(remote登録)は未実施
  　 →(3/15更新)GitHub登録しました。https://github.com/NITworkers/Review

- ページ遷移を要整理。HTMLとGO側で整合性が必要。各ページも内容適当

- react対応。HTMLべた書きから、reactでのコンテンツ作成・モジュール化に移行したい
　 →NITworker1がやるよ
　 →(3/15更新)ログインフォームをreactで作りました。フォームのname属性変えたので、
  -  controller側もそこだけ変更してます。
  -  他画面も随時作成していくので、webpackのentryを複数にしました。
  -  それに伴って、[name].bundle.jsの形でbuildに出力するよう変更しています。

- DB設定を要整理。
    - パスワードがプレーンで保存されている
    - ゲストユーザーの管理方法がはっきりしてない
    - 　ログインなしの場合でも参照や評価は可能とする。
    　　→(3/22)ログインなしの場合は参照のみ
    - 　ただしController側���制御を考えたらダミーのユーザーID振ったほうがよさそう
    - 　→振り方、Viewからの情報の渡し方、コンテンツ制御方法を考える必要あり

### 情報共有

「notion」というすてきなツールを見つけました。
これで進捗管理とかwikiとか作れるか試してみます。
 - URL:
 　https://www.notion.so/362cf86569e14bcaae2ec10335a59bc8?v=091145ca899942bba1b543b65d1b98f8
 - 編集にはユーザ登録と招待が必要なので、気になった方はまずは登録を！
 - 参考サイト:
   https://unleashmag.com/2019/01/31/notion/

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
