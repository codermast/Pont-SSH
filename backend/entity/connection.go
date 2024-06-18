package entity

// SSHConfig SSH 连接配置结构体
type SSHConfig struct {
	Server   string `json:"server" yaml:"server"`
	Port     int    `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Name     string `json:"name" json:"name"`
	Edit     bool   `json:"edit" yaml:"edit"`
}
