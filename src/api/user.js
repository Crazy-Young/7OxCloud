import request from './request'

export const UserLogin = (data) => request.post('/user/login', data)
export const UserRegister = (data) => request.post('/user/register', data)
export const UserInfo = (uid) => request.get('/user/get_user' + (uid ? `?uid=${uid}` : ''))
export const UserUpdate = (data) => request.put('/user/update_user', data)
export const UserPasswordReset = (data) => request.put('/user/reset_password', data)