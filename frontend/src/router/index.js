import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.use(VueRouter)

const router = new VueRouter({
    mode: 'history',
    routes: [
        {
            name: 'index',
            path: '/',
            component: () => import('@/views/index'),
            meta:{
                title: "首页"
            }
        },
        {
            name: 'register',
            path: '/register',
            component: () => import('@/views/register'),
            meta:{
                title: "注册"
            }
        },
        {
            name: 'login',
            path: '/login',
            component: () => import('@/views/login'),
            meta:{
                title: "登录"
            }
        },
        {
            name: 'detail',
            path: '/detail/:pid',
            component: () => import('@/views/detail'),
            meta:{
                title: "详情"
            }
        },
        {
            name: 'board',
            path: '/board/:bid',
            component: () => import('@/views/board'),
            meta:{
                title: "板块"
            }
        }
    ]
})

router.beforeEach((to,from,next)=>{//beforeEach是router的钩子函数，在进入路由前执行
    if(to.meta.title){//判断是否有标题
        document.title = to.meta.title
    }
    next()  //执行进入路由，如果不写就不会进入目标页
})

export default router