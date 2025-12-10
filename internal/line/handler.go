package line

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

// LINE Bot クライアントを作成
func CreateClient(channelSecret, channelToken string) (*Client, error) {
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return nil, err
	}
	return &Client{bot: bot}, nil // *Client型を返す
}

// 指定ユーザーにメッセージを送信
func (c *Client) PushMessage(userID, message string) error {
	_, err := c.bot.PushMessage(userID, linebot.NewTextMessage(message)).Do()
	if err != nil {
		log.Printf("LINE送信エラー: %v\n", err)
		return err
	}
	return nil
}

// 受け取ったLINEイベントをGinサーバーにログ出力する関数
func HandleEvents(req *http.Request, client *Client) {
	events, err := client.bot.ParseRequest(req)
	if err != nil {
		log.Println("Webhook Parse エラー:", err)
		return
	}

	for _, ev := range events {
		switch ev.Type {
		case linebot.EventTypeMessage:
			log.Println("メッセージイベント:", ev.Source.UserID)
			switch msg := ev.Message.(type) {
			case *linebot.TextMessage:
				log.Println("テキスト:", msg.Text)
			case *linebot.ImageMessage:
				log.Println("写真:", msg.ID)
			case *linebot.StickerMessage:
				log.Println("スタンプ:", msg.PackageID, msg.StickerID)
			}
		case linebot.EventTypeFollow:
			log.Println("友だち追加:", ev.Source.UserID)
		case linebot.EventTypeUnfollow:
			log.Println("友だち解除:", ev.Source.UserID)
		case linebot.EventTypeJoin:
			log.Println("Botが参加:", ev.Source.GroupID, ev.Source.RoomID)
		case linebot.EventTypeLeave:
			log.Println("Botが退出:", ev.Source.GroupID, ev.Source.RoomID)
		case linebot.EventTypePostback:
			log.Println("ポストバック:", ev.Postback.Data)
			if ev.Postback.Params != nil {
				log.Println("パラメータ:", ev.Postback.Params)
			}
		default:
			log.Println("未対応イベント:", ev.Type)
		}
	}
}
