<script setup lang="ts">
import CreateContainer from "../dialogs/NewConnect.vue";
import { useDialogStore } from "../stores/dialogStore";
import ExecCommand from "../components/ExecCommand.vue";
import SiderBar from "../components/SiderBar.vue";

const dialogStore = useDialogStore();

// 新增连接弹窗
function newConnectVisible() {
  dialogStore.newConnectDialogVisible = true;
}

</script>

<template>
  <n-grid cols="6" justify="center">
    <n-grid-item span="1">
      <div class="left-sidebar">
        <!-- 左侧边栏 -->
        <SiderBar></SiderBar>
        <n-button type="primary" @click="newConnectVisible">添加连接</n-button>
        <!-- 创建连接弹窗 -->
        <CreateContainer></CreateContainer>
      </div>
    </n-grid-item>
    <n-grid-item span="5">
      <n-grid cols="1">
        <n-grid-item>
          <div class="green">
            <ExecCommand style="height: 65vh"></ExecCommand>
          </div>

        </n-grid-item>
        <n-grid-item>
          <div class="yellow">
            这里是文件展示区
          </div>

        </n-grid-item>
      </n-grid>
    </n-grid-item>
  </n-grid>


</template>

<style scoped>
.left-sidebar {
  display: block;
  height: 100%;
}

.green {
  height: 70vh;
}

.yellow {
  height: 30vh;
}

.left-sidebar > * {
  margin-left: 5px;
}

p {
  padding: 0;
  margin: 0;
}


</style>