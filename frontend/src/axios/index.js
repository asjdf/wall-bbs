import axios from "axios";
import router from "@/router"
axios.defaults.withCredentials = true;

const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_URL,
    credentials: 'include',
    withCredentials: true,
    timeout: 5000
});

// http request 拦截器
service.interceptors.request.use(
    config => {
        if (localStorage.token) { //判断token是否存在
            config.headers.Authorization = "token " + localStorage.token;  //将token设置成请求头
        }
        return config;
    },
    err => {
        return Promise.reject(err);
    }
);

// http response 拦截器
service.interceptors.response.use(
    response => {
        if (response.data.code === 40001) {
            localStorage.removeItem('token');
            localStorage.removeItem('hasToken');
            localStorage.removeItem('right');
            localStorage.removeItem('uid');
            router.replace('/login');
            console.log("token过期");
        }
        return response;
    },
    error => {
        return Promise.reject(error);
    }
);

export default service;