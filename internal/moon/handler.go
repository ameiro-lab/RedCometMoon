package moon

import (
	"log"
	"net/http"
	"os"

	"github.com/ameiro-lab/RedCometMoon/internal/line"
	"github.com/gin-gonic/gin"
)

// CheckMoonHandler は /check-moon エンドポイントのハンドラー
// lineClient を渡せるようにクロージャに変更
func CheckMoonHandler(lineClient *line.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		lat := 35.0
		lon := 139.0
		date := "2025-12-05"

		// 月が見えるか判定を実行
		result := IsMoonVisibleForDay(lat, lon, date)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "500エラーが発生しました"})
			return
		}

		// 月が見える場合は、
		if result.Visible {
			userID := os.Getenv("LINE_TARGET_USER_ID") // .env で管理
			// LINE に通知する
			err := lineClient.PushMessage(userID, "今日は月が見えます")
			if err != nil {
				log.Println("LINEメッセージ送信エラー:", err)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"visible": result.Visible,
			"azimuth": result.Azimuth,
		})
	}
}
