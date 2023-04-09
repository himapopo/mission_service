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
- 達成条件のテーブル毎で、未達成のミッションを全てチェックするようにしているので、以下の様な既存のミッションの上位版や対象が違うだけのミッションを、現状のソースコードのまま追加可能です。
  - 特定のモンスターを倒す(既存)
    - モンスターBを倒す
    - 特定のモンスターを2体倒す
  - モンスターAのレベルが５になる(既存)
    - モンスターBのレベルが10になる
    - モンスターAのレベルが10になる
  - レベル５以上のモンスターが２体(既存)
    - レベル５以上のモンスターが3体
    - レベル10以上のモンスターが2体
  - ２０００コイン貯まる(既存)
    - ３０００コイン貯まる
  - アイテムAを所有する(既存)
    - アイテムBを所有する
    - アイテムAを2つ所有する

### DB設計
- ミッションの達成条件をソースコードではなくDBで持つ様にしました。
  - 今後ミッションが増えていった場合に、マスタデータ追加で対応できるように。
- ミッションの報酬テーブル(`mission_reward_items, mission_reward_coins`)のレコードをミッション(`missions`)に紐づける形にして、複数の報酬が存在するミッションの追加に対応しています。
- ミッションの達成条件テーブル(`monster_kill_missions`など)のレコードをミッション(`missions`)に紐づける形にして、複数の達成条件を持ったミッションの追加に対応しています。（※複数条件対応時、ソースコードの修正も必要）
- ミッション解放管理テーブル(`mission_releases`)のレコードをミッション(`missions`)に紐づける形にして、一つのミッション達成時に複数のミッションを解放できるように対応しています。
- 不要になったミッションは`missions.is_deleted`をtrueにすることで簡単に無効化できます。
- `missions.mission_type`を用意したことで、同じ達成条件でも全期間, デイリーミッションなどの種類が分けれます。

#  時間がなく実装できなかった点
- APIの認証が実装できていない。(Authorizationヘッダを利用してBearer token認証する予定でした。)
- request bodyの値に対してバリデーションの実装ができていない。
- テストが全体的に書けていない。
- データの中身を確認できるエンドポイントを用意できていない。
