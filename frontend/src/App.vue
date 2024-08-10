<script setup lang="ts">
import { Environment, WindowToggleMaximise } from '../wailsjs/runtime'
import { onMounted, ref } from "vue";
import { TerminalOutline } from '@vicons/ionicons5'

let os = ''

let logoPaddingLeft = ref(80)


onMounted(async () => {
  await loadEnvironment()

  if (isWindows()) {
    logoPaddingLeft.value = 10;
  }
})

async function loadEnvironment() {
  const env = await Environment()
  console.log("env", env)
  os = env.platform
}

function isMacOS() {
  return os === 'darwin'
}

function isWindows() {
  return os === 'windows'
}


</script>

<template>

  <!-- title bar -->
  <div
      style="--wails-draggable: drag;
      user-select:none;
        height: 38px;
        background-color: #f5f6f7;
         display: flex;
        align-items: center;
"
      @dblclick="WindowToggleMaximise">

    <!-- title -->
    <div
        :style="{
            paddingLeft: `${logoPaddingLeft}px`,
            display: `flex`,
            alignItems: `center`,
        }">
      <n-icon :size="20">
        <TerminalOutline/>
      </n-icon>

      <div
          style="
                min-width: 68px;
                white-space: nowrap;
                font-weight: 800;
                padding-left: 10px;
            "
      >
        Pont SSH
      </div>
    </div>
  </div>
  <router-view></router-view>

</template>

<style scoped>

</style>