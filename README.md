# pgroll-sample-app
Pgrollを動かすためのサンプルアプリケーション

## 利用方法
1. pgrollをインストールする
```sh
curl -LO https://github.com/xataio/pgroll/releases/download/v0.4.1/pgroll.linux.amd64
mv pgroll.linux.amd64 pgroll
chmod +x pgroll
```
2. Docker buildする
``` sh
docker build -t pgroll-sample-app:v1 ./app
```

3. 環境を初期化する
``` sh
docker compose down
docker compose up -d
./pgroll init --postgres-url postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
./pgroll start ddl/users.json --postgres-url postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable --complete
```

4. 問い合せを実行する
``` sh
# ユーザーを登録するAPI(データはランダム)
curl -X POST http://localhost:8082/users/new
# 登録したユーザーを検索するAPI
curl -X GET http://localhost:8082/users/${number}
```
