<script setup lang="ts">
import { ref } from 'vue'

import { ExecCommand } from '../../wailsjs/go/service/Connection';
import { useMessage } from "naive-ui";
import Terminal, { Message } from "vue-web-terminal"
import 'vue-web-terminal/lib/theme/light.css'

const message = useMessage();
// 指令变量
// let command = ref("");


let context = ref("ubuntu")
let initLog = ref<Message[]>([{
      "type": "normal",
      "content": "Welcome using Pont SSH Management Tool",
      "class": "success",
      "tag": "Tag success"
    },
    {
      "type": "normal",
      "content": "<a href='https://www.codermast.com/'>编程桅杆</a>",
      "class": "info",
      "tag": "Tag success"
    }]);

// 执行指令
function execCommand() {
  // console.log(command.value);
}

function onExecCmd(key: any, command: any, success: any, failed: any) {
  console.log(key);
  console.log(command);

  ExecCommand(command).then((result) => {
    console.log(result);

    if (result.code == 200) {
      success({
        type: 'normal',
        class: 'success',
        tag: '成功',
        content: result.msg
      })
    } else {
      failed(result.msg);
    }
  })
}


</script>

<template>
  <!--  <n-input placeholder="搜索" v-model:value="command">-->
  <!--    <template #prefix>-->
  <!--      <n-icon :component="FlashOutline"/>-->
  <!--    </template>-->
  <!--  </n-input>-->

  <!--  <n-button @click="execCommand">执行</n-button>-->
  <div class="terminal">
    <terminal name="my-terminal"
              :context="context"
              :init-log="initLog"
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
