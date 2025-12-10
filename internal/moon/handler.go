package moon

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckMoonHandler は /check-moon エンドポイントのハンドラー
func CheckMoonHandler(c *gin.Context) {
	lat := 35.0
	lon := 139.0
	date := "2025-12-05"

	// 月が見えるか判定を実行
	result := IsMoonVisibleForDay(lat, lon, date)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "500エラーが発生しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"visible": result.Visible,
		"azimuth": result.Azimuth,
	})
}
