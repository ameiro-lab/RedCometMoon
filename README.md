* Goのインストール
```
$ go version
go version go1.25.2 darwin/amd64
```

* Ginの導入
```
$ go get github.com/gin-gonic/gin
```

* cronパッケージの利用
```
$ go get github.com/robfig/cron/v3
```

* SDKを取得
```
$ go get github.com/line/line-bot-sdk-go/linebot
```

* ngrokの導入
```
$ brew install --cask ngrok
```
アカウント登録とトークンの設定が必要。https://dashboard.ngrok.com/ できたらコマンドを実行
```
$ ngrok http 8080
Forwarding    https://stereotactically-windrode-jewel.ngrok-free.dev -> http://localhost:8080
```
https://以降のURL + `/callback` を LINE Developers の `Webhook URL` に設定する。


* ディレクトリ構成
```
RedCometMoon/
├ cmd/
│   └ server/
│       └ main.go                // エントリーポイント
│
├ internal/　← internalキーワードは import させない“非公開パッケージ”
│   ├ moon/
│   │   ├ checker.go             // 月が見えるか判定
│   │   └ types.go               // もし構造体が増えるなら
│   │
│   ├ line/
│   │   ├ client.go              // LINE Bot の初期化
│   │   └ notifier.go            // 通知まわり
│   │
│   ├ cron/
│   │   └ scheduler.go           // 定期実行処理をまとめる
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

* 機能要件ざっと
1. 月の可視判定ロジック
   * 与えられた日時・位置情報（緯度/経度）から「月が自分の部屋の窓から見えるか」 を判定する。
2. LINE への通知機能
   * LINE Bot SDK を用いて LINE の指定ユーザーへメッセージを送信する。
3. LINE Webhook 受信機能（TBL）
   * ユーザー操作を受けたい場合は必要。【拡張予定】
4. 月の可視チェックのスケジューラー
   * 一定間隔（例：毎日 18:00）に自動で月の可視判定を行う。定期実行の処理。
5. 設定（Config）読み込み
   * .env から環境変数をロード（godotenv）LINE のトークンや緯度経度などを安全に管理。
6. API（簡易ステータスチェック）
   * デバック用のAPI。Gin で “現在の設定 / 次のチェック時間 / 前回の月判定結果” を返す。

* 非機能要件ざっと
1. ログ管理
   * Go の標準 log or zap でアプリの動作ログ（起動・判定結果・通知結果）を記録。
2. エラーハンドリング
   * APIの失敗、ネットワークエラー、LINE通知のエラーを共通化。
3. 設定のセキュリティ
   * LINE トークンなど機密情報が漏れないようにする。
   * .env を gitignore、環境変数で管理、リポジトリに埋め込まない
4. 内部モジュールの分離
   * internal で他のプロジェクトから import できない構造にする。