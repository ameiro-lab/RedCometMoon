package main

import (
	"net/http"

	"github.com/ameiro-lab/RedCometMoon/internal/moon"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 簡易疎通確認
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// 月の可視判定
	r.GET("/check-moon", func(c *gin.Context) {

		// パラメーター設定
		lat := 35.0
		lon := 139.0
		date := "2025-12-05" // テスト用で固定

		// 月が見えるか判定メソッドを呼ぶ
		result := moon.IsMoonVisibleForDay(lat, lon, date)
		if result.Error != nil {
			// クライアントに返す
			c.JSON(http.StatusInternalServerError, gin.H{"error": "500エラーが発生しました"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"visible": result.Visible,
			"azimuth": result.Azimuth,
		})
	})

	r.Run(":8080")
}
