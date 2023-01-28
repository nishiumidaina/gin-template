# gin-template  
![Go](https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=plastic)
![MySQL](https://img.shields.io/badge/-Mysql-4479A1.svg?logo=mysql&style=plastic)  
 Dockerで使用するGinのテンプレート。  
 [MySQLテンプレート](https://github.com/nishiumidaina/mysql-template)を使用するとexample.envの設定でDB接続ができます。
## 構成
[Gin](https://github.com/gin-gonic/gin)：フレームワーク  
[GORM](https://gorm.io/ja_JP/docs/index.html)：ORMライブラリ  
[golang-migrate](https://github.com/golang-migrate/migrate)：マイグレーションツール  
```
gin-template
├── app
│   ├── src
|   |   ├──controllers
|   |   |   ├──router.go(ルーティングの設定)
|   |   |   └──index_controller.go
|   |   |   └──user_controller.go
|   |   |
|   |   ├──database(マイグレーション・SQLファイル)
|   |   ├──models(DB処理)
|   |   |   ├──databese.go(DBの接続設定)
|   |   |   └──user_model.go
|   |   |
|   |   ├──tmp(errorログ)
|   |   ├──.air.toml
|   |   ├──go.mod
|   |   ├──go.sum
|   |   ├──main.go
|   |   └──.env(※example.envを参考にして配置する)
|   |
│   └── Dockerfile
└── docker-compose.yaml
```
## 環境構築
1.リポジトリをclone
```
git clone https://github.com/nishiumidaina/gin-template.git
```
2.作業ディレクトリに入る
```
cd gin-template
```
3.コンテナを起動(ログが出ます)
```
docker-compose up
```
4.別ターミナルでコンテナに入る
```
docker-compose exec app sh
```
5.マイグレーションの作成(例：usersテーブル)
```
migrate create -ext sql -dir database/migrations -seq create_users
```
6.マイグレーションの実行(例：usersテーブル)
```
migrate -database="mysql://root:root@tcp(host.docker.internal:3306)/my_testdb?multiStatements=true" -path=database/migrations/ up 1
```
