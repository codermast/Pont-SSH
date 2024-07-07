<script setup lang="ts">

import Sidebar from "./Sidebar.vue";
import { onMounted } from "vue";
import { GetConfig } from "../../wailsjs/go/service/Config";
import {useConfigStore} from "../stores/configStore";
import { useMessage } from "naive-ui";
const configStore = useConfigStore();
const message = useMessage();

// 1. 获取配置信息
onMounted(() => {
  GetConfig().then((result) => {
    console.log(result)
    if (result.code == 200) {
      // 将数据存储到 configStore 中
      configStore.systemConfig = result.data;
    } else {
      message.error(result.msg)
    }
  })
})

</script>

<template>
  <div class="content-container">
    <n-grid :cols="4">
      <n-gi :span="1">

        <Sidebar></Sidebar>
      </n-gi>
      <n-gi :span="3">
        <router-view></router-view>
      </n-gi>
    </n-grid>

  </div>
</template>

<style scoped>
.content-container {
  height: 100vh;
  width: 100%;
}
</style>