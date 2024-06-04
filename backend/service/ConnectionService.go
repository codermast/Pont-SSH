package service

import (
	"PontSsh/backend/utils"
	"context"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

// Connection 连接结构体
type Connection struct {
	ctx context.Context
}

// SSHConnection SSH 连接配置结构体
type SSHConnection struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	KeyPath  string `json:"key_path"`
}

// 连接 Client 全局变量
var sshClient *ssh.Client

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
func (c *Connection) CreateConnection(sshConnection SSHConnection) utils.Result {
	server := sshConnection.Server
	port := sshConnection.Port
	socket := fmt.Sprintf("%s:%d", server, port)
	username := sshConnection.Username
	password := sshConnection.Password

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
		log.Fatalf("Failed to dial: %s", err)
	}
	defer client.Close()

	// 此时连接成功
	sshClient = client

	return utils.SuccessMsg("连接成功！")
}

// ExecCommand 执行 指令
func (c *Connection) ExecCommand(command string) utils.Result {
	// 创建一个新的会话
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	// 运行远程命令并捕获输出
	output, err := session.CombinedOutput(command)
	if err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}

	return utils.SuccessMsg(string(output))
}
