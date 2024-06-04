import { createRouter, createWebHistory } from 'vue-router';

import Connection from "../pages/Connection.vue";
import ExecCommand from "../pages/ExecCommand.vue";

const routes = [
    {
        path: '/',
        name: 'Connection',
        component: Connection
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