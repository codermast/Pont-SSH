package service

import (
	"PontSsh/backend/database"
	"PontSsh/backend/entity"
	"PontSsh/backend/utils"
	_ "bytes"
	"context"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"net/http"
	_ "os"
	"strings"
	_ "strings"

	"github.com/gorilla/websocket"
)

// Connection 连接结构体
type Connection struct {
	ctx context.Context
}

// 连接 Client 全局变量
var (
	sshClient *ssh.Client
)

// 启动端口
var webSocketPort int

// 将 http 连接升级为 websocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// NewConnection 创建一个 Connection 对象
func NewConnection() *Connection {
	return &Connection{}
}

// Startup 启动方法，用来注册一些资源
func (c *Connection) Startup(ctx context.Context) {
	c.ctx = ctx

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handleWebSocket)

	// Create a listener on any available port (port 0)
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	// Output the allocated port
	port := listener.Addr().(*net.TCPAddr).Port
	log.Printf("WebSocket server running on port %d\n", port)

	webSocketPort = port

	// Serve the HTTP server using the listener
	go func() {
		if err := http.Serve(listener, mux); err != nil {
			log.Fatal("Failed to serve:", err)
		}
	}()

	// 初始化数据库
	err = database.InitDatabase()
	if err != nil {
		log.Printf("Failed to init database: %s", err)
	}
}

// Shutdown 销毁方法，用户释放一些资源
func (c *Connection) Shutdown(ctx context.Context) {
	err := sshClient.Close()
	if err != nil {
		log.Println("Failed to close ssh client:", err)
	}
}

// WebSocket 处理器
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("websocket upgrade:", err)
		return
	}
	defer conn.Close()

	if sshClient == nil {
		log.Println("ssh client is nil")
	}

	// 创建一个新的会话
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	// 请求伪终端
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // 是否回显输入
		ssh.TTY_OP_ISPEED: 14400, // 输入速度
		ssh.TTY_OP_OSPEED: 14400, // 输出速度
	}

	if err := session.RequestPty("xterm-256color", 80, 100, modes); err != nil {
		log.Fatalf("request for pseudo terminal failed: %s", err)
	}

	// 启动一个Shell
	stdoutPipe, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed to create stdout pipe: %s", err)
	}
	stderrPipe, err := session.StderrPipe()
	if err != nil {
		log.Fatalf("Failed to create stderr pipe: %s", err)
	}
	stdinPipe, err := session.StdinPipe()
	if err != nil {
		log.Fatalf("Failed to create stdin pipe: %s", err)
	}

	// 输入控制
	go func() {
		for {
			messageType, message, err := conn.ReadMessage()

			fmt.Println(message)
			fmt.Println(messageType)

			if err != nil {
				log.Println("websocket read:", err)
				break
			}

			if _, err := stdinPipe.Write(append(message, '\n')); err != nil {
				log.Println("Failed to write to stdin:", err)
				break
			}
		}
	}()

	// 输出控制
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stdoutPipe.Read(buf)
			if err != nil {
				log.Println("stdout read:", err)
				break
			}

			if strings.TrimSpace(string(buf[:n])) == "reboot" {
				if err := conn.WriteMessage(websocket.TextMessage, []byte("System is going to reboot, disconnecting...")); err != nil {
					log.Println("Failed to reboot:", err)
				}
			}

			if err := conn.WriteMessage(websocket.TextMessage, buf[:n]); err != nil {
				log.Println("websocket write:", err)
				break
			}
		}
	}()

	// 控制异常
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stderrPipe.Read(buf)
			if err != nil {
				log.Println("stderr read:", err)
				break
			}
			if err := conn.WriteMessage(websocket.TextMessage, buf[:n]); err != nil {
				log.Println("websocket write:", err)
				break
			}
		}
	}()

	if err := session.Shell(); err != nil {
		log.Printf("Failed to run: %s", err)
	}

	if err := session.Wait(); err != nil {
		log.Printf("Failed to run: %s", err)
	}
}

// TestConnection 连接测试
func (c *Connection) TestConnection(sshConfig entity.SSHConfig) entity.Result {
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
	_, err := ssh.Dial("tcp", socket, config)
	if err != nil {
		return utils.Error(fmt.Sprintf("Failed to dial: %s", err))
	}

	return utils.SuccessMsg("连接成功！")
}

// SaveConnection 保存新连接
func (c *Connection) SaveConnection(sshConfig entity.SSHConfig) entity.Result {
	err := database.SaveSshConnect(sshConfig)
	if err != nil {
		return utils.Error("保存失败！")
	}
	return utils.SuccessMsg("保存成功！")
}

// ServerConnection 连接服务器
func (c *Connection) ServerConnection(sshConfig entity.SSHConfig) entity.Result {
	server := sshConfig.Server
	port := sshConfig.Port
	socket := fmt.Sprintf("%s:%d", server, port)
	log.Println(socket)
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

// GetWebSocketPort 返回 WebSocketPort
func (c *Connection) GetWebSocketPort() int {
	return webSocketPort
}

// GetServerList 查询服务器列表
func (c *Connection) GetServerList() entity.Result {
	serverList, err := database.GetServerList()
	if err != nil {
		log.Println("Query server list error:", err)
		return utils.Error("查询失败！")
	}

	return utils.Success(serverList, "查询成功！")
}
