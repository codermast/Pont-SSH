// Dialog 控制
import { defineStore } from 'pinia';
import { ref } from 'vue';

export const useDialogStore = defineStore('dialogStore', () => {
	// 创建连接弹窗
	let newConnectDialogVisible = ref(false);

	// 编辑连接弹窗
	let editConnectDialogVisible = ref(false);

	// 关于作者弹窗
	let aboutAuthorDialogVisible = ref(false);

	// 连接异常弹出框
	let errorConnectDialogVisible = ref(false);

	return {newConnectDialogVisible, editConnectDialogVisible, aboutAuthorDialogVisible, errorConnectDialogVisible}
})
