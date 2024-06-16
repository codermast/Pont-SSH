import { createRouter, createWebHistory } from 'vue-router';
import ConnectList from '../components/ConnectList.vue';
import Terminal from "../components/Terminal.vue";

const routes = [
    {
        path: '/',
        name: 'ConnectList',
        component: ConnectList
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