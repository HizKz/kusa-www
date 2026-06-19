# 草www

> 今日、草なくて草。

`kusa-www` は、GitHub のその日のコントリビューションがまだ 0 件だったときに、夜に向かって少しずつ圧を強めながらリマインドを送る Public GitHub Action を目指すプロジェクトです。  
狙いは「空コミットで草を生やす」ことではなく、「README を 1 行直す」「テストを 1 件足す」くらいの意味ある小さな一歩を促すことです。

## 最終的な形

最終的には、利用者が自分の Public Repository に workflow を 1 つ置き、Webhook を GitHub Secrets に登録するだけで使える GitHub Action として公開します。

想定している利用フローは次の通りです。

1. `schedule` または `workflow_dispatch` で Action を実行する
2. GitHub GraphQL API から対象ユーザーの当日 contribution 数を取得する
3. 設定した最低件数に届いていなければ通知レベルを決める
4. Discord Webhook に通知する
5. 判定結果を GitHub Actions outputs として返す

公開時の利用イメージは次の形です。

```yaml
uses: hizume0308/kusa-www@v1
with:
  github-token: ${{ github.token }}
  username: ${{ github.repository_owner }}
  timezone: Asia/Tokyo
  minimum-contributions: "1"
  notification-provider: discord
  webhook-url: ${{ secrets.DISCORD_WEBHOOK_URL }}
  tone: kusa
  dry-run: "false"
```

## 目指している体験

- GitHub の状態と連動した「今日まだ草がない」を知らせる
- Duolingo のストリーク通知のように、時間が遅くなるほど少し圧が強くなる
- 中央サーバーや中央 DB を持たず、利用者の Repository と Secrets だけで完結する
- 強い権限の Token を前提にせず、まずは Public な contribution 確認を中心にする
- 通知文はユーモアを持たせつつ、意味のある小さな作業を促す

## MVP で入れたいもの

- 当日 contribution 数の取得
- IANA タイムゾーン指定
- 最低 contribution 件数の指定
- 基準未満のときだけ通知
- Discord Webhook 通知
- 通知 tone の切り替え
- 時刻に応じた通知レベル変更
- `dry-run` 実行
- `workflow_dispatch` での手動確認
- GitHub Actions outputs の返却

## MVP の対象外

- Web 管理画面
- GitHub OAuth ログイン
- スマホアプリ
- 友達ランキングやチームランキング
- 中央 DB
- AI による通知文生成
- 自動コミット
- 自動 Issue 作成
- GitHub App 版

## 想定構成

```text
利用者の Public Repository
        │
        │ schedule / workflow_dispatch
        ▼
GitHub Actions Runner
        │
        ├── GitHub GraphQL API で contribution 数を取得
        └── 草がなければ Discord Webhook へ通知
```

## 現在の状態

このリポジトリはまだ開発途中です。現在は Go で CLI / Action の土台を段階的に組み立てている段階で、最終公開形の仕様は主に設計書とフェーズ文書にあります。

- 設計書: [docs/kusa-www-design.md](/Users/apple/Documents/solo-projects/kusa-www/docs/kusa-www-design.md)
- 実装フェーズ: [docs/phases/README.md](/Users/apple/Documents/solo-projects/kusa-www/docs/phases/README.md)

README は「このプロジェクトを最終的に何へ持っていくか」を先に示すためのものとして整理しています。実装が進んだら、セットアップ手順、実際の `action.yml`、サンプル workflow、CI 方針をここに追加していきます。
