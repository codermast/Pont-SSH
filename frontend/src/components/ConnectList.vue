<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Add, ArrowDown, Pencil } from '@vicons/ionicons5'
import NewConnect from "../dialogs/NewConnect.vue";
import { useDialogStore } from "../stores/dialogStore";
import { NIcon, useMessage } from 'naive-ui'
import Search from "./Search.vue";
import { DeleteServer, GetServerList, ServerConnection } from "../../wailsjs/go/service/Connection";
import emitter from '../utils/emitter';
import { useRouter } from "vue-router";
import EditConnect from "../dialogs/EditConnect.vue";
import { entity } from "../../wailsjs/go/models";
import DeleteIcon from "../icons/DeleteIcon.vue";

const router = useRouter()
const message = useMessage();

const dialogStore = useDialogStore();

// 新增连接弹窗
function newConnectVisible() {
  dialogStore.newConnectDialogVisible = true;
}

// 服务器列表
let serverList = ref<entity.SSHConfig[]>()

onMounted(() => {
  getServerList();
})

emitter.on('reloadServerList', getServerList)

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
function editServer(serverConfig: entity.SSHConfig) {
  // 发送数据
  emitter.emit('sendEditSSHConfig', serverConfig)
  // 展示编辑模态框，并将数据进行传递
  dialogStore.editConnectDialogVisible = true;
}

// 删除服务器连接
function deleteServer(sshConfig: entity.SSHConfig) {

  console.log("sshConfig", sshConfig)

  DeleteServer(sshConfig.id).then((result) => {
    if (result.code == 200) {
      message.success(result.msg);
    } else {
      message.error(result.msg);
    }
  })
}

// 更新搜索结果
function changeServerList(sshConfig: entity.SSHConfig[]) {
  serverList.value = sshConfig;
}

// 获取服务器操作系统Logo
function getOSLogo(osType: number): string {
  let logoUrl = "/assets/"
  switch (osType) {
    case 1:
      logoUrl += "ubuntu"
      break;
    case 2:
      logoUrl += "centos"
      break;
    case 3:
      logoUrl += "debian"
      break;
    default:
      logoUrl += "linux"
  }
  return logoUrl + "-logo.png";
}
</script>

<template>
  <div class="server-list">
    <n-grid :cols="1">
      <n-gi :span="1">
        <Search :change-server-list="changeServerList"></Search>
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
              <n-flex>
                <n-image
                    :src="getOSLogo(server.type)"
                    width="140px"
                    height="100%"
                    style="margin-left: -20px"
                >
                </n-image>
                <div>
                  <p>{{ server.server }}</p>
                  <p>{{ server.username }}</p>
                </div>
                <!-- 编辑按钮 -->
                <n-tooltip trigger="hover" placement="top-start">
                  <template #trigger>
                    <n-button
                        round
                        type="success"
                        style="float: right;
                                position: absolute;
                                bottom: 20px;
                                right: 20px;"
                        v-show="server.edit"
                        @click.stop="editServer(server)"
                    >
                      <template #icon>
                        <n-icon size="20">
                          <pencil></pencil>
                        </n-icon>
                      </template>

                    </n-button>
                  </template>
                  编辑连接
                </n-tooltip>

                <!-- 删除按钮 -->
                <n-tooltip trigger="hover" placement="top-start">
                  <template #trigger>
                    <n-button
                        round
                        type="error"
                        style="float: right;
                                position: absolute;
                                bottom: 60px;
                                right: 20px;"
                        v-show="server.edit"
                        @click.stop="deleteServer(server)"
                    >
                      <template #icon>
                        <n-icon size="20">
                          <delete-icon></delete-icon>
                        </n-icon>
                      </template>

                    </n-button>
                  </template>
                  删除连接
                </n-tooltip>
              </n-flex>

            </n-card>
          </n-gi>

        </n-grid>

      </n-scrollbar>
    </n-config-provider>
    <NewConnect></NewConnect>
    <!-- 编辑连接模态框 -->
    <EditConnect></EditConnect>

  </div>

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
  height: 150px;
  width: 95%;
  margin: 10px auto 10px;
}

</style>