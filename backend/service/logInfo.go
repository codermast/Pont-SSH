package service

import (
	"PontSsh/backend/database"
	"PontSsh/backend/entity"
	"PontSsh/backend/utils"
	"context"
	"log"
)

// LogInfo 连接结构体
type LogInfo struct {
	ctx context.Context
}

// NewLogInfo 构建 LogInfo 对象
func NewLogInfo() *LogInfo {
	return &LogInfo{}
}

// GetLogInfoList 获取日志信息
func (l *LogInfo) GetLogInfoList(serverId string) entity.Result {
	logs := database.GetLogInfoList(serverId)

	return utils.SuccessData(logs)
}

// SaveLogInfo 添加日志信息
func (l *LogInfo) SaveLogInfo(logInfo entity.LogInfo) entity.Result {
	err := database.SaveLogInfo(logInfo)

	if err != nil {
		log.Printf("日志保存异常：%v", err)
	}

	return utils.SuccessData(logInfo)
}
