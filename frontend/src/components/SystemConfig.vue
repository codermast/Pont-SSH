<script setup lang="ts">
import { onMounted, ref, toRefs } from 'vue'
import { Save } from '@vicons/ionicons5'
import Cancel from '../icons/Cancel.vue'
import Reset from '../icons/Reset.vue'
import { useMessage } from "naive-ui"
import { fontFamilies, languages } from "../utils/options";
import { GetConfig, ResetConfig, SaveConfig } from "../../wailsjs/go/service/Config";
import {useConfigStore} from "../stores/configStore";
const configStore = useConfigStore();

const message = useMessage();

let config = ref({
  general: {
    Theme: "auto",
    Language: "default",
    FontFamily: "default",
    FontSize: "12",
  },
  terminal: {
    Theme: "auto",
    Language: "default",
    FontFamily: "default",
    FontSize: "13",
    LineNumber: true,
    JumpTarget: true,
  },
  network: {
    Speed: false,
  },
  other: {}
});

// 解构配置项
let {general, terminal, network, other} = toRefs(config.value)

onMounted(() => {
  // 获取配置信息数据
  general.value = configStore.systemConfig.General;
  terminal.value = configStore.systemConfig.Terminal;
  network.value = configStore.systemConfig.Network;
  other.value = configStore.systemConfig.Other;
})



// 重置配置
function resetConfig() {
  ResetConfig().then((result) => {
    if (result.code == 200) {
      general.value = result.data.General;
      terminal.value = result.data.Terminal;
      network.value = result.data.Network;
      other.value = result.data.Other;
      message.success(result.msg);
    } else {
      message.error(result.msg);
    }
  })
}

// 应用
function applyConfig() {
  SaveConfig(config.value).then((result) => {
    if (result.code == 200) {
      message.success(result.msg)
    } else {
      message.error(result.msg)
    }
  })
}

// 取消
function cancelConfig() {
  general.value = configStore.systemConfig.General;
  terminal.value = configStore.systemConfig.Terminal;
  network.value = configStore.systemConfig.Network;
  other.value = configStore.systemConfig.Other;

  message.info("配置取消修改，已还原");
}
</script>

<template>
  <n-card title="系统设置" class="systemConfig">

    <template #header-extra>
      <n-gradient-text type="info">
        欢迎使用 「Pont SSH」连接工具！
      </n-gradient-text>
    </template>
    <n-tabs
        placement="left"
        default-value="general"
    >
      <n-tab-pane name="general" tab="常规配置">
        <n-form label-placement="left">
          <n-form-item label="主题">
            <n-radio-group
                v-model:value="general.Theme"
                name="left-size"
            >
              <n-radio-button value="light">
                浅色
              </n-radio-button>
              <n-radio-button value="auto">
                自动
              </n-radio-button>
              <n-radio-button value="dark">
                深色
              </n-radio-button>
            </n-radio-group>
          </n-form-item>

          <n-form-item label="语言">
            <n-select default-value="系统默认" :options="languages" v-model:value="general.Language">
            </n-select>
          </n-form-item>

          <n-form-item label="字体">
            <n-select default-value="系统默认" :options="fontFamilies" v-model:value="general.FontFamily"></n-select>
          </n-form-item>
        </n-form>
      </n-tab-pane>

      <n-tab-pane name="terminal" tab="终端配置">
        <n-form label-placement="left">
          <n-form-item label="主题">
            <n-radio-group
                v-model:value="terminal.Theme"
                name="left-size"
            >
              <n-radio-button value="light">
                浅色
              </n-radio-button>
              <n-radio-button value="auto">
                自动
              </n-radio-button>
              <n-radio-button value="dark">
                深色
              </n-radio-button>
            </n-radio-group>
          </n-form-item>

          <n-form-item label="字体">
            <n-select default-value="系统默认" :options="fontFamilies" v-model:value="general.FontFamily"></n-select>
          </n-form-item>

          <n-form-item label="字号">
            <n-input-number v-model:value="terminal.FontSize" :show-button="false">
              <template #suffix>
                px
              </template>
            </n-input-number>
          </n-form-item>

          <n-checkbox
              v-model:checked="terminal.LineNumber"
          >
            显示行号
          </n-checkbox>

          <n-checkbox
              v-model:checked="terminal.JumpTarget"
          >
            超链接跳转
          </n-checkbox>

        </n-form>

      </n-tab-pane>

      <n-tab-pane name="internal" tab="网络配置">
        <n-checkbox
            v-model:checked="network.Speed"
        >启用海外代理加速
        </n-checkbox>
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
            <Reset/>
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
              <save/>
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
              <Cancel/>
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

.funcButt {
  float: right;
}

.apply {
  margin: 0 10px;
}


</style>