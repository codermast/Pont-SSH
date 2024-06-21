<script setup lang="ts">
import { GetWebSocketPort } from '../wailsjs/go/service/Connection'
import { useConfigStore } from "./stores/configStore";
import { onMounted, } from "vue";

const configStore = useConfigStore();

onMounted(() => {
  // 初始化 WebSocket 端口
  GetWebSocketPort().then(port => {
    configStore.webSocketPort = port;

    // console.log('ws://localhost:' + port + '/ws');
  })
})
</script>

<template>

  <div class="container">
    <n-dialog-provider>
      <n-modal-provider>
        <n-message-provider>
          <n-grid :cols="1">
            <n-gi :span="1" style="height: 100vh;width: 100%">
              <router-view></router-view>
            </n-gi>
          </n-grid>
        </n-message-provider>
      </n-modal-provider>
    </n-dialog-provider>
  </div>


</template>

<style scoped>

.container {
  overflow: hidden;
  padding: 0;
  margin: 0;
}
</style>