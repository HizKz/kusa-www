# Phase 4: メッセージ生成と通知を作る

## 目的

判定結果から、ユーザーに見せるメッセージを作り、必要なら通知できる状態にします。

この Phase では、ロジックと通知先を分けて設計することが重要です。

## この Phase でやること

1. メッセージ生成を独立させる
2. Notifier interface を作る
3. 標準出力 notifier を作る
4. Discord notifier を作る
5. `dry-run` を扱う

## 作業内容

### 1. メッセージ生成を作る

対象:

- `internal/message/`

やること:

- 判定結果を受け取り、通知文を返す
- tone ごとの文面切り替えをできるようにする

最初に扱う tone:

- `kusa`

後から追加してよいもの:

- `gentle`
- `warning`
- `final`

ポイント:

- まずは文字列 1 本を返すだけでよい
- いきなり Discord embed 全体を組み立てなくてもよい

### 2. Notifier interface を作る

対象:

- `internal/notifier/`

やること:

- `Notify(...) error` のような interface を作る

理由:

- 通知先の差し替えがしやすい
- テストでモックしやすい

### 3. 標準出力 notifier を作る

やること:

- コンソールへメッセージを出す notifier を作る

狙い:

- Discord を使わなくても動作確認できる
- 開発初期の確認が速い

### 4. Discord notifier を作る

やること:

- webhook URL を受け取る
- JSON payload を POST する

最初はここまでで十分:

- `content` に本文を入れる
- 必要なら後から embeds を追加する

### 5. `dry-run` を入れる

やること:

- `dry-run=true` のときは Discord 送信しない
- 代わりに「送るはずだった内容」を表示する

理由:

- 本番 webhook への誤送信を防げる

### 6. 秘密情報をログに出さない

注意:

- webhook URL をログに出さない
- token をログに出さない
- エラー文に秘密情報が混ざる場合はマスクする

## この Phase で作るとよいもの

- `internal/message/message.go`
- `internal/message/message_test.go`
- `internal/notifier/notifier.go`
- `internal/notifier/stdout.go`
- `internal/notifier/discord.go`

## 完了条件

- 判定結果から通知文を作れる
- ローカルでは stdout 通知で確認できる
- Discord webhook に送信できる
- `dry-run` で送信抑止できる

## 次に進む基準

- API取得結果と判定結果を、通知という形で外へ出せる

## 詰まりやすいポイント

- message 生成と notifier の責務が混ざること
- Discord payload を最初から複雑にしすぎること
- `dry-run` なのに実際に送ってしまうこと
