<script setup lang="ts">
import { Save } from '@vicons/ionicons5'
import Cancel from '../icons/Cancel.vue'
import Reset from '../icons/Reset.vue'

// 重置配置
function resetConfig(){

}

// 应用
function applyConfig(){

}

// 取消
function cancelConfig(){

}
</script>

<template>
  <n-card title="系统设置" class="systemConfig">

    <n-tabs
        placement="left"
        default-value="general"
    >
      <n-tab-pane name="general" tab="常规配置">
        <n-form>
          <n-form-item></n-form-item>
        </n-form>
      </n-tab-pane>

      <n-tab-pane name="terminal" tab="终端配置">
        终端配置
      </n-tab-pane>

      <n-tab-pane name="internal" tab="网络配置">
        网络配置
      </n-tab-pane>

      <n-tab-pane name="other" tab="其他配置">
        其他配置
      </n-tab-pane>
    </n-tabs>

    <template #footer>
      <n-button
          type="tertiary"
          class="reset"
          @click="resetConfig"
      >
        <template #icon>
          <n-icon>
            <Reset />
          </n-icon>
        </template>
        重置
      </n-button>
      <div class="funcButt">

        <n-button
            type="success"
            class="apply"
            @click="applyConfig"
        >
          <template #icon>
            <n-icon>
              <save />
            </n-icon>
          </template>
          应用
        </n-button>

        <n-button
            type="error"
            class="cancel"
            @click="cancelConfig"
        >
          <template #icon>
            <n-icon>
              <Cancel />
            </n-icon>
          </template>
          取消
        </n-button>
      </div>
    </template>
  </n-card>

</template>

<style scoped>

.systemConfig {
  width: 90%;
  margin: auto;
}

.reset {
  float: left;
}

.funcButt{
  float: right;
}

.apply{
  margin: 0 10px;
}


</style>