package model

// System全体の状態を管理するモデルクラス
type SystemInfo struct {
	Status int `json: "status"`
	Message string `json: "message"`
	Time string `json: "time"`
}