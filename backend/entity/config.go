package entity

type Config struct {
	General  GeneralConfig  `yaml:"general"`
	Terminal TerminalConfig `yaml:"terminal"`
	Network  NetworkConfig  `yaml:"network"`
	Other    OtherConfig    `yaml:"other"`
}

// GeneralConfig 常规配置结构体
type GeneralConfig struct {
	Theme      string `yaml:"theme"`
	Language   string `yaml:"language"`
	FontFamily string `yaml:"font-family"`
}

// TerminalConfig 终端配置结构体
type TerminalConfig struct {
	Theme      string `yaml:"theme"`
	FontFamily string `yaml:"font-family"`
	FontSize   int    `yaml:"font-size"`
	LineNumber bool   `yaml:"line-number"`
	JumpTarget bool   `yaml:"jump-target"`
}

// NetworkConfig 网络配置结构体
type NetworkConfig struct {
	Speed bool `yaml:"speed"`
}

// OtherConfig 其他配置结构体
type OtherConfig struct {
}
