<script setup lang="ts">

import { FlashOutline } from "@vicons/ionicons5";
</script>

<template>
  <n-grid :cols="1">
    <n-gi :span="1">
      <n-input-group class="search-input">
        <!--  ssh 输入框  -->
        <n-input placeholder="Find a Server or ssh user@hostname...">
          <template #prefix>
            <n-icon :component="FlashOutline"/>
          </template>
        </n-input>
        <n-button class="search-submit">连接</n-button>
      </n-input-group>
    </n-gi>
  </n-grid>
</template>

<style scoped>

.search-input {
  display: flex;
  margin: 10px auto 10px;
  width: 73vw;
  border-radius: 20px;
}
</style>