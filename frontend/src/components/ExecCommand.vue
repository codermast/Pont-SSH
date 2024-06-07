<script setup lang="ts">
import { ref } from 'vue'
import { ExecCommand } from '../../wailsjs/go/service/Connection';
import Terminal, { Message } from "vue-web-terminal"
import 'vue-web-terminal/lib/theme/dark.css'

let context = ref("ubuntu")
let initLog = ref<Message[]>([{
  "type": "normal",
  "content": "Welcome using Pont SSH Management Tool",
  "class": "success",
  "tag": "Tag success"
}]);

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
