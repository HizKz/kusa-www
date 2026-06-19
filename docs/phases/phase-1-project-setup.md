# Phase 1: Go プロジェクトを立ち上げる

## 目的

Go で実装を始めるための最小構成を作ります。

この Phase では、まだ GitHub API も Discord 通知も扱いません。
まずは `go run` でアプリが動く状態にすることが目標です。

## この Phase でやること

1. Go モジュールを初期化する
2. エントリーポイントを作る
3. 最小ディレクトリ構成を作る
4. 設定値を受け取る方法を決める
5. ローカル実行できることを確認する

## 作業内容

### 1. Go モジュールを初期化する

やること:

- `go mod init <module-path>` を実行する
- `go.mod` が生成されることを確認する

補足:

- GitHub 公開を想定するなら `github.com/<your-name>/kusa-www` 形式にすると自然です
- モジュール名は後で変更できますが、早めに決めた方が混乱しません

### 2. `main.go` を作る

やること:

- `cmd/kusa-www/main.go` を作る
- ひとまず `hello kusa-www` のような文字列を出す

この段階の狙い:

- Go の実行経路を理解する
- まず 1 本の実行可能プログラムを持つ

### 3. ディレクトリ構成を作る

やること:

- `cmd/kusa-www/`
- `internal/config/`
- `internal/date/`
- `internal/judgment/`
- `internal/message/`
- `internal/github/`
- `internal/notifier/`

補足:

- 最初は空ディレクトリでも構いません
- 後からファイルを増やす場所を先に決めておくと迷いにくいです

### 4. 設定値の持ち方を決める

最初に扱う設定値:

- `username`
- `timezone`
- `minimum-contributions`
- `dry-run`

おすすめ:

- 最初は環境変数でもフラグでもよい
- Go 初心者なら最初は固定値でも構いません
- ただし、値を `main.go` にベタ書きし続けないこと

### 5. 実行確認

やること:

- `go run ./cmd/kusa-www` で起動する
- エラーなく終了することを確認する

## この Phase で作るとよいもの

- `go.mod`
- `cmd/kusa-www/main.go`
- 必要なら `internal/config/config.go`

## 完了条件

- `go run ./cmd/kusa-www` が成功する
- リポジトリの基本ディレクトリ構成ができている
- 次の Phase でロジックを追加できる土台がある

## 次に進む基準

- Go ファイルの追加場所に迷わない
- `main.go` から別 package を呼ぶイメージが持てる

## 詰まりやすいポイント

- `package main` とそれ以外の package の違い
- `cmd/` と `internal/` の役割
- `go run .` と `go run ./cmd/kusa-www` の違い

## 学習メモ

この Phase では次の理解ができれば十分です。

- Go は package 単位でコードを分ける
- `main` package が実行開始地点になる
- `internal` はアプリ内部のコード置き場として使う
