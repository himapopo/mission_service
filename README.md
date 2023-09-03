# ミッションサービス
- モンスター育成・バトル系のゲームのミッション機能の実装

# 使用技術
- Go1.20
- PostgreSQL

### ライブラリ
- gin(FW)
- sql-migrate(migration tool)
- sqlboiler(ORM)
- air(hot reload)

# 実行方法
1. git cloneして移動
```sh
git clone git@github.com:himapopo/mission_service.git
cd mission_service
```
2. .envファイルを作成
```sh
make cp-env
```
3. build 
```sh
make build
```
4. API, DB立ち上げ
```sh
make up
```
5. migration
```sh
make migration/up
```
6. テストデータ作成
```sh
make test-data-init
```

# ER図

![mission_service - public1](https://user-images.githubusercontent.com/62779514/230787447-c20c1050-e442-4248-8e9e-a1ab23cbb456.png)

# ストーリー
※データを確認できるエンドポイントを用意できていないので、GUIツールなどからデータの確認をお願いいたします。

### 土曜の午後１０時

#### 1. ログインする

endpoint
```text
POST: http://localhost:8080/login
```

RequestHeaders
```sh
Content-Type: application/json
```

body
```json
{
    "user_id": 1,
    "requested_at": "2023-04-08T13:00:00Z"
}
```

#### 2. モンスターAが任意のモンスターを倒す

endpoint
```text
POST: http://localhost:8080/monster_kill
```

RequestHeaders
```sh
Content-Type: application/json
```

body
```json
{
    "user_id": 1,
    "user_monster_id": 1,
    "opponent_monster_id": 2,
    "requested_at": "2023-04-08T13:00:00Z"
}
```

#### 3. モンスターAのレベルが２上がる

endpoint
```text
POST: http://localhost:8080/monster_level_up
```

RequestHeaders
```sh
Content-Type: application/json
```

body
```json
{
    "user_id": 1,
    "user_monster_id": 1,
    "amount": 2,
    "requested_at": "2023-04-08T13:00:00Z"
}
```

### 日曜の午後１０時

#### 1. ログインする

endpoint
```text
POST: http://localhost:8080/login
```

RequestHeaders
```sh
Content-Type: application/json
```

body
```json
{
    "user_id": 1,
    "requested_at": "2023-04-09T13:00:00Z"
}
```

#### 2. モンスターBが特定のモンスターを倒す

endpoint
```text
POST: http://localhost:8080/monster_kill
```

RequestHeaders
```sh
Content-Type: application/json
```

body
```json
{
    "user_id": 1,
    "user_monster_id": 2,
    "opponent_monster_id": 1,
    "requested_at": "2023-04-09T13:00:00Z"
}
```

#### 3. モンスターBがレベル１上がる

endpoint
```text
POST: http://localhost:8080/monster_level_up
```

RequestHeaders
```sh
Content-Type: application/json
```

body
```json
{
    "user_id": 1,
    "user_monster_id": 2,
    "amount": 1,
    "requested_at": "2023-04-09T13:00:00Z"
}
```
