# Phase 2: 日付と判定ロジックを作る

## 目的

GitHub API を呼ばなくても作れる、純粋なロジック部分を先に完成させます。

この Phase が強いと、後の実装がかなり楽になります。

## この Phase でやること

1. timezone から「今日」を求める
2. contribution 数の判定ロジックを作る
3. 通知レベルを決める
4. 単体テストを書く

## 作業内容

### 1. 日付処理を切り出す

対象:

- `internal/date/`

やること:

- IANA timezone 文字列を受け取る
- その timezone での現在日時を求める
- 「今日の日付」を `YYYY-MM-DD` など扱いやすい形で返す

理由:

- GitHub の contribution 判定は日付のズレに弱いです
- Asia/Tokyo と UTC の差を早い段階で意識した方が安全です

### 2. contribution 判定を作る

対象:

- `internal/judgment/`

やること:

- 今日の contribution 数を受け取る
- 最低必要数 `minimum-contributions` を受け取る
- 条件を満たしたか返す

例:

- `0 < 1` なら未達
- `1 >= 1` なら達成
- `2 >= 1` なら達成

### 3. 通知レベルを定義する

設計書にある通知レベルを、まずはシンプルに表現します。

候補:

- `none`
- `gentle`
- `warning`
- `final`

最初のおすすめ:

- MVP では `none` と `gentle` だけでもよい
- 後から `warning` と `final` を追加する

### 4. 出力用の結果構造体を作る

やること:

- 判定結果を struct にまとめる
- 例として次のような情報を持たせる

候補:

- `Date`
- `ContributionCount`
- `MinimumContributions`
- `ThresholdMet`
- `NotificationLevel`

理由:

- `main.go` で扱いやすくなる
- 後で GitHub Action の outputs にもつなげやすいです

### 5. 単体テストを書く

特にテストしたいもの:

- timezone が正しく解釈できるか
- contribution 数の境界値
- `0`, `1`, `minimum - 1`, `minimum`, `minimum + 1`

## この Phase で作るとよいもの

- `internal/date/date.go`
- `internal/date/date_test.go`
- `internal/judgment/judgment.go`
- `internal/judgment/judgment_test.go`

## 完了条件

- timezone を入力して今日の日付を求められる
- contribution 数と閾値から達成判定を返せる
- 単体テストが通る

## 次に進む基準

- GitHub API の値がまだなくても、仮の contribution 数で動作確認できる
- ロジックが `main.go` に直接書かれていない

## 詰まりやすいポイント

- `time.LoadLocation` の使い方
- `time.Now()` を直接多用してテストしにくくすること
- 判定結果を `bool` だけで済ませて後で情報不足になること
