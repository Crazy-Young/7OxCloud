import axios from "axios";
import { getToken, removeToken, setToken } from '@/utils/token'
import Vue from "vue";

// 设置取消请求令牌
let cancelToken = axios.CancelToken.source()

// 创建axios实例
const request = axios.create({
    baseURL: process.env.VUE_APP_API_BASE_URL,
    cancelToken: cancelToken.token
})

// 添加请求拦截器，在发送请求之前，携带token
request.interceptors.request.use((config) => {
    const token = getToken();
    config.headers.authorization = `${token}`
    return config
}, error => {
    console.log(error)
});

// 添加响应拦截器，在接收响应之前，重新设置token，避免过期
request.interceptors.response.use((response) => {
    const { authorization } = response.headers;
    const token = authorization || response.data?.token
    token && setToken(token);
    return response;
}, (error) => {
    // 处理 401 Unauthorized 情况，重新登录
    if (error.response?.status === 401) {
        // 移除token
        removeToken();
        // 派发退出事件，清除当前用户数据
        Vue.prototype.$store?.dispatch('UserLogout');
        // 唤起登录框
        Vue.prototype.$bus.$emit('toggleLogin');
        return Promise.reject(error.response.data.msg);
    }
    return error.response
});

export default request;

// 取消所有axios请求
export const CancelToken = () => {
    // 取消请求
    cancelToken.cancel()
    // 创建新的请求
    cancelToken = axios.CancelToken.source()
}