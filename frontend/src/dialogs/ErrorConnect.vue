<script setup lang="ts">
import { useDialogStore } from "../stores/dialogStore";
import { useMessage } from 'naive-ui'
import { useRouter } from "vue-router";

const dialogStore = useDialogStore();
const message = useMessage();
const router = useRouter();

// 返回首页
function goHome() {
  router.push({name: "Index"})
  // 关闭模态框
  dialogStore.errorConnectDialogVisible = false;
}

// 重新连接
function reConnect() {
  dialogStore.errorConnectDialogVisible = false;
}

// 退出连接
function exitConnect() {
  router.back()
  dialogStore.errorConnectDialogVisible = false;
}

</script>

<template>
  <!-- 连接断开、异常弹出框-->
  <n-modal
      v-model:show="dialogStore.errorConnectDialogVisible"
      transform-origin="center"
  >
    <n-card class="operationList">
      <template #header>
        连接异常操作列表
      </template>
      <n-button type="default" @click="goHome">返回首页</n-button>
      <n-button type="info" @click="reConnect">重新连接</n-button>
      <n-button type="warning" @click="exitConnect">退出连接</n-button>
    </n-card>
  </n-modal>
</template>

<style scoped>
.operationList {
  display: flex;
  flex-direction: column;
  width: 200px;
  text-align: center;
}
</style>