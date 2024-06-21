<script setup lang="ts">
import { ref, watch } from 'vue'
import { entity } from '../../wailsjs/go/models'
import { SaveConnection, TestConnection } from '../../wailsjs/go/service/Connection'
import { useMessage } from 'naive-ui'
import { useDialogStore } from "../stores/dialogStore";
import { storeToRefs } from "pinia";
import { useRouter } from "vue-router";

const router = useRouter()

const dialogStore = useDialogStore()


const message = useMessage();
const sshConfig = ref<entity.SSHConfig>({
  id: "123",
  server: "10.211.55.17",
  port: 22,
  username: "codermast",
  password: "dongpeng@123.",
  name: "string",
  edit: false,
});

function testConnect() {
  TestConnection(sshConfig.value).then((result) => {
    // 1. 判断是否成功
    if (result.code == 200) {
      message.success(result.msg);
      // 关闭模态框
      // showModal.value = false;
      // router.push({name: "Cmd"});
    } else {
      message.error(result.msg);
    }
  })
}

// 保存点击事件
function saveConnect() {
  // 将 ssh 连接信息存入后端的数据库中
  SaveConnection(sshConfig.value).then((result) => {
    if (result.code == 200) {
      message.success(result.msg)
      dialogStore.newConnectDialogVisible = false;
    } else {
      message.error(result.msg)
    }
  })
}

// 取消点击
function cancelClick() {
  dialogStore.newConnectDialogVisible = false;
}
</script>

<template>
  <n-modal v-model:show="dialogStore.newConnectDialogVisible"
           transform-origin="center"
           :mask-closable="false"
           class="main"
  >
    <n-card title="新建连接">
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
              <n-input v-model:value="sshConfig.password" type="password" placeholder="输入密码"/>
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
  bottom: 0;
  left: 0;
  right: 0;
  margin: auto;
  width: 500px;
  max-height: 100vh;
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