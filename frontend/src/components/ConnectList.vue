<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Add, ArrowDown, Pencil } from '@vicons/ionicons5'
import NewConnect from "../dialogs/NewConnect.vue";
import { useDialogStore } from "../stores/dialogStore";
import { NIcon, useMessage } from 'naive-ui'
import Search from "./Search.vue";
import { GetServerList, ServerConnection } from "../../wailsjs/go/service/Connection";

import { useRouter } from "vue-router";
import EditConnect from "../dialogs/EditConnect.vue";

const router = useRouter()
const message = useMessage();

const dialogStore = useDialogStore();

// 新增连接弹窗
function newConnectVisible() {
  dialogStore.newConnectDialogVisible = true;
}

// 服务器列表
let serverList = ref()

onMounted(() => {
  getServerList();
})

// 获取服务器列表
function getServerList() {
  GetServerList().then(result => {
    console.log(result)

    if (result.code == 200) {
      serverList.value = result.data
    } else {
      message.error(result.msg)
    }
  })
}

// 连接到服务器
function connectServer(sshConfig: any) {
  console.log("sshConfig", sshConfig)
  console.log("sshConfig.value", sshConfig.value)
  ServerConnection(sshConfig).then((result) => {
    // 1. 判断是否成功
    if (result.code == 200) {
      message.success(result.msg);
      router.push({name: "Cmd"});
    } else {
      message.error(result.msg);
    }
  })
}

// 编辑服务器连接
function editServer(sshConfig: any) {
  // 展示编辑模态框，并将数据进行传递
  dialogStore.editConnectDialogVisible = true;
}
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
            @click="newConnectVisible"
        >
          <template #icon>
            <n-icon>
              <add/>
            </n-icon>
          </template>
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
            @click="newConnectVisible"
        >
          <template #icon>
            <n-icon>
              <add/>
            </n-icon>
          </template>
          添加新连接
        </n-button>
      </n-gi>
      <n-gi :span="1">
        <n-button
            class="submitButton"
            strong
            type="error"
            @click="newConnectVisible"
        >
          <template #icon>
            <n-icon>
              <add/>
            </n-icon>
          </template>
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
          <n-gi :span="1" v-for="server in serverList">
            <n-card
                :title="server.name"
                class="server-card"
                size="small"
                hoverable
                @mouseover="server.edit = true"
                @mouseout="server.edit = false"
                @click="connectServer(server)"
            >
              <p style="line-height: 20px;vertical-align: center">{{ server.server + "," + server.username }}
                <n-button
                    round
                    dashed
                    type="tertiary"
                    style="float: right;"
                    v-show="server.edit"
                    @click.stop="editServer(server)"
                >
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

  <!-- 创建连接模态框 -->
  <NewConnect></NewConnect>

  <!-- 编辑连接模态框 -->
  <EditConnect
  >
  </EditConnect>

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