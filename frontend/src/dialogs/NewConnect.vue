<script setup lang="ts">
import { ref, watch } from 'vue'
import { entity } from '../../wailsjs/go/models'
import { CreateConnection,SaveConnection } from '../../wailsjs/go/service/Connection'
import { useMessage } from 'naive-ui'
import { useDialogStore } from "../stores/dialogStore";
import { storeToRefs } from "pinia";

const dialogStore = useDialogStore()
let {newConnectDialogVisible} = storeToRefs(dialogStore)

let showModal = ref(newConnectDialogVisible)

watch(newConnectDialogVisible, (newValue) => {
  showModal.value = newValue
})

const message = useMessage();
let sshConfig = ref(entity.SSHConfig.createFrom({
  server: "",
  port: 22,
  username: "",
  password: "",
  keyPath: ""
}))

function testConnect() {
  CreateConnection(sshConfig.value).then((result) => {

    // 1. 判断是否成功
    if (result.code == 200) {
      message.success(result.msg);
    } else {
        message.error(result.msg);
    }
  })
}

// 确认点击
function saveConnect() {
  // 将 ssh 连接信息存入后端的数据库中
  SaveConnection(sshConfig.value).then((result) => {
    if (result.code == 200) {
      message.success(result.msg)
      showModal.value = false;
    }else{
      message.error(result.msg)
    }
  })
}

// 取消点击
function cancelClick() {
  showModal.value = false;
}
</script>

<template>
  <n-modal v-model:show="showModal"
           transform-origin="center"
           :mask-closable="false"
           class="main"
  >
    <n-card title="新建连接">

      <template #header>
        表头
      </template>

      <n-tabs
          type="line"
          animated
          placement="left"
          tab-style="justify-content: right; font-weight: 420;"
      >
        <n-tab-pane name="basic" tab="基本设置">
          <n-form
              :label-width="80"
              :model="sshConfig"
              label-placement="left"
              style="max-width: 600px;"
          >
            <n-form-item label="服务器IP">
              <n-input v-model:value="sshConfig.server" placeholder="输入服务器IP"/>
            </n-form-item>
            <n-form-item label="端口">
              <n-input v-model:value="sshConfig.port" placeholder="输入端口"/>
            </n-form-item>
            <n-form-item label="用户名">
              <n-input v-model:value="sshConfig.username" placeholder="输入用户名"/>
            </n-form-item>
            <n-form-item label="密码">
              <n-input v-model:value="sshConfig.password" placeholder="输入密码"/>
            </n-form-item>
          </n-form>
        </n-tab-pane>
        <n-tab-pane name="advance" tab="高级设置">
          高级设置
        </n-tab-pane>
        <n-tab-pane name="proxy" tab="代理设置">
          代理设置
        </n-tab-pane>
        <n-tab-pane name="other" tab="其他设置">
          其他设置
        </n-tab-pane>
      </n-tabs>
      <template #footer>
        <div class="buttons-container">
          <n-button class="submitButton" type="info" @click="testConnect">
            验证
          </n-button>

          <n-button class="submitButton" type="primary" @click="saveConnect">
            保存
          </n-button>

          <n-button class="submitButton" type="warning" @click="cancelClick">
            取消
          </n-button>
        </div>
      </template>
    </n-card>
  </n-modal>
</template>

<style scoped>

.main {
  position: absolute;
  top: -30vh;
  bottom: 0;
  left: 0;
  right: 0;
  margin: auto;
  width: 500px;
  height: 400px;
}

.buttons-container {
  width: 100%;
  display: flex;
  /* 水平居中 */
  justify-content: center;
  /* 垂直居中 */
  align-items: center;
  /* 边距 */
  padding: 10px;
}

.submitButton {
  margin: 10px;
}
</style>