<script setup lang="ts">

</script>

<template>
  <div class="container">
    <n-modal-provider>
      <n-message-provider>
        <router-view></router-view>
      </n-message-provider>
    </n-modal-provider>
  </div>
</template>
<style scoped>
.container {
  overflow: hidden;
  padding: 0;
  margin: 0;
}
</style>