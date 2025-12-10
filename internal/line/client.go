package line

import "github.com/line/line-bot-sdk-go/linebot"

// Client はラッパー型
type Client struct {
	bot *linebot.Client
}
