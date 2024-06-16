import { createRouter, createWebHistory } from 'vue-router';
import Home from '../pages/Home.vue';
import Terminal from "../components/Terminal.vue";

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
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