# gin-template
 Dockerで使用するGinのテンプレート。  
## 構成
[Gin](https://github.com/gin-gonic/gin)：フレームワーク  
[GORM](https://gorm.io/ja_JP/docs/index.html)：ORMライブラリ  
[golang-migrate](https://github.com/golang-migrate/migrate)：マイグレーションツール  
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
