<script setup lang="ts">

import NewConnect from "../dialogs/NewConnect.vue";
import { useDialogStore } from "../stores/dialogStore";

const dialogStore = useDialogStore();

// 新增连接弹窗
function newConnectVisible() {
  dialogStore.newConnectDialogVisible = true;
}

</script>

<template>

  <n-button type="primary" @click="newConnectVisible">添加连接</n-button>
  <!-- 创建连接弹窗 -->
  <NewConnect></NewConnect>
</template>

<style scoped>
</style>