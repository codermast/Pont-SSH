package entity

type LogInfo struct {
	Id       string `json:"id"`
	ServerId string `json:"server_id"`
	ExecDate string `json:"exec_date"`
	Command  string `json:"command"`
	Time     string `json:"time"`
}
