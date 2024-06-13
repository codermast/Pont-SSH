package service

import (
	"PontSsh/backend/database"
	"PontSsh/backend/entity"
	"PontSsh/backend/utils"
	_ "bytes"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net/http"
)

// Connection 连接结构体
type Connection struct {
	ctx context.Context
}

// 连接 Client 全局变量
var (
	sshClient *ssh.Client
)

// NewConnection 创建一个 Connection 对象
func NewConnection() *Connection {
	return &Connection{}
}

// Http 连接升级到 Websocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Startup 启动方法，用来注册一些资源
func (c *Connection) Startup(ctx context.Context) {
	c.ctx = ctx

	// 启动时，开启 WebSocket
	http.HandleFunc("/terminal", wsEndpoint)

	// 异步启动 HTTP 服务器，监听 8080 端口
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("Handle server start error : %v", err)
		}
	}()
}

var (
	sshWs *websocket.Conn
)

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to upgrade to websocket: %s", err)
	}
	sshWs = websocket
}

// Shutdown 销毁方法，用户释放一些资源
func (c *Connection) Shutdown(ctx context.Context) {
	// 管理 WebSocket 连接
	err := sshWs.Close()
	if err != nil {
		return
	}
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

	sshClient = client

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
	// 先判断 Client 是否为空
	if sshClient == nil {
		return utils.Error(fmt.Sprintf("Client 为空！"))
	}

	// 获取 Session 对象
	session, err := sshClient.NewSession()
	if err != nil {
		return utils.Error(fmt.Sprintf("Failed to create session: %s", err))
	}

	defer session.Close()

	// 0. 先获取标准输出流，以免数据遗漏
	stdout, err := session.StdoutPipe()
	if err != nil {
		return utils.Error(fmt.Sprintf("Failed to create stdout: %s", err))
	}

	// 1. 启动命令执行后，并不会阻塞，而是会继续向后执行
	err = session.Start(command)
	if err != nil {
		return utils.Error(fmt.Sprintf("Command Start Execption: %s", err))
	}

	// 2. 获取标准 输入、输出、异常 流

	if sshWs == nil {
		return utils.Error(fmt.Sprintf("WebSocket Init Execption: %s", err))
	}

	// 读取标准输出数据并发送到 WebSocket
	go func() {
		for {
			// 读取标准输出流中的数据
			buf := make([]byte, 4096)
			n, err := stdout.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Printf("Failed to read stdout: %v", err)
				break
			}
			// 发送数据到 WebSocket 连接
			if err := sshWs.WriteMessage(websocket.TextMessage, buf[:n]); err != nil {
				log.Printf("Failed to send message to WebSocket: %v", err)
				break
			}
		}
	}()

	// 3. 等待指令执行完毕
	err = session.Wait()
	if err != nil {
		return utils.Error(fmt.Sprintf("Command execution failed: %v", err))
	}

	return utils.SuccessMsg("执行成功！")
}
