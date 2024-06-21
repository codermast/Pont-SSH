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

	log.Println("sql:", sql)

	rows, err := Query(sql, serverId)

	if err != nil {
		log.Printf("日志查询SQL执行异常：%v", err)
	}

	// 日志数据
	logInfos := make([]entity.LogInfo, 0)

	for rows.Next() {
		var logInfo entity.LogInfo
		err := rows.Scan(&logInfo.Id, &logInfo.ServerId, &logInfo.Time, &logInfo.Command, &logInfo.ExecDate /* 指定字段 */)
		if err != nil {
			log.Printf("扫描行出错：%v", err)
			// 处理错误，例如跳过或者返回错误信息
			continue
		}
		logInfos = append(logInfos, logInfo)
	}

	rowsJson, err := json.Marshal(logInfos)

	if err != nil {
		log.Printf("JSON 转换异常：%s", err)
	}

	log.Printf("Log日志数据%s", string(rowsJson))

	return logInfos
}

// SaveLogInfo 添加日志信息
func SaveLogInfo(logInfo entity.LogInfo) error {
	sql := "INSERT INFO main.logs (id,server_id,command,exec_date,time) VALUES(?,?,?,?,?)"

	err := Execute(sql, logInfo.Id, logInfo.ServerId, logInfo.Command, logInfo.ExecDate, logInfo.Time)

	return err
}
