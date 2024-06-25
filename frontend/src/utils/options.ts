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

// 语言选项列表
export let languages = ref([
	{
		label: '系统默认',
		value: "default",
	}, {
		label: '简体中文',
		value: "zh_cn",
	}, {
		label: '繁体中文',
		value: "zh_hk",
	}, {
		label: 'English',
		value: "en",
	}
])

// 字体选项列表
export let fontFamilies = ref([
	{
		label : '系统默认',
		value: "default",
	},{
		label : 'Apple',
		value: "apple",
	}
])