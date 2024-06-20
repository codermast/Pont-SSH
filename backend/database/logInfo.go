package database

import (
	"PontSsh/backend/entity"
	"encoding/json"
	"log"
)

// GetLogInfoList 查询日志列表
func GetLogInfoList(serverId string) []entity.LogInfo {
	sql := "select * from logs"

	// 查询指定服务器日志，则拼接条件
	if serverId != "all" {
		sql += " where server_id = ?"
	}

	rows, err := Query(sql)

	if err != nil {
		log.Printf("日志查询SQL执行异常：%v", err)
	}

	rowsJson, err := json.Marshal(rows)

	if err != nil {
		log.Printf("JSON 转换异常：%s", err)
	}

	log.Printf("Log日志数据%s", string(rowsJson))

	return nil
}

// SaveLogInfo 添加日志信息
func SaveLogInfo(logInfo entity.LogInfo) error {
	sql := "INSERT INFO main.logs (id,server_id,command,exec_date,time) VALUES(?,?,?,?,?)"

	err := Execute(sql, logInfo.Id, logInfo.ServerId, logInfo.Command, logInfo.ExecDate, logInfo.Time)

	return err
}
