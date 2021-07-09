import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.use(VueRouter)

export default new VueRouter({
    mode: 'history',
    routes: [
        {
            name: 'index',
            path: '/',
            component: () => import('@/views/index')
        },
        {
            name: 'register',
            path: '/register',
            component: () => import('@/views/register')
        },
        {
            name: 'login',
            path: '/login',
            component: () => import('@/views/login')
        },
        {
            name: 'detail',
            path: '/detail/:pid',
            component: () => import('@/views/detail')
        },
        {
            name: 'board',
            path: '/board/:bid',
            component: () => import('@/views/board')
        }
    ]
})