# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

MySQLデータベースを使用したブログ/記事プラットフォームのGo REST APIサーバーです。HTTPハンドラ、ビジネスロジック、データアクセスを明確に分離した3層アーキテクチャを採用しています。

## アーキテクチャ

アプリケーションは以下の層状アーキテクチャパターンに従っています：

**レイヤーフロー**: HTTPリクエスト → Controllers → Services → Repositories → MySQLデータベース

- **`main.go`**: エントリーポイント - データベース接続を確立し（`.env`で設定）、ポート8080でHTTPサーバーを起動
- **`api/router.go`**: gorilla/muxを使用したルート定義、コントローラとサービス層を連携
- **`controllers/`**: HTTPハンドラ - JSONリクエストをデコード、サービスを呼び出し、JSONレスポンスをエンコード
- **`controllers/services/services.go`**: コントローラで使用する依存性注入のためのインターフェース定義（`ArticleServicer`, `CommentServicer`）
- **`services/`**: ビジネスロジック層 - サービスインターフェースを実装し、リポジトリ呼び出しを調整
- **`repositories/`**: データアクセス層 - すべてのSQLクエリとデータベース操作を含む
- **`models/models.go`**: 全レイヤーで共有されるデータモデル（`Article`, `Comment`）

### 主要なアーキテクチャパターン

1. **インターフェースベースの依存性注入**: コントローラは具体的な実装ではなく、サービスインターフェース（`controllers/services/`で定義）に依存。テスタビリティと疎結合を実現。

2. **トランザクション管理**: データベーストランザクションはリポジトリ層で処理（`repositories/articles.go:83-114`の`UpdateNiceNum`を参照）

3. **サービスの組み合わせ**: サービス層（`MyAppService`）は複数のリポジトリからデータを集約（例：`GetArticleService`は記事詳細とコメントリストを組み合わせる - `services/article_service.go:8-22`）

## データベースセットアップ

プロジェクトはDockerで実行されるMySQL 8.0を使用：

```bash
# MySQLコンテナを起動
docker compose up -d

# テーブルを作成
mysql -h 127.0.0.1 -u docker sampledb --password=docker < createTable.sql

# テストデータを挿入（オプション）
mysql -h 127.0.0.1 -u docker sampledb --password=docker < insertData.sql
```

データベース認証情報は`.env`で設定：
- ユーザー: `docker`
- パスワード: `docker`
- データベース: `sampledb`
- ホスト: `127.0.0.1:3306`

## アプリケーションの実行

```bash
# サーバーを起動（MySQLが起動している必要があります）
go run main.go

# サーバーは http://localhost:8080 で起動します
```

## APIエンドポイント

- `POST /article` - 新しい記事を作成
- `GET /article/list?page=1` - 記事一覧を取得（ページネーション、1ページ5件）
- `GET /article/{id}` - 記事詳細をコメント付きで取得
- `POST /article/nice` - 記事の「いいね」カウントをインクリメント
- `POST /comment` - 記事にコメントを追加

## テスト

テストは`repositories/`パッケージにあり、実行中のMySQLデータベースが必要です：

```bash
# repositoriesパッケージのすべてのテストを実行
go test ./repositories/...

# 特定のテストを実行
go test ./repositories -run TestSelectArticleDetail

# 詳細出力でテストを実行
go test -v ./repositories/...
```

### テストアーキテクチャ

- **`repositories/main_test.go`**: `TestMain`を使用したテストのセットアップ/クリーンアップ - DB接続を確立し、テスト実行前に`testdata/`ディレクトリのSQLファイルからテストデータをロード、テスト後にクリーンアップ
- **テストデータ管理**: `mysql`コマンドラインツールを使用してSQLファイルを実行し、セットアップ/クリーンアップを行う
- **テスト構造**: テーブル駆動テストパターンを使用（`repositories/articles_test.go:25-68`の`TestSelectArticleDetail`を参照）
- **クリーンアップパターン**: テストは`t.Cleanup()`を使用して、変更後にデータベースの状態を復元

**重要**: テストは直接のMySQL CLI アクセスとハードコードされた認証情報（`docker`/`docker`/`sampledb`）を前提としています

## モジュール情報

- モジュールパス: `github.com/tetsuya-stn/go-api-server-handson`
- Goバージョン: 1.22.1
- 依存関係:
  - `github.com/go-sql-driver/mysql` - MySQLドライバ
  - `github.com/gorilla/mux` - HTTPルーター

## 開発時の注意事項

- **ページネーション**: 記事一覧は`repositories/articles.go:12`で定義された定数`articleNumPerPage = 5`を使用
- **エラーハンドリング**: エラーは層を通じて伝播 - リポジトリがエラーを返し、サービスがそれを渡し、コントローラがHTTPステータスコードに変換
- **SQLインジェクション対策**: すべてのクエリはパラメータ化されたステートメント（`?`プレースホルダ）を使用
- **時刻の扱い**: 接続文字列の`parseTime=true`により、MySQLのDATETIMEフィールドがGoの`time.Time`にマッピングされる
