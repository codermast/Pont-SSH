<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { FlashOutline, Reload } from "@vicons/ionicons5";
import DeleteIcon from "../icons/DeleteIcon.vue";
import { GetServerList } from "../../wailsjs/go/service/Connection";
import { GetLogInfoList } from "../../wailsjs/go/service/LogInfo";
import { useMessage } from 'naive-ui'
import { entity } from '../../wailsjs/go/models';

const message = useMessage();

let options = [
  {
    label: '全部',
    value: 'all'
  }
]

let all = ref("all")

let keyword = ref(null)

let logColumns = [
  {
    title: '服务器',
    key: 'server_id'
  },
  {
    title: '执行命令',
    key: 'command'
  },
  {
    title: '执行时间',
    key: 'exec_date'
  },
  {
    title: '执行结果',
    key: 'time'
  }
]

let logData = ref<entity.LogInfo[]>([

])

onMounted(() => {
  // 获取服务器列表
  GetServerList().then(result => {
    console.log(result)
    if (result.code == 200) {
      let data = result.data

      if (data != null) {
        for (let i = 0; i < data.length; i ++) {
          options.push({
            label: data[i].name,
            value: data[i].id
          })
        }
      }
    }
  })

  // 查询所有服务器日志
  getLogInfoList(options[0].value)
})

// 查询指定服务器日志
function getLogInfoList(serverId: string) {
  // 1. 先进行数据初始化
  logData.value = []

  // 2. 然后添加数据
  GetLogInfoList(serverId).then((result) => {
    console.log(result)
    if (result.code == 200) {
      let data = result.data

      if (data != null) {
        // 这里拼接日志信息
        for (let i = 0; i < data.length; i ++) {
          logData.value.push(data[i])
        }
      }
    } else {
      message.error("获取日志失败！");
    }

    console.log("logData", logData)
  })
}
</script>

<template>
  <div class="log-main">
    <n-grid :cols="1" :y-gap="20">
      <n-gi :span="1">
        <h2>日志管理</h2>
      </n-gi>
      <n-gi :span="1">
        <n-grid :cols="4">
          <n-gi :span="1">
            <span>筛选服务器</span>
            <n-select
                v-model:value="all"
                :options="options"
                style="width: 80%"
                @change="getLogInfoList"
            />
          </n-gi>
          <n-gi :span="3">
            <div style="display: flex;flex-direction: column">
              <span>关键字</span>
              <div>
                <n-input v-model:value="keyword"
                         type="text"
                         placeholder="搜索"
                         style="width: 70%">
                  <template #prefix>
                    <n-icon :component="FlashOutline"/>
                  </template>
                </n-input>
                <n-button style="margin-left: 10px">
                  <template #icon>
                    <n-icon>
                      <reload/>
                    </n-icon>
                  </template>
                </n-button>

                <n-button style="margin-left: 10px">
                  <template #icon>
                    <n-icon>
                      <delete-icon/>
                    </n-icon>
                  </template>
                </n-button>
              </div>
            </div>
          </n-gi>
        </n-grid>
      </n-gi>

      <n-gi :span="1">
        <n-data-table
            class="logTable"
            :columns="logColumns"
            :data="logData"
        />
      </n-gi>
    </n-grid>
  </div>
</template>

<style scoped>
.log-main {
  width: 90%;
  margin: auto;
}

.logTable {
  height: 530px;
}
</style>