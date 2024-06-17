<script setup lang="ts">

import Search from "../components/Search.vue";
import Sidebar from "../components/Sidebar.vue";
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