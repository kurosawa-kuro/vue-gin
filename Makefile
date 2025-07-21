.PHONY: setup dev build clean frontend backend help

# 初期セットアップ
setup:
	@echo "Setting up frontend dependencies..."
	cd src/frontend && npm install
	@echo "Setting up backend dependencies..."
	cd src/backend && go mod tidy
	@echo "Setup complete!"

# 開発サーバー起動（並行実行）
dev:
	@echo "Starting development servers..."
	@trap 'kill %1 %2' INT; \
	cd src/backend && go run cmd/main.go & \
	cd src/frontend && npm run dev & \
	wait

# フロントエンドのみ起動
frontend:
	cd src/frontend && npm run dev

# バックエンドのみ起動  
backend:
	cd src/backend && go run cmd/main.go

# ビルド
build:
	@echo "Building frontend..."
	cd src/frontend && npm run build
	@echo "Building backend..."
	cd src/backend && go build -o bin/server cmd/main.go
	@echo "Build complete!"

# クリーンアップ
clean:
	@echo "Cleaning up..."
	cd src/frontend && rm -rf dist node_modules
	cd src/backend && rm -rf bin data.db
	@echo "Cleanup complete!"

# ヘルプ
help:
	@echo "Available commands:"
	@echo "  setup    - Install dependencies"
	@echo "  dev      - Start development servers"
	@echo "  frontend - Start frontend only"
	@echo "  backend  - Start backend only"
	@echo "  build    - Build both frontend and backend"
	@echo "  clean    - Clean build artifacts"
	@echo "  help     - Show this help message"
