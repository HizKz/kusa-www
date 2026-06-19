# Phase 6: GitHub Action と公開品質に仕上げる

## 目的

ローカルで動く CLI を GitHub Action として公開可能な形に仕上げます。

この Phase は「動くものを作る」より、「他人が安全に使えるようにする」工程です。

## この Phase でやること

1. `action.yml` を作る
2. Dockerfile を作る
3. outputs を整える
4. README を整備する
5. テストと CI を足す
6. リリース準備をする

## 作業内容

### 1. `action.yml` を作る

設計書の action interface をベースに作ります。

最低限入れたい inputs:

- `github-token`
- `username`
- `timezone`
- `minimum-contributions`
- `notification-provider`
- `webhook-url`
- `tone`
- `dry-run`

最初は省略してよいもの:

- `message-template`
- `mention`
- `suggest-small-task`

理由:

- 最初から全入力を実装しなくても MVP は成立するためです

### 2. Dockerfile を作る

やること:

- Go バイナリを build する
- 実行専用の軽いイメージへ載せる

ポイント:

- Multi-stage build にする
- 余計なツールを本番イメージに残さない

### 3. GitHub Actions outputs を返す

返したい値:

- `date`
- `contribution-count`
- `minimum-contributions`
- `threshold-met`
- `notified`
- `notification-level`
- `result`

注意:

- 出力形式を安定させる
- ログと output を混同しない

### 4. Workflow 例を書く

やること:

- README に最小 workflow を載せる
- `workflow_dispatch` 例を入れる
- `dry_run` 例を入れる

狙い:

- 他人がコピペしてすぐ試せる状態にする

### 5. テストを拡充する

最低限ほしいテスト:

- date
- judgment
- message
- config
- notifier の HTTP テスト

余力があれば:

- integration test
- action 実行確認

### 6. CI を入れる

最低限の CI:

- `go test ./...`
- `go vet ./...`
- build 確認

### 7. セキュリティ面を整える

やること:

- secrets をログに出さない
- `permissions: contents: read` の最小権限を使う
- Action の利用例では SHA pin の説明も書く

### 8. バージョンを切る

やること:

- `v0.x` か `v1.0.0` のタグ戦略を決める
- Marketplace 公開を見据えて README と説明文を整える

## この Phase で作るとよいもの

- `action.yml`
- `Dockerfile`
- `.github/workflows/ci.yml`
- README の導入手順

## 完了条件

- GitHub Action として実行できる
- README を見れば他人が導入できる
- 最低限のテストと CI がある
- 公開しても困らない品質になっている

## 次に考えること

この Phase 後に検討するとよいもの:

- 重複通知対策
- 圧レベルの追加
- small task suggestion
- GitHub Marketplace 最適化
- GitHub App 版

## 詰まりやすいポイント

- CLI 前提の実装を Action に無理やり載せること
- inputs と環境変数の責務が曖昧になること
- README を後回しにして利用方法が分かりにくくなること
