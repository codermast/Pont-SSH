package service

import (
	"PontSsh/backend/database"
	"PontSsh/backend/entity"
	"PontSsh/backend/utils"
	"context"
	"fmt"
	"golang.org/x/crypto/ssh"
)

// Connection 连接结构体
type Connection struct {
	ctx context.Context
}

// 连接 Client 全局变量
var (
	sshSession *ssh.Session
)

// NewConnection 创建一个 Connection 对象
func NewConnection() *Connection {
	return &Connection{}
}

// Startup 启动方法，用来注册一些资源
func (c *Connection) Startup(ctx context.Context) {
	c.ctx = ctx
}

// Shutdown 销毁方法，用户释放一些资源
func (c *Connection) Shutdown(ctx context.Context) {

}

// CreateConnection 创建一个新连接
func (c *Connection) CreateConnection(sshConfig entity.SSHConfig) utils.Result {
	server := sshConfig.Server
	port := sshConfig.Port
	socket := fmt.Sprintf("%s:%d", server, port)
	username := sshConfig.Username
	password := sshConfig.Password

	// 创建 SSH 客户端配置
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password), // 使用密码进行身份验证
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 在生产环境中应使用更加安全的方法来验证主机密钥
	}

	// 连接到 SSH 服务器
	client, err := ssh.Dial("tcp", socket, config)
	if err != nil {
		return utils.Error(fmt.Sprintf("Failed to dial: %s", err))
	}

	// 运行远程命令
	session, err := client.NewSession()
	if err != nil {
		return utils.Error(fmt.Sprintf("Failed to create session: %s", err))
	}

	// 全局变量保存
	sshSession = session

	return utils.SuccessMsg("连接成功！")
}

// SaveConnection 保存新连接
func (c *Connection) SaveConnection(sshConfig entity.SSHConfig) utils.Result {
	err := database.SaveSshConnect(sshConfig)
	if err != nil {
		return utils.Error("保存失败！")
	}
	return utils.SuccessMsg("保存成功！")
}

// ExecCommand 执行 指令
func (c *Connection) ExecCommand(command string) utils.Result {
	fmt.Println("要执行的命令：" + command)
	if sshSession == nil {
		return utils.Error(fmt.Sprintf("Session 为空！"))
	}

	output, err := sshSession.CombinedOutput(command)
	if err != nil {
		return utils.Error(fmt.Sprintf("Failed to execute command: %s", err))
	}
	fmt.Println(string(output))
	return utils.SuccessMsg(string(output))
}
