
### Go製マイグレーションツール sql-migrate

* インストール

```
$ go get github.com/rubenv/sql-migrate/...
```

* 設定ファイル

``` 
デフォルトはdbconfig.yml

dbconfig.yml
development:
  dialect: mysql
  dir: db/migrations
  datasource: tosashimanto1:123456@tcp(127.0.0.1)/go_test001?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true

production:
  dialect: postgres
  dir: migrations/postgres
  datasource: dbname=myapp sslmode=disable
  table: migrations
    
```

* コマンド

```
# ヘルプ
$ sql-migrate --help
```

* マイグレーション作成
```
# 以下のコマンドでは「20181005103536-create_users.sql」みたいなマイグレーションファイルが作成される
# ファイル名のdatetimeは自動で付与

$ sql-migrate new create_users
```

* マイグレーションの実行

```
$ sql-migrate up
# マイグレーションをdryrunで実行。実行予定のsqlが出力される
$ sql-migrate up -dryrun

```
* マイグレーションのロールバック
```
$ sql-migrate down
```

* マイグレーションの実行状態確認
```
$ sql-migrate status

```

* マイグレーションファイル

```
マイグレーション作成すると.sqlとしてファイルが生成される。sqlコメント文を使用しupの処理なのかdownの処理なのかを-- +migration Up、-- +migration Downでそれぞれ指定する。

20181222191112-create_users.sql
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (id int);

-- +migrate Down
DROP TABLE IF EXISTS users;

```