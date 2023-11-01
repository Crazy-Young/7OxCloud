import axios from "axios";
import { getToken, removeToken, setToken } from '@/utils/token'

const request = axios.create({
    baseURL: process.env.VUE_APP_API_BASE_URL
})

request.interceptors.request.use((config) => {
    const token = getToken();
    config.headers.authorization = `${token}`
    return config
}, error => {
    console.log(error)
});

request.interceptors.response.use((response) => {
    const { authorization } = response.headers;
    const token = authorization || response.data?.token
    token && setToken(token);
    return response;
}, (error) => {
    if (error.response?.status === 401) {
        removeToken();
        return Promise.reject(error.response.data.msg);
    }
    return error.response
});

export default request;