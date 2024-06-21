<script setup lang="ts">
import { useRouter } from "vue-router";
import { BrowserOpenURL } from "../../wailsjs/runtime";
import { useDialogStore } from "../stores/dialogStore";
import AboutAuthor from "../dialogs/AboutAuthor.vue";

const dialogStore = useDialogStore();
const router = useRouter();

// 连接列表
function connectListClick() {
  router.push({name: 'ConnectList'});
}

// 日志管理
function logListClick() {
  router.push({name: 'LogList'});
}

// 系统设置
function systemConfigClick() {
  router.push({name: 'SystemConfig'});
}

// helpInfo
function helpInfoClick() {
  // TODO: Update HelpInfo
  BrowserOpenURL("https://www.github.com/CoderMast/Pont-SSH");
}

// AboutAuthor
function aboutAuthorClick() {
  dialogStore.aboutAuthorDialogVisible = true;
}

// Contribute
function contributeClick() {
  // TODO: Update Contribute Introduce Info
  BrowserOpenURL("https://www.github.com/CoderMast/Pont-SSH");
}

// Github
function githubClick() {
  BrowserOpenURL("https://www.github.com/CoderMast/Pont-SSH");
}

// 跳转到首页
function backToIndex() {
  router.push({name: 'Index'});
}
</script>

<template>

  <n-grid :cols="1">
    <n-gi :span="1" class="logo">
      <n-grid :cols="3">
        <n-gi :span="3">
          <n-image preview-disabled
                   src="https://www.codermast.com/logo.png"
                   width="160vw"
                   alt="Logo"
                   @click="backToIndex"
          />
        </n-gi>
      </n-grid>

    </n-gi>

    <n-gi :span="1">
      <div class="sidebar">
        <n-scrollbar style="max-height: 530px">
          <n-list hoverable clickable bordered>
            <n-list-item @click="connectListClick">
              <n-thing title="连接管理" content-style="margin-top: 10px;">
                <template #description>
                  <n-space size="small" style="margin-top: 4px">
                    管理你的 SSH 连接
                  </n-space>
                </template>
              </n-thing>
            </n-list-item>
            <n-list-item @click="logListClick">
              <n-thing title="日志管理" content-style="margin-top: 10px;">
                <template #description>
                  <n-space size="small" style="margin-top: 4px">
                    连接日志管理
                  </n-space>
                </template>
              </n-thing>
            </n-list-item>
            <n-list-item @click="systemConfigClick">
              <n-thing title="系统设置" content-style="margin-top: 10px;">
                <template #description>
                  <n-space size="small" style="margin-top: 4px">
                    自定义系统设置
                  </n-space>
                </template>
              </n-thing>
            </n-list-item>
            <n-list-item @click="helpInfoClick">
              <n-thing title="帮助文档" content-style="margin-top: 10px;">
                <template #description>
                  <n-space size="small" style="margin-top: 4px">
                    查看帮助文档
                  </n-space>
                </template>
              </n-thing>
            </n-list-item>
            <n-list-item @click="aboutAuthorClick">
              <n-thing title="关于作者" content-style="margin-top: 10px;">
                <template #description>
                  <n-space size="small" style="margin-top: 4px">
                    访问 Github 仓库
                  </n-space>
                </template>
              </n-thing>
            </n-list-item>
            <n-list-item @click="contributeClick">
              <n-thing title="参与贡献" content-style="margin-top: 10px;">
                <template #description>
                  <n-space size="small" style="margin-top: 4px">
                    参与本项目的开发
                  </n-space>
                </template>
              </n-thing>
            </n-list-item>
            <n-list-item @click="githubClick">
              <n-thing title="Github" content-style="margin-top: 10px;">
                <template #description>
                  <n-space size="small" style="margin-top: 4px">
                    访问 Github 仓库
                  </n-space>
                </template>
              </n-thing>
            </n-list-item>
          </n-list>
        </n-scrollbar>
      </div>
    </n-gi>
  </n-grid>

  <!-- 关于作者模态框 -->
  <AboutAuthor></AboutAuthor>
</template>

<style scoped>

.sidebar {
  width: 25vw;
  position: fixed; /* 固定定位 */
  left: 0; /* 紧贴左侧 */
  bottom: 0; /* 紧贴底部 */
}


.logo {
  display: block;
  margin: auto;
}
</style>