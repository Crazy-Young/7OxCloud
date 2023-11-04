import request from './request'

/**
 * 登录
 * @param {Object} data 登录信息
 * @returns {Promise}
 * @memberof User
 * @example
 * UserLogin({
    mobile: '13900000001',
    password: 'a12345678',
    captcha: '123456',
    captchaId: '123456',
    type: 1
 })
 * @example
 * UserLogin({
     mobile: '13900000001',
     captcha: '123456',
     type: 2
 })
 */
export const UserLogin = (data) => request.post('/user/login', data)

/**
 * 注册
 * @param {Object} data 注册信息
 * @returns {Promise}
 * @memberof User
 * @example
 * UserRegister({
    mobile: '13900000001',
    password: 'a12345678',
    captcha: '123456',
 })
 */
export const UserRegister = (data) => request.post('/user/register', data)

/**
 * 退出登录
 * @returns {Promise}
 * @memberof User
 */
export const UserLogout = () => request.get('/user/logout')

/**
 * 获取用户信息
 * @param {Number} uid 用户ID
 * @returns {Promise}
 * @memberof User
 */
export const UserInfo = (uid) => request.get('/user/get_user' + (uid ? `?uid=${uid}` : ''))

/**
 * 修改用户信息
 * @param {Object} data 修改信息
 * @returns {Promise}
 * @memberof User
 * @example
 * UserUpdate({
     age: 18,
     avatar: 'http://s38ddu7np.hn-bkt.clouddn.com/avatar.jpg',
     gender: 1,
     location: '北京',
     signature: '这个人很懒,什么都没留下',
     username: '张三'
 })
 */
export const UserUpdate = (data) => request.put('/user/update_user', data)

/**
 * 重置密码
 * @param {Object} data 修改信息
 * @returns {Promise}
 * @memberof User
 * @example
 * UserPasswordReset({
     mobile: '13900000001',
     captcha: '123456',
     password: 'b12345678'
 })
 */
export const UserPasswordReset = (data) => request.put('/user/reset_password', data)

/**
 * 关注/取消关注用户
 * @param {Object} data
 * @returns {Promise}
 * @memberof User
 * @example 关注
 * FollowUser({
     type: 1,
     userId: 1
 })
 * @example 取消关注
 * FollowUser({
     type: 2,
     userId: 1
 })
 */
export const FollowUser = (data) => request.post('/social/follow', data)

/**
 * 移除粉丝
 * @param {Number} fanId 粉丝ID
 * @returns {Promise}
 * @memberof User
 */
export const RemoveFan = (fanId) => request.delete('/social/remove_fan' + (fanId ? `?fanId=${fanId}` : ''))

/**
 * 获取关注列表
 * @param {Number} uid 用户ID
 * @returns {Promise}
 * @memberof User
 */
export const UserFollowList = (uid) => request.get('/social/following' + (uid ? `?uid=${uid}` : ''))

/**
 * 获取粉丝列表
 * @param {Number} uid 用户ID
 * @returns {Promise}
 * @memberof User
 */
export const UserFanList = (uid) => request.get('/social/fan' + (uid ? `?uid=${uid}` : ''))

/**
 * 搜索关注用户
 * @param {String} keyword 关键字
 * @returns {Promise}
 * @memberof User
 */
export const SearchFollow = (keyword) => request.get('/social/search_following' + (keyword ? `?keyword=${keyword}` : ''))

/**
 * 搜索粉丝
 * @param {String} keyword 关键字
 * @returns {Promise}
 * @memberof User
 */
export const SearchFan = (keyword) => request.get('/social/search_fan' + (keyword ? `?keyword=${keyword}` : ''))