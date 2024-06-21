<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Terminal } from 'xterm';
import { FitAddon } from 'xterm-addon-fit'
import { useMessage } from 'naive-ui'
import 'xterm/css/xterm.css';
import { useRouter } from 'vue-router'
import { useConfigStore } from "../stores/configStore";
import { WebLinksAddon } from '@xterm/addon-web-links'
import { BrowserOpenURL } from "../../wailsjs/runtime";
import ErrorConnect from "../dialogs/ErrorConnect.vue";
import { useDialogStore } from "../stores/dialogStore";

const router = useRouter()
const message = useMessage();
const configStore = useConfigStore();
const dialogStore = useDialogStore();

const terminal = new Terminal({
      fontSize: 16,
      theme: {
        background: '#2f3440', // 设置背景颜色
        foreground: '#ffffff'  // 设置前景（文本）颜色
      }
    }
);
const terminalContainer = ref(null)

// 自适应
const fitAddon = new FitAddon();
// url 可点击
const webLinksAddon = new WebLinksAddon(handleLink);

let socket: WebSocket;

onMounted(() => {
  // WebSocket Url
  const url = "ws://localhost:" + configStore.webSocketPort + '/ws';
  socket = new WebSocket(url);
  console.log(url)

  if (terminalContainer.value) {
    // 绑定自适应
    terminal.loadAddon(fitAddon);
    terminal.loadAddon(webLinksAddon);

    terminal.open(terminalContainer.value)

    // 刷新
    fitAddon.fit();
  }

  socket.onopen = (event) => {
    terminal.writeln('Welcome Using Pont Ssh Tool！');
  };

  socket.onmessage = (event) => {
    terminal.write(event.data);
  };

  socket.onclose = (event) => {
    terminal.write('连接关闭');
    dialogStore.errorConnectDialogVisible = true;
  };

  socket.onerror = (error) => {
    terminal.write('连接异常');
    dialogStore.errorConnectDialogVisible = true;
  };
});

// 存储用户输入
let buffer: string[] = [];

// 输入事件绑定
terminal.onData(data => {
  if (data === '\r') {
    console.log("回车了")
    terminal.write('\n\r');
    // buffer.push("\r")
    const cmd = buffer.join('')
    console.log(cmd);
    console.log(cmd.length);
    socket.send(cmd)
    buffer = [];
  } else if (data === '\x7F' || data === '\b' || data === '\x08') {
    console.log(buffer.join(""))
    // 删除
    if (buffer.length > 0) {
      terminal.write("\b \b");
      buffer.pop()
    } else {
      message.warning("已经清空")
    }
  } else if (data === '\u001b[A') {
    message.error("上箭头");
    // 处理上箭头逻辑
  } else if (data === '\u001b[B') {
    message.success("下箭头");
    // 处理下箭头逻辑
  } else if (data === '\u001b[C') {
    message.warning("右箭头");
    // 处理右箭头逻辑
  } else if (data === '\u001b[D') {
    message.info("左箭头");
    // 处理左箭头逻辑
  } else if (data === '\x03') { //  Ctrl+C
    console.log("发送中断 Ctrl + C 请求")
    socket.send("\x03");
  } else if (data === '\x16') { //  Ctrl+v
    // Ignore Ctrl+v
    console.log("ctrl + v");
  } else {
    terminal.write(data);
    buffer.push(data)
    console.log(buffer)
  }
})


// 键盘按键事件处理
terminal.onKey(event => {
  // 如果按下的是 Ctrl 键并且同时按下了 C 键，则发送 Ctrl + C 组合键信息
  if (event.domEvent.ctrlKey && event.key === 'c') {

  }

  if (event.key === 'v' && event.domEvent.ctrlKey) {
    // Ctrl+V 被按下
    // 进行处理...
    console.log("Ctrl+V 被按下");
  }
});


// 监听窗口大小变化事件，并调整终端大小
window.addEventListener('resize', () => {

  fitAddon.fit();
});

// 自定义链接处理函数
function handleLink(event: any, url: string) {
  // 1. 阻止原本的作用
  event.preventDefault();

  // 2. 执行自定义跳转操作
  BrowserOpenURL(url);
}
</script>


<template>
  <div class="terminal-container">
    <div class="terminal-box">
      <div class="terminal" ref="terminalContainer"></div>
    </div>
  </div>
  <ErrorConnect></ErrorConnect>

</template>

<style scoped>

.terminal-container {
  height: 100vh;
  width: 100%;
  background-color: #2f3440;
}

.terminal-box {
  position: absolute;
  top: 30px;
  left: 10px;
  right: 10px;
  bottom: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.terminal {
  overflow: hidden;
  flex-grow: 1;
  width: 100%;
  height: 100%;
}

</style>
