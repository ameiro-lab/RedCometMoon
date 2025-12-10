package moon

import (
	"log"
	"net/http"
	"os"

	"github.com/ameiro-lab/RedCometMoon/internal/line"
	"github.com/gin-gonic/gin"
)

// "time"
// date := time.Now().Format("2006-01-02") // 今日の日付

// CheckMoonHandler は /check-moon エンドポイントのハンドラー
func CheckMoonHandler(lineClient *line.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		visible, azimuth, err := CheckMoonAndNotify(lineClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "500エラー"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"visible": visible, "azimuth": azimuth})
	}
}

// 月の可視判定とLINE通知を行う共通関数
func CheckMoonAndNotify(lineClient *line.Client) (visible bool, azimuth float64, err error) {
	lat := 35.0
	lon := 139.0
	date := "2025-12-05"

	// 月が見えるか判定を実行
	result := IsMoonVisibleForDay(lat, lon, date)
	if result.Error != nil { // これ何返してるのks
		return false, 0, result.Error
	}

	// 月が見える場合は、
	if result.Visible {
		userID := os.Getenv("LINE_TARGET_USER_ID") // .env で管理
		// LINE に通知する
		err := lineClient.PushMessage(userID, "月が見えます！")
		if err != nil {
			log.Println("LINEメッセージ送信エラー:", err)
		} else {
			log.Println("LINEメッセージ送信完了")
		}
	} else {
		// log.Println("月は見えません")
	}

	return result.Visible, result.Azimuth, nil
}
