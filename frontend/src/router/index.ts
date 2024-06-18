import { createRouter, createWebHistory } from 'vue-router';
import Home from '../components/Home.vue';
import Terminal from "../components/Terminal.vue";
import ConnectList from "../components/ConnectList.vue";
import LogList from "../components/LogList.vue";
import SystemConfig from "../components/SystemConfig.vue";
import Index from "../components/Index.vue";
const routes = [
	{
		path: '/',
		name: 'Home',
		component: Home,
		children: [
			{
				path: '/',
				name: 'Index',
				component: Index,
			},
			{
				path: '/connectList',
				name: 'ConnectList',
				component: ConnectList,
			}, {
				path: '/logList',
				name: 'LogList',
				component: LogList
			},{
				path: '/systemConfig',
				name: 'SystemConfig',
				component: SystemConfig
			}
		]
	},
	{
		path: '/cmd',
		name: 'Cmd',
		component: Terminal
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes
});

export default router;