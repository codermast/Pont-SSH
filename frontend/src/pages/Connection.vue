<script setup lang="ts">
import { service } from '../../wailsjs/go/models'
import { ref } from "vue";
import { CreateConnection } from '../../wailsjs/go/service/Connection'
import { useRouter } from 'vue-router'

const router = useRouter()
let sshConnection = ref(service.SSHConnection.createFrom({
  server: "10.211.55.17",
  port: 22,
  username: "codermast",
  password: "dongpeng@123.",
  keyPath: "null"
}))

function testConnect() {
  CreateConnection(sshConnection.value).then((connection) => {

    // 1. 判断是否成功
    if (connection.code == 200) {
      console.log("成功",connection);
      router.push("/exec")
    } else {
      console.log("失败",connection);
    }
  })
}

function saveConnect() {
  console.log("保存")
}
</script>

<template>
  <n-form
      :label-width="80"
      :model="sshConnection"
  >
    <n-form-item label="服务器IP">
      <n-input v-model:value="sshConnection.server" placeholder="输入服务器IP"/>
    </n-form-item>
    <n-form-item label="端口">
      <n-input v-model:value="sshConnection.port" placeholder="输入端口"/>
    </n-form-item>
    <n-form-item label="用户名">
      <n-input v-model:value="sshConnection.username" placeholder="输入用户名"/>
    </n-form-item>
    <n-form-item label="密码">
      <n-input v-model:value="sshConnection.password" placeholder="输入密码"/>
    </n-form-item>
    <n-form-item>
      <n-button attr-type="button" @click="testConnect">
        验证
      </n-button>

      <n-button attr-type="button" @click="saveConnect">
        保存
      </n-button>
    </n-form-item>
  </n-form>
</template>

<style scoped>

</style>