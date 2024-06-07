<template>
  <div class="sidebar">
    <p>同步状态：<span class="status-icon">✔️</span></p>
    <p class="ip-container">
      IP {{ ipAddress }}
      <n-button @click="copyIp" type="info" >复制</n-button>
    </p>
    <p>运行 {{ runStatus }}</p>
    <p>负载 {{ loadStatus }}</p>
    <p>CPU
      <n-progress type="line" :percentage="cpuUsage" indicator-placement="inside" status="success"/>
    </p>
    <p>内存
      <n-progress type="line" :percentage="memoryUsage" indicator-placement="inside" status="info"/>
    </p>
    <p>交换
      <n-progress type="line" :percentage="swapUsage" indicator-placement="inside" status="warning"/>
    </p>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { NProgress, useMessage } from 'naive-ui';
const message = useMessage();

const ipAddress = ref('192.168.1.1');
const runStatus = ref('*');
const loadStatus = ref('*');
const cpuUsage = ref(18);
const memoryUsage = ref(36);
const swapUsage = ref(54);

const copyIp = () => {
  navigator.clipboard.writeText(ipAddress.value);
  message.success('IP 已复制');
};

</script>

<style scoped>

.sidebar {
}

</style>
