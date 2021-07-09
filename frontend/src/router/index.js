import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.use(VueRouter)

export default new VueRouter({
    mode: 'history',
    routes: [
        {
            name: '首页',
            path: '/',
            component: () => import('@/views/index')
        },
        {
            name: '注册',
            path: '/register',
            component: () => import('@/views/register')
        },
        {
            name: '登录',
            path: '/login',
            component: () => import('@/views/login')
        }
    ]
})