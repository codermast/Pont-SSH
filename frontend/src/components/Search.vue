<script setup lang="ts">

import { FlashOutline } from "@vicons/ionicons5";

import { ref } from 'vue'
import NewConnect from "../dialogs/NewConnect.vue";
import { entity } from "../../wailsjs/go/models";
import { useDialogStore } from "../stores/dialogStore";

const dialogStore = useDialogStore();

// SSH 连接指令
let sshCommand = ref<string>("ssh codermast@10.211.55.17");

let sshConnectConfig = ref<entity.SSHConfig>({
  id: "123",
  server: "10.211.55.17",
  port: 22,
  username: "codermast",
  password: "dongpeng@123.",
  name: "string",
  edit: false,
});

// SSH 指令连接提交
function sshCommandConnect() {
  let username = sshCommand.value.split("@")[0];
  let server = sshCommand.value.split("@")[1];

  console.log(username);
  console.log(server);
  sshConnectConfig.value.port = 22;
  sshConnectConfig.value.username = username;
  sshConnectConfig.value.server = server;

  dialogStore.newConnectDialogVisible = true;
}

</script>

<template>
  <n-grid :cols="1">
    <n-gi :span="1">
      <n-input-group class="search-input">
        <!--  ssh 输入框  -->
        <n-input
            type="text"
            placeholder="Input ssh user@hostname to link your server"
            v-model:value="sshCommand"
        >
          <template #prefix>
            <n-icon :component="FlashOutline"/>
          </template>
        </n-input>
        <n-button @click="sshCommandConnect" class="search-submit">连接</n-button>
      </n-input-group>

    </n-gi>
  </n-grid>

  <!-- 新建连接模态框 -->
  <NewConnect :ssh-connect-config="sshConnectConfig"></NewConnect>
</template>

<style scoped>

.search-input {
  display: flex;
  margin: 10px auto 10px;
  width: 73vw;
  border-radius: 20px;
}
</style>