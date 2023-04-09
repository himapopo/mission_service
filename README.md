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
1. git clone
```sh
git clone git@github.com:himapopo/mission_service.git
```
2. build 
```sh
make build
```
3. API, DB立ち上げ
```sh
make up
```
4. migration
```sh
make migration/up
```
5. テストデータ作成
```sh
make test-data-init
```

# ストーリー

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