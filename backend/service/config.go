package service

import (
	"PontSsh/backend/database"
	"PontSsh/backend/entity"
	"PontSsh/backend/utils"
	"context"
	"encoding/json"
	"log"
)

type Config struct {
	ctx context.Context
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Startup(ctx context.Context) {
	// 1. 加载配置文件
	database.LoadConfig()

	// 2. 加载配置启动器
	database.ConfigStartup()
}

func (c *Config) Shutdown(ctx context.Context) {
	// 1. 加载配置结束器
	database.ConfigShutdown()
}

// GetConfig 获取配置信息
func (c *Config) GetConfig() entity.Result {
	config, err := database.GetConfig()

	if err != nil {
		return utils.Error("获取配置信息失败！")
	}

	return utils.SuccessData(config)
}

// ResetConfig 重置配置信息
func (c *Config) ResetConfig() entity.Result {
	err := database.ResetConfig()
	if err != nil {
		return utils.Error("重置配置信息失败")
	}

	config, err := database.GetConfig()

	if err != nil {
		return utils.Error("获取配置信息失败！")
	}

	return utils.Success(config, "重置配置信息成功")
}

// SaveConfig 保存配置信息
func (c *Config) SaveConfig(config entity.Config) entity.Result {
	data, _ := json.Marshal(config)
	log.Printf(string(data))
	err := database.SaveConfig(config)

	if err != nil {
		return utils.Error("保存失败")
	}

	return utils.SuccessMsg("保存成功")
}
