<script setup lang="ts">

import { FlashOutline } from "@vicons/ionicons5";

import { ref } from 'vue'
import { useDialogStore } from "../stores/dialogStore";
import { useMessage } from "naive-ui";
import { SearchConnection } from "../../wailsjs/go/service/Connection";

const message = useMessage();

const props = defineProps(['changeServerList'])

// 关键字
let keyword = ref<string>("");

// 关键词搜索
function searchClick() {
  if (keyword.value === "") {
    message.error("请输入关键词！");
  } else {
    SearchConnection(keyword.value).then((result: any) => {
      if (result.code === 200) {
        message.success("搜索成功！");
        props.changeServerList(result.data);
      } else {
        message.error(result.msg);
      }
    })
  }
}

</script>

<template>
  <n-grid :cols="1">
    <n-gi :span="1">
      <n-input-group class="search-input">
        <!--  搜索框  -->
        <n-input
            type="text"
            placeholder="Please enter the keyword that your want to search for ..."
            v-model:value="keyword"
        >
          <template #prefix>
            <n-icon :component="FlashOutline"/>
          </template>
        </n-input>
        <n-button @click="searchClick" class="search-submit">搜索</n-button>
      </n-input-group>

    </n-gi>
  </n-grid>

</template>

<style scoped>

.search-input {
  display: flex;
  margin: 10px auto 10px;
  width: 73vw;
  border-radius: 20px;
}
</style>