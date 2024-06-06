import { createRouter, createWebHistory } from 'vue-router';
import Home from '../pages/Home.vue';
import ExecCommand from "../components/ExecCommand.vue";

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/exec',
        name: 'ExecCommand',
        component: ExecCommand
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;