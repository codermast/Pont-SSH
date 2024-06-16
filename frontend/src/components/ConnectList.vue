<script setup lang="ts">

import NewConnect from "../dialogs/NewConnect.vue";

import { h, ref } from 'vue'
import { useDialogStore } from "../stores/dialogStore";

import { Add } from '@vicons/ionicons5'
import { NIcon } from 'naive-ui'

const dialogStore = useDialogStore();

// 新增连接弹窗
function newConnectVisible() {
  dialogStore.newConnectDialogVisible = true;
}


function renderIcon() {
  return h(NIcon, null, {
    default: () => h(Add)
  })
}

const modal = ref(true);

</script>

<template>
  <div class="content-container">

        <n-button
            secondary
            strong
            :render-icon="renderIcon"
            @click="newConnectVisible"
        >
          添加新连接
        </n-button>

  </div>

  <!-- 创建连接弹窗 -->
  <NewConnect></NewConnect>
</template>

<style scoped>
.content-container {
  background-color: lightblue;
  height: 100vh;
  width: 100%;
}

</style>