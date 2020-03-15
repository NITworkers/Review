USE web_db;
/*
  レビュー情報管理テーブル
  ・レビューそのものの内容のみ。
  ・評価はEvaluationテーブルで管理
*/
CREATE TABLE IF NOT EXISTS reviews (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT, /*DB管理用ID（自動採番）*/
  `created_at` datetime DEFAULT NULL, /*作成日時（GORM管理用）*/
  `updated_at` datetime DEFAULT NULL, /*更新日時（GORM管理用）*/
  `deleted_at` datetime DEFAULT NULL, /*削除日時（GORM管理用）*/
  `user_id` int(10) unsigned NOT NULL,  /*投稿ユーザID*/
  `store_name` varchar(255) DEFAULT NULL, /*店名*/
  `store_url` varchar(255) DEFAULT NULL, /*店のURL*/
  `text` varchar(255) DEFAULT NULL, /*店についてのコメント・評価*/
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES idpasses(`id`),  /*ユーザー情報管理テーブルとの紐づけ*/
  KEY `idx_idpasses_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*テストデータ登録用SQL文書（IDは自動採番）*/
INSERT INTO reviews(`user_id`, `store_name`, `store_url`, `text`) values
(2,"google", "http://wwww.google.co.jp", "useful to search");
INSERT INTO reviews(`user_id`, `store_name`, `store_url`, `text`) values
(3,"テスト", "http://wwww.google.co.jp", "日本語テスト");