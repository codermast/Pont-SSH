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

// NewConnection 创建一个 Connection 对象
func NewConnection() *Connection {
	return &Connection{}
}

// Startup 启动方法，用来注册一些资源
func (c *Connection) Startup(ctx context.Context) {
	c.ctx = ctx

	// 启动 WebSocket 服务器
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("WebSocket server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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
			_, message, err := conn.ReadMessage()
			fmt.Print(string(message))
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

// GetPrompt 获取主机头
func (c *Connection) GetPrompt() utils.Result {

	// 先判断 Client 是否为空
	if sshClient == nil {
		log.Printf("Client 为空！")
		return utils.Error(fmt.Sprintf("Client 为空！"))
	}

	// 获取 Session 对象
	session, err := sshClient.NewSession()
	if err != nil {
		log.Printf("Failed to create session: %s", err)
		return utils.Error(fmt.Sprintf("Failed to create session: %s", err))
	}
	defer session.Close()

	// 启动之前先拼接主机头信息
	prompt, err := session.CombinedOutput("pwd_short=\"${PWD/#$HOME/}\"  \n[[ $pwd_short == / ]] && pwd_short='~' || pwd_short=\"~/$pwd_short\"  \nprintf \"%s\" \"$(whoami)@$(hostname):$pwd_short\"")
	if err != nil {
		log.Printf("prompt failed: %s", err)
		return utils.Error(fmt.Sprintf("prompt failed: %s", err))
	}

	return utils.SuccessData(string(prompt))
}

//// ExecCommand 执行指令
//func (c *Connection) ExecCommand(command string) {
//	// 先判断 Client 是否为空
//	if sshClient == nil {
//		log.Printf("Client 为空！")
//		return
//	}
//
//	// 获取 Session 对象
//	session, err := sshClient.NewSession()
//	if err != nil {
//		log.Printf("Failed to create session: %s", err)
//		return
//	}
//	defer session.Close()
//
//	// 请求伪终端
//	modes := ssh.TerminalModes{
//		ssh.ECHO:          1,     // 是否回显输入
//		ssh.TTY_OP_ISPEED: 14400, // 输入速度
//		ssh.TTY_OP_OSPEED: 14400, // 输出速度
//	}
//	err = session.RequestPty("xterm", 80, 40, modes)
//	if err != nil {
//		log.Printf("request for pseudo terminal failed: %s", err)
//		return
//	}
//
//	// 获取标准输出流和标准错误流
//	stdout, err := session.StdoutPipe()
//	if err != nil {
//		log.Printf("Failed to create stdout: %s", err)
//		return
//	}
//	stderr, err := session.StderrPipe()
//	if err != nil {
//		log.Printf("Failed to create stderr: %s", err)
//		return
//	}
//
//	// 启动命令执行
//	err = session.Shell()
//	if err != nil {
//		log.Printf("Failed to start command: %s", err)
//		handleError(c.ctx, stderr)
//		return
//	}
//
//	// 使用 bufio.Scanner 按行读取标准输出流
//	scanner := bufio.NewScanner(stdout)
//	for scanner.Scan() {
//		log.Printf("stdout: %s", scanner.Text())
//		// 触发事件，向前端传递数据
//		runtime.EventsEmit(c.ctx, "sendResult", scanner.Text())
//	}
//
//	// 等待命令执行完成
//	err = session.Wait()
//	if err != nil {
//		log.Printf("Command execution failed: %v", err)
//		handleError(c.ctx, stderr)
//	}
//}
//
//// handleError 用于处理标准错误流
//func handleError(ctx context.Context, stderr io.Reader) {
//	scanner := bufio.NewScanner(stderr)
//	for scanner.Scan() {
//		log.Printf("stderr: %s", scanner.Text())
//		// 触发事件，向前端传递数据
//		runtime.EventsEmit(ctx, "sendResult", scanner.Text())
//	}
//}
