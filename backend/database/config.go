package database

import (
	"PontSsh/assets"
	"PontSsh/backend/constant"
	"PontSsh/backend/entity"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// 配置文件
var configFile *os.File

// LoadConfig 加载配置文件
func LoadConfig() {
	log.Println("配置文件加载！")

	filePath := constant.ConfigFilePath
	data := assets.Config

	// 0. 判断是否需要初始化
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		log.Println("系统配置文件不存在，则需要进行初始化")
		// 1. 写入配置文件
		err = os.WriteFile(filePath, data, 0644)

		if err != nil {
			log.Printf("文件写入异常%v", err)
		}
	} else {
		log.Println("系统配置文件已经存在，无需进行初始化")
	}
}

// ConfigStartup 配置文件启动方法
func ConfigStartup() {
	// 打开配置文件
	file, err := os.Open(constant.ConfigFilePath)

	if err != nil {
		log.Printf("配置文件打开异常！")
	}

	configFile = file

}

// ConfigShutdown 配置文件结束方法
func ConfigShutdown() {
	configFile.Close()
}

// GetConfig 获取配置信息
func GetConfig() (entity.Config, error) {
	// 配置对象
	var config entity.Config

	// 重新定位文件指针到开头
	_, err := configFile.Seek(0, 0)

	if err != nil {
		fmt.Println("Error seeking file:", err)
		return entity.Config{}, err
	}

	// 创建解析器
	decoder := yaml.NewDecoder(configFile)

	// 解析 YAML 数据
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding YAML:", err)
		return entity.Config{}, err
	}
	return config, nil
}

// ResetConfig 重置配置信息
func ResetConfig() error {
	err := os.Remove(constant.ConfigFilePath)

	if err != nil {
		log.Println("删除文件失败")
		return err
	}

	LoadConfig()
	return nil
}

// SaveConfig 保存配置信息
func SaveConfig(config entity.Config) error {
	file, err := os.OpenFile(constant.ConfigFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		log.Printf("读出文件异常%v", err)
		return err
	}

	encoder := yaml.NewEncoder(file)

	err = encoder.Encode(config)
	if err != nil {
		log.Printf("写入文件异常%v", err)
		return err
	}
	return nil
}
