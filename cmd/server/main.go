package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/ameiro-lab/RedCometMoon/internal/line"
	"github.com/ameiro-lab/RedCometMoon/internal/moon"
	"github.com/gin-gonic/gin"
)

// .envを環境変数として設定する関数
func loadEnv(path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		t := s.Text()
		if !strings.Contains(t, "=") {
			continue
		}
		kv := strings.SplitN(t, "=", 2)
		os.Setenv(strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1]))
	}
}

func main() {
	// .envをロードする
	loadEnv(".env")
	// 環境変数から LINE Bot の秘密情報を取得
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	channelToken := os.Getenv("LINE_CHANNEL_TOKEN")
	// 秘密情報を使って LINE Bot SDK クライアントを初期化
	lineClient, err := line.CreateClient(channelSecret, channelToken)
	if err != nil {
		log.Fatalf("LINE 初期化エラー: %v", err)
	}

	// ---ginサーバー起動
	r := gin.Default()
	// サーバーが起動しているか確認
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// --- WebhookでLINEイベントを検知する
	r.POST("/callback", func(c *gin.Context) {
		// 処理関数を呼び出し
		line.HandleEvents(c.Request, lineClient)
		c.Status(200)
	})

	// --- 月が見えるか判定を実行
	r.GET("/check-moon", moon.CheckMoonHandler)

	r.Run(":8080")
}
