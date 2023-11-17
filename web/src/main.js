import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { createRouter, createWebHistory } from 'vue-router'

createApp(App).use(createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: () => import('./components/AttendanceInput.vue') },
        { path: '/attendance', component: () => import('./components/AttendanceList.vue') }
]})).mount('#app')