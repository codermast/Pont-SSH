<script setup lang="ts">
import {ref} from 'vue'
import { FlashOutline } from '@vicons/ionicons5'

import { ExecCommand } from  '../../wailsjs/go/service/Connection';

// 指令变量
let command = ref("");

// 执行指令
function execCommand(){
  ExecCommand(command.value).then((result) => {
    console.log(result);
  })
}
</script>

<template>
  <n-input placeholder="搜索" v-model="command">
    <template #prefix>
      <n-icon :component="FlashOutline"/>
    </template>
  </n-input>

  <n-button @click="execCommand">执行</n-button>
</template>

<style scoped>

</style>