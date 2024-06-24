import { ref } from "vue";

// 操作系统选项列表
export let osTypeOptions = ref([
	{
		label: 'Ubuntu',
		value: 1,

	},
	{
		label: 'Centos',
		value: 2,
	},
	{
		label: 'Debian',
		value: 3,
	}, {
		label: '其他',
		value: 0
	}
]);
