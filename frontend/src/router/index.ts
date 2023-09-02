import {createRouter, createWebHistory} from 'vue-router'
import HomePage from '../pages/home/index.vue';
import NotePage from '../pages/note/index.vue';

const routes = [
    {
        path: '/',
        component: HomePage,
    },
    {
        path: '/note',
        component: NotePage,
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router
