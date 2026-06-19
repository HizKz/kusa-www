# kusa-www Phases

`kusa-www` を Go で段階的に実装するための作業ガイドです。

Go 初心者を前提に、いきなり GitHub Action 全体を作るのではなく、ローカルで動く最小構成から少しずつ積み上げます。

## Phase 一覧

1. [Phase 1: Go プロジェクトを立ち上げる](./phase-1-project-setup.md)
2. [Phase 2: 日付と判定ロジックを作る](./phase-2-date-and-judgment.md)
3. [Phase 3: GitHub の contribution 数を取得する](./phase-3-github-client.md)
4. [Phase 4: メッセージ生成と通知を作る](./phase-4-message-and-notifier.md)
5. [Phase 5: アプリケーションとしてつなぎ込む](./phase-5-application-flow.md)
6. [Phase 6: GitHub Action と公開品質に仕上げる](./phase-6-action-and-release.md)

## 読み方

- 各 Phase には目的、やること、完了条件を書いています
- 上から順に進める前提です
- 途中で詰まったら、1つ前の Phase の完了条件を満たせているか確認してください

## 実装の基本方針

- 最初はローカル CLI として動かす
- 複雑な機能は後回しにして、1回動く最小機能を先に作る
- 純粋関数にしやすい部分から先に実装する
- テストしやすい構造を意識して `internal/` に分割する

## 想定ディレクトリ

```text
kusa-www/
  cmd/kusa-www/main.go
  internal/config/
  internal/date/
  internal/judgment/
  internal/message/
  internal/github/
  internal/notifier/
  docs/phases/
```
