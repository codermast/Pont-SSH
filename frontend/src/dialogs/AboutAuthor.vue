<script setup lang="ts">
import { useDialogStore} from "../stores/dialogStore";
import { BrowserOpenURL } from "../../wailsjs/runtime";

const dialogStore = useDialogStore();

function authorInfoClick() {
  BrowserOpenURL("https://www.github.com/CoderMast/Pont-SSH");
}

function authorInfoFocusout(){
  dialogStore.aboutAuthorDialogVisible = false;
}
</script>

<template>
  <n-modal v-model:show="dialogStore.aboutAuthorDialogVisible" transform-origin="center">
    <n-image
        width="500vw"
        src="../assets/friend-wechat-logo.jpg"
        preview-disabled
        @click="authorInfoClick"
        @focusout="authorInfoFocusout"
    />
  </n-modal>
</template>

<style scoped>

</style>