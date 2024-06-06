<script setup lang="ts">
import CreateContainer from "../dialogs/NewConnect.vue";
import { useDialogStore } from "../stores/dialogStore";

const dialogStore = useDialogStore();

// 新增连接弹窗
function newConnectVisible() {
  dialogStore.newConnectDialogVisible = true;
}

</script>

<template>
  <n-button type="primary" @click="newConnectVisible">添加连接</n-button>
  <CreateContainer></CreateContainer>
</template>

<style scoped>
</style>