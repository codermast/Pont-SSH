<script setup lang="ts">
import { useMessage } from 'naive-ui'
import { LogoGithub } from '@vicons/ionicons5'

const message = useMessage()

function Duck() {
  message.success('Duck')
}
</script>

<template>
  <div class="ribbon">
    <n-tooltip trigger="hover">
      <template #trigger>
        <n-icon size="40" v-show="true" @click="Duck">
          <logo-github/>
        </n-icon>
      </template>
      如果它长得像鸭子，走起来像鸭子，叫起来也像鸭子，那它一定是个鸭子。
    </n-tooltip>
  </div>

</template>

<style scoped>
.ribbon {
  background-color: pink;
  width: 100%;
  height: 100vh;
}
</style>