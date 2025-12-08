* Goのインストール
```
$ go version
go version go1.25.2 darwin/amd64
```

* Ginの導入
```
go get github.com/gin-gonic/gin
```

* ディレクトリ構成
```
RedCometMoon/
├ cmd/
│   └ server/
│       └ main.go                // エントリーポイント
│
├ internal/
│   ├ moon/
│   │   ├ checker.go             // 月が見えるか判定
│   │   └ types.go               // もし構造体が増えるなら
│   │
│   ├ line/
│   │   ├ client.go              // LINE Bot の初期化
│   │   └ notifier.go            // 通知まわり
│   │
│   └ server/
│       ├ router.go              // Gin のルーティング
│       ├ handler.go             // ルートの実処理
│       └ middleware.go          // ミドルウェア（必要なら）
│
├ pkg/                           // 外部公開したい便利パッケージ（必要なら）
│
├ go.mod
└ go.sum
```