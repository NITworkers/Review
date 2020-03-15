USE web_db;
/*
  レビュー評価テーブル
  ・レビューについての他ユーザーの評価
  ・件数カウント用なので大した管理はしてない
*/
CREATE TABLE IF NOT EXISTS evaluations (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT, /*DB管理用ID（自動採番）*/
  `created_at` datetime DEFAULT NULL, /*作成日時（GORM管理用）*/
  `updated_at` datetime DEFAULT NULL, /*更新日時（GORM管理用）*/
  `deleted_at` datetime DEFAULT NULL, /*削除日時（GORM管理用）*/
  `review_id` int(10) unsigned NOT NULL,  /*評価対象レビュー*/
  `evaluate_user_id` int(10) unsigned NOT NULL,  /*評価ユーザID*/
  PRIMARY KEY (`id`),
  UNIQUE (`review_id`, `evaluate_user_id`), /*１レビューについて、ユーザーは1回のみ評価可能*/
  FOREIGN KEY (`review_id`) REFERENCES reviews(`id`)  /*レビュー情報との紐づけ*/
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*テストデータ登録用SQL文書（IDは自動採番）*/
INSERT INTO evaluations(`review_id`, `evaluate_user_id`) values
(3, 3);
INSERT INTO evaluations(`review_id`, `evaluate_user_id`) values
(4, 2);