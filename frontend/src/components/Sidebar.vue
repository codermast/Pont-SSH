<script setup lang="ts">
import { useRouter } from "vue-router";
import { BrowserOpenURL } from "../../wailsjs/runtime";
const router = useRouter();

// 连接列表
function connectListClick(){
    router.push({ name : 'ConnectList' });
}

// 日志管理
function logListClick(){
  router.push({ name : 'LogList' });
}

// 系统设置
function systemConfigClick(){
  router.push({ name : 'SystemConfig' });
}

// Github
function githubClick(){
  BrowserOpenURL("https://www.github.com/CoderMast/Pont-SSH");
}
</script>

<template>
  <div class="left">
    <n-grid :cols="1">

      <n-gi :span="1" class="logo">
        <n-grid :cols="3">
          <n-gi :span="3">
            <n-image preview-disabled
                     src="https://www.codermast.com/logo.png"
                     width="160vw"
                     alt="Logo"/>
          </n-gi>
        </n-grid>

      </n-gi>

      <n-gi :span="1">
        <n-scrollbar style="max-height: 530px">
          <n-list hoverable clickable bordered >
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
                  管理你的 SSH 连接
                </n-space>
              </template>
            </n-thing>
          </n-list-item>
          <n-list-item @click="systemConfigClick">
            <n-thing title="系统设置" content-style="margin-top: 10px;">
              <template #description>
                <n-space size="small" style="margin-top: 4px">
                  管理你的 SSH 连接
                </n-space>
              </template>
            </n-thing>
          </n-list-item>
          <n-list-item>
            <n-thing title="帮助文档" content-style="margin-top: 10px;">
              <template #description>
                <n-space size="small" style="margin-top: 4px">
                  查看帮助文档
                </n-space>
              </template>
            </n-thing>
          </n-list-item>
          <n-list-item>
            <n-thing title="关于作者" content-style="margin-top: 10px;">
              <template #description>
                <n-space size="small" style="margin-top: 4px">
                  访问 Github 仓库
                </n-space>
              </template>
            </n-thing>
          </n-list-item>
          <n-list-item >
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
      </n-gi>
    </n-grid>

  </div>

</template>

<style scoped>
.left {
  height: 100vh;
}


.logo {
  display: block;
  margin: auto;
}
</style>