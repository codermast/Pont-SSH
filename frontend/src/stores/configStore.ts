// 系统配置 store
import { defineStore } from 'pinia';
import { ref } from "vue";

export const useConfigStore = defineStore('configStore', () => {
	// WebSocket Port
	let webSocketPort = ref(54321);

	return { webSocketPort }
})
