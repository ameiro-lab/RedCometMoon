package cron

import (
	"log"
	"os"
	"time"

	"github.com/ameiro-lab/RedCometMoon/internal/line"
	"github.com/ameiro-lab/RedCometMoon/internal/moon"
	"github.com/robfig/cron/v3"
)

// StartScheduler は定期実行を開始する
func StartScheduler(lineClient *line.Client) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	c := cron.New(cron.WithLocation(loc))

	// .env から3つの部分を取得
	minute := os.Getenv("MOON_CHECK_MINUTE") // 例: "0"
	hour := os.Getenv("MOON_CHECK_HOUR")     // 例: "11"
	other := os.Getenv("MOON_CHECK_OTHER")   // 例: "* * *" （日 月 曜日）

	// cron 式を組み立て
	cronExpr := minute + " " + hour + " " + other

	// 毎日11:00に月判定
	c.AddFunc(cronExpr, func() {
		log.Println("月の可視チェック（cron）開始")
		visible, azimuth, err := moon.CheckMoonAndNotify(lineClient)
		if err != nil {
			log.Println("月判定エラー:", err)
		} else {
			log.Printf("定期チェック: visible=%v azimuth=%.2f\n", visible, azimuth)
		}
	})

	c.Start()
}
