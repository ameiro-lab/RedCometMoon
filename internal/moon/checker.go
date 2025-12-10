package moon

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 月の可視判定結果
type CheckResult struct {
	Visible bool    // 月が見えるかどうか
	Azimuth float64 // 月の方角（°）
	Error   error   // API 呼び出しや計算でのエラー
}

// 月が見えるか判定
func IsMoonVisibleForDay(lat, lon float64, date string) CheckResult {

	// リクエスト URL を組み立てる
	url := fmt.Sprintf(
		"https://mgpn.org/api/moon/v2position.cgi?time=%sT18:00&lat=%f&lon=%f&loop=4&interval=60",
		date, lat, lon,
	)

	// まぢぽん製作所にAPIリクエスト
	resp, err := http.Get(url)
	if err != nil {
		return CheckResult{Visible: false, Error: fmt.Errorf("APIリクエストエラー: %w", err)}
	}
	defer resp.Body.Close() // HTTP 接続のクローズ処理

	// JSONデータを Go の構造体に変換
	var data struct {
		Result []struct {
			Time     string  `json:"time"`     // 計算時刻（ISO形式）
			Azimuth  float64 `json:"azimuth"`  // 月の方角
			Altitude float64 `json:"altitude"` // 月の高度
		} `json:"result"`
		Status int `json:"status"` // APIステータス
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return CheckResult{Visible: false, Error: fmt.Errorf("JSON デコードエラー: %w", err)}
	}

	// デバック用：レスポンスをログで確認
	fmt.Printf("デバッグ: %+v\n", data)

	// 判定ロジック：東北東 (67.5° ± 20°)
	center := 67.5
	tolerance := 20.0
	for _, pos := range data.Result {
		inDirection := abs(pos.Azimuth-center) <= tolerance || abs(pos.Azimuth-center) >= 360-tolerance
		if inDirection {
			return CheckResult{
				Visible: true,
				Azimuth: pos.Azimuth,
				Error:   nil,
			}
		}
	}

	// 条件を満たさなければ false
	return CheckResult{Visible: false}
}

// abs は float64 の絶対値
func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
