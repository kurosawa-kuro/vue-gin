# Vue + Gin Full Stack Application

Vue 3 + Tailwind CSS フロントエンドと Go + Gin + GORM バックエンドのサンプルアプリケーション

## 🛠 技術スタック

### フロントエンド
- Vue 3 + Composition API
- Pinia (状態管理)
- Tailwind CSS
- Axios (HTTP クライアント)
- Zod (データ検証)

### バックエンド
- Go 1.20+
- Gin (Web フレームワーク)
- GORM (ORM)
- SQLite (データベース)
- CORS サポート

## 📋 前提条件

- Go 1.20+
- Node.js 16+
- npm または yarn

## 🚀 セットアップ手順

### 1. リポジトリクローン
```bash
git clone https://github.com/kurosawa-kuro/vue-gin.git
cd vue-gin
```

### 2. 依存関係インストール
```bash
# フロントエンド依存関係
cd src/frontend
npm install

# バックエンド依存関係
cd ../backend
go mod tidy
```

### 3. 環境変数設定
```bash
# src/backend/.env ファイル作成
DATABASE_DSN=data.db
PORT=3001
```

### 4. 開発サーバー起動

#### 方法1: Makefile使用（推奨）
```bash
# プロジェクトルートで
make dev
```

#### 方法2: 個別起動
```bash
# ターミナル1: バックエンド
cd src/backend
go run cmd/main.go

# ターミナル2: フロントエンド  
cd src/frontend
npm run dev
```

## 🌐 アクセス

- フロントエンド: http://localhost:5173
- バックエンドAPI: http://localhost:3001

## 📡 API エンドポイント

- `GET /users` - ユーザー一覧取得
- `GET /microposts` - 投稿一覧取得
- `GET /microposts/:id` - 個別投稿取得
- `POST /microposts` - 新規投稿作成

## 🗄 データモデル

### User
- ID (uint, primary key)
- Name (string)
- CreatedAt, UpdatedAt (time)
- Microposts (関連)

### Micropost
- ID (uint, primary key)
- Title (string)
- UserID (uint, foreign key)
- CreatedAt, UpdatedAt (time)
- User (関連)

## 🔧 開発コマンド

```bash
make setup    # 初期セットアップ
make dev      # 開発サーバー起動
make build    # ビルド
make clean    # クリーンアップ
```

## 🧪 動作確認

1. 両方のサーバーが起動していることを確認
2. ブラウザで http://localhost:5173 にアクセス
3. Dashboard、Users、Microposts ページが正常に表示されることを確認
4. 新規投稿作成機能をテスト

## 📝 開発ノート

- フロントエンドは Vue 3 + Composition API を使用
- バックエンドは Go + Gin + GORM で RESTful API を提供
- データベースは SQLite を使用（開発環境）
- CORS 設定により フロントエンド ⇄ バックエンド 間の通信を許可
