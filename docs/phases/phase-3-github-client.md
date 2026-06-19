# Phase 3: GitHub の contribution 数を取得する

## 目的

GitHub GraphQL API を使って、対象ユーザーの「今日の contribution 数」を取得します。

ここで初めて外部 API が登場します。

## この Phase でやること

1. GitHub API クライアントの責務を決める
2. GraphQL Query を作る
3. API レスポンスを Go の struct にマッピングする
4. 今日の contribution 数を取り出す
5. エラー処理を整える

## 作業内容

### 1. インターフェースを定義する

対象:

- `internal/github/`

やること:

- 設計書の `ContributionClient` に近い interface を作る

目的:

- 本物の API 呼び出しと、テスト用の偽物を差し替えやすくする

### 2. GraphQL Query を用意する

設計書の GitHub API 設計を元に、最低限必要なデータだけ取得します。

最初に狙うこと:

- 指定ユーザー
- 指定日付範囲
- contribution calendar から当日の数を得る

ポイント:

- 最初から全部の項目を取らない
- 今日の数を出すための最小レスポンスに絞る

### 3. HTTP リクエストを送る

やること:

- `https://api.github.com/graphql` に POST する
- Authorization ヘッダーに token を付ける
- `Content-Type: application/json` を付ける

設定値:

- `github-token`
- `username`

### 4. レスポンス構造体を作る

やること:

- GraphQL の JSON に対応する struct を定義する
- 必要な field だけを持たせる

注意:

- 最初は deeply nested な struct になっても構いません
- 後で整理すればよいです

### 5. 今日の contribution 数を抜き出す

やること:

- 指定した local date に一致する contribution を探す
- 数字を返す

補足:

- GitHub の返す日付と local timezone の扱いを混同しないこと
- まずは `today` の count を 1 つ返せれば十分です

### 6. エラー処理を入れる

最低限見るべき失敗:

- token が空
- HTTP エラー
- GraphQL errors が返る
- 指定日の contribution が見つからない

## この Phase で作るとよいもの

- `internal/github/client.go`
- `internal/github/client_test.go`
- 必要なら `internal/github/types.go`

## 動作確認方法

最初の確認方法:

- `main.go` で token と username を渡す
- 取得した contribution 数を標準出力に表示する

成功の目安:

- 実際の GitHub アカウントに対して数字が出る

## 完了条件

- GitHub GraphQL API から対象ユーザーの contribution 数を取得できる
- 失敗時に原因が分かるエラーを返せる

## 次に進む基準

- 仮の値ではなく、実データを使って Phase 2 の判定に流せる

## 詰まりやすいポイント

- GraphQL の JSON struct 定義
- token 未設定時の失敗
- timezone と GitHub 側の日付の対応
- private contributions を MVP でどう扱うか

## 実装の割り切り

初心者なら、最初は次の割り切りで十分です。

- 認証は `GITHUB_TOKEN` だけ
- GitHub user 1人だけ対応
- 今日の contribution 数だけ返す
- retry や複雑な rate limit 対応は後回し
