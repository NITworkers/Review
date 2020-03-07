-- データベースの初期化
DROP DATABASE IF EXISTS web_db;
CREATE DATABASE web_db;

-- ユーザー情報の初期化
DROP USER 'web_user'@'localhost';
CREATE USER IF NOT EXISTS 'web_user'@'localhost' IDENTIFIED BY 'NITWeb_service';

-- ユーザー権限の初期化
GRANT ALL ON web_db.* TO 'web_user'@'localhost';