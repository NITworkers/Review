USE user_info;
/*ユーザー情報管理テーブル*/
CREATE TABLE IF NOT EXISTS `idpasses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT, /*DB管理用ID（自動採番）*/
  `created_at` datetime DEFAULT NULL, /*作成日時（GORM管理用）*/
  `updated_at` datetime DEFAULT NULL, /*更新日時（GORM管理用）*/
  `deleted_at` datetime DEFAULT NULL, /*削除日時（GORM管理用）*/
  `name` varchar(255) DEFAULT NULL,   /*ユーザー名 SNS内アカウント名*/
  `pass` varchar(255) DEFAULT NULL,   /*パスワード*/
  PRIMARY KEY (`id`),
  UNIQUE (`name`), /*ユーザーは重複不可*/
  KEY `idx_idpasses_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*テストデータ登録用SQL文書（IDは自動採番）*/
INSERT INTO idpasses(`name`, `pass`) values
("kunishige", "kunishig");
INSERT INTO idpasses(`name`, `pass`) values
("kiyopon", "kiyopon");