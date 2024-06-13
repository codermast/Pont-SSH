<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { ExecCommand } from '../../wailsjs/go/service/Connection';
import { utils } from '../../wailsjs/go/models'
import Terminal, { Message } from "vue-web-terminal"
import 'vue-web-terminal/lib/theme/dark.css'
import {TerminalApi} from "vue-web-terminal"
// // 页面被挂载后
// onMounted(() => {
//   ExecCommand("run-parts /etc/update-motd.d/").then((result: utils.Result) => {
//     console.log(result)
//     if (result.code == 200) {
//       let welcomeMessage: string = result.msg;
//       let initialMessage: Message = {
//         "type": "html",
//         "content": welcomeMessage,
//         "class": "success",
//         "tag": "Tag success"
//       }
//       initLog.value.push(initialMessage)
//     }
//   })
// })

let context = ref("ubuntu")
let initLog = ref<Message[]>([{
  "type": "normal",
  "content": "Welcome using Pont SSH Management Tool",
  "class": "success",
  "tag": "Tag success"
}]);

function formatTextToHtml(inputText: string): string {
  // 将输入文本按行拆分成数组
  const trimmedText = inputText.trim();
  const lines = trimmedText.split('\n');

  // 如果行数小于阈值，则直接返回原始文本
  // TODO: 将阈值提取到系统设置中
  if (lines.length < 20) {
    return inputText;
  }

  // 遍历数组，将每行内容包裹在 <p> 标签内，并在结尾加上 <br />
  const formattedLines = lines.map(line => `<span>${ line }</span><br />`);

  // 将所有格式化后的行连接成一个字符串
  return formattedLines.slice(0, - 1).join('');
}

//  key 是 command 中的第一个单词，Command 是完整的命令
function onExecCmd(key: any, command: any, success: any, failed: any) {
  const ws = new WebSocket("ws://localhost:8080/terminal");

  // ExecCommand(command).then((result: utils.Result) => {
  //   if (result.code == 200) {
  //     success({
  //       type: 'normal',
  //       class: 'success',
  //       tag: '成功',
  //       content: formattedHtml
  //     })
  //   } else {
  //     failed(result.msg);
  //   }
  // })

  ExecCommand(command);

  ws.onmessage = function (event) {
    console.log(event.data);
    TerminalApi.appendMessage('my-terminal', event.data);
  };

}
</script>

<template>

  <div class="terminal">
    <terminal name="my-terminal"
              :context="context"
              :init-log="initLog"
              line-space="10"
              :show-header="false"
              @exec-cmd="onExecCmd">

      <template #header>
        This is my custom header
      </template>
    </terminal>
  </div>
</template>

<style scoped>
.terminal {
  display: block;
  height: 100vh;
}
</style>
