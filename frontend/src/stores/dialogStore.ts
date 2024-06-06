// Dialog 控制
import { defineStore } from 'pinia';
import { ref } from 'vue';

export const useDialogStore = defineStore('dialogStore', () => {
    // 创建连接弹窗
    let newConnectDialogVisible = ref(false);

    return { newConnectDialogVisible }
})
