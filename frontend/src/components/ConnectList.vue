<script setup lang="ts">
import { h, ref } from 'vue'
import { Add, ArrowDown, Pencil } from '@vicons/ionicons5'
import NewConnect from "../dialogs/NewConnect.vue";
import { useDialogStore } from "../stores/dialogStore";
import { NIcon } from 'naive-ui'
import Search from "./Search.vue";

const dialogStore = useDialogStore();

// 新增连接弹窗
function newConnectVisible() {
  dialogStore.newConnectDialogVisible = true;
}

function renderIcon() {
  return h(NIcon, null, {
    default: () => h(Add)
  })
}

let edit = ref(false)

</script>

<template>
  <div class="server-list">
    <n-grid :cols="1">
      <n-gi :span="1">
        <Search></Search>
      </n-gi>
    </n-grid>

    <n-grid :cols="4">
      <n-gi :span="1">
        <n-button
            class="submitButton"
            type="info"
            strong
            :render-icon="renderIcon"
            @click="newConnectVisible"
        >
          添加新连接
        </n-button>

      </n-gi>
      <n-gi :span="1">

        <n-button class="submitButton"
                  icon-placement="left"
                  strong
                  @click="newConnectVisible"
                  type="primary"
        >
          <template #icon>
            <n-icon>
              <arrow-down/>
            </n-icon>
          </template>
          导入连接项
        </n-button>
      </n-gi>
      <n-gi :span="1">
        <n-button
            class="submitButton"
            strong
            type="warning"
            :render-icon="renderIcon"
            @click="newConnectVisible"
        >
          添加新连接
        </n-button>
      </n-gi>
      <n-gi :span="1">
        <n-button
            class="submitButton"
            strong
            type="error"
            :render-icon="renderIcon"
            @click="newConnectVisible"
        >
          添加新连接
        </n-button>
      </n-gi>
    </n-grid>
    <n-config-provider
        :theme-overrides="{
                      Scrollbar: {
                          width: '8px',
                          railInsetHorizontal: '4px 4px 4px auto',
                          borderRadius: 0
                      }
                  }"
    >
      <n-scrollbar style="max-height: 80vh">
        <n-grid :cols="2">
          <n-gi :span="1" v-for="n in 15" :key="n">
            <n-card
                title="卡片"
                class="server-card"
                size="small"
                hoverable
                @mouseover="edit = true"
                @mouseout="edit = false">
              <p style="line-height: 20px;vertical-align: center">卡片内容
                <n-button round dashed type="tertiary" style="float: right;" v-show="edit">
                  <template #icon>
                    <n-icon size="20">
                      <pencil></pencil>
                    </n-icon>
                  </template>

                </n-button>
              </p>
            </n-card>
          </n-gi>

        </n-grid>

      </n-scrollbar>
    </n-config-provider>
  </div>

  <!-- 创建连接弹窗 -->
  <NewConnect></NewConnect>
</template>

<style scoped>


.submitButton {
  height: 40px;
  width: 90%;
  border-radius: 0;
  margin: 10px;
}

.submitButton:first-child {
  margin-left: 10px;
}

.submitButton:last-child {
  margin-right: 0;
}


.server-list {
  height: 100vh;
}

.server-card {
  width: 90%;
  margin: 10px auto 10px;
}

</style>