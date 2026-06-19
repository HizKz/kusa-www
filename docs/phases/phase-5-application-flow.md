# Phase 5: アプリケーションとしてつなぎ込む

## 目的

これまで作った部品をまとめて、`kusa-www` を 1 本の CLI アプリとして完成させます。

この Phase が終わると、ローカルで実用レベルの動作確認ができます。

## この Phase でやること

1. 設定読み込みを整理する
2. UseCase を作る
3. 各部品を `main.go` から組み立てる
4. 結果を標準出力へ出す
5. 統合的に動作確認する

## 作業内容

### 1. config をまとめる

対象:

- `internal/config/`

やること:

- 設定 struct を作る
- 必須入力をまとめる
- 環境変数やフラグから値を読む

まず必要な設定:

- `github-token`
- `username`
- `timezone`
- `minimum-contributions`
- `notification-provider`
- `webhook-url`
- `tone`
- `dry-run`

### 2. UseCase を作る

対象:

- `internal/usecase/` を追加してもよい

やること:

- アプリの主処理を 1 箇所にまとめる

流れ:

1. 設定を受け取る
2. 今日の日付を決める
3. GitHub から contribution 数を取る
4. 判定する
5. メッセージを作る
6. 必要なら通知する
7. 結果を返す

理由:

- `main.go` を薄く保てる
- テストしやすくなる

### 3. `main.go` を薄くする

`main.go` の役割は次に絞るのがおすすめです。

- config を読む
- 依存オブジェクトを作る
- usecase を実行する
- exit code を決める

### 4. 出力結果を整える

やること:

- 実行結果を読みやすく表示する
- 最低限、次の情報を出す

候補:

- date
- contribution count
- threshold met
- notified
- notification level

### 5. 統合的な確認をする

確認パターン:

- contribution が 0 の日
- contribution がある日
- `dry-run=true`
- webhook なしで stdout 通知
- token 不正

## この Phase で作るとよいもの

- `internal/config/config.go`
- `internal/config/config_test.go`
- `internal/usecase/check_grass.go`
- `internal/usecase/check_grass_test.go`

## 完了条件

- コマンド 1 つで全体の処理が動く
- 判定から通知までが一連でつながる
- `main.go` に細かいロジックがべた書きされていない

## 次に進む基準

- ローカル CLI としては一通り実用になる
- GitHub Action 化するだけの土台がある

## 詰まりやすいポイント

- `main.go` に全部書いてしまうこと
- config 読み込みと業務ロジックが混ざること
- 失敗時の終了コードやメッセージが曖昧になること
