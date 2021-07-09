import axios from "axios";

axios.defaults.withCredentials = true;

const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_URL,
    credentials: 'include',
    withCredentials: true,
    timeout: 5000
});

export default service;