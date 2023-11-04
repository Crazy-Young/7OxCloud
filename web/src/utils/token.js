/**
 * @description: 获取token
 * @returns {string} token
 * @author: 涂国彬
 */
const getToken = () => {
    const token = localStorage.getItem('token')
    return token
}

/**
 * @description: 设置token
 * @param token
 * @author: 涂国彬
 */
const setToken = (token) => {
    localStorage.setItem('token', token)
}

/**
 * @description: 删除token
 * @author: 涂国彬
 */
const removeToken = () => {
    localStorage.removeItem('token')
}

export {
    getToken,
    setToken,
    removeToken
}