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

![mission_service - public](https://user-images.githubusercontent.com/62779514/230784240-aef48956-e008-4fa7-99e4-9d57781a9780.png)

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

# 工夫した点

### ソース

- 責務分けを意識したディレクトリ構成で実装しました。
  - router
  - controller
  - usecase
  - model
  - repository
  - infrastructure
- DIを使用し単体テストがしやすい実装、他レイヤーの変更の影響を受けにくい実装をしました。
- ビジネスロジックを記述するusecaseファイルが肥大しすぎないように、目的ごとでファイル分けしました。
  - usecase/event/event.go(リクエストが来るイベント毎のメソッドを用意)
  - usecase/mission/daily(デイリーミッション)
  - usecase/mission/normal(全期間ミッション)
  - usecase/mission/release(ミッション解放)
  - usecase/mission/reward(ミッション報酬, 報酬に関わるミッション)
  - usecase/mission/weekly(週間ミッション)
- 現状のソースコードで以下のような条件のミッションが追加可能です。
  - 特定のモンスター討伐ミッション
    - モンスターBを倒す
    - モンスターAを2体倒す
  - 特定のモンスターレベルアップミッション
    - モンスターBのレベルが10になる
    - モンスターAのレベルが10になる
  - 一定レベル以上のモンスター獲得ミッション
    - レベル５以上のモンスターが3体
    - レベル10以上のモンスターが2体
  - 獲得コイン数ミッション
    - ３０００コイン貯まる
  - アイテム獲得ミッション
    - アイテムAを2つ所有

### DB設計
- ミッションの達成条件をソースコードではなくDBで持つ様にしました。
  - 今後ミッションが増えていった場合に、マスタデータ追加で対応できるように。
- ミッションの報酬テーブル(`mission_reward_items, mission_reward_coins`)のレコードをミッション(`missions`)に紐づける形にして、複数の報酬が存在するミッションの追加に対応しています。
- ミッションの達成条件テーブル(`monster_kill_missions`など)のレコードをミッション(`missions`)に紐づける形にして、複数の達成条件を持ったミッションの追加に対応しています。
- ミッション解放管理テーブル(`mission_releases`)のレコードをミッション(`missions`)に紐づける形にして、一つのミッション達成時に複数のミッションを解放できるように対応しています。
- 使わなくなった、ミッションは`missions.is_deleted`をtrueにすることで簡単に使用できなくしてあります。
- `missions.mission_type`を用意したことで、同じ達成条件でも全期間, デイリーミッションなどの種類が分けれます。

#  時間がなく実装できなかった点
- APIの認証が実装できていない。(Authorizationヘッダを利用してBearer token認証する予定でした。)
- request bodyのバリデーションが実装できていない。
- テストが全体的に書けていない。
- データの中身を確認できるエンドポイントを用意できていない。
