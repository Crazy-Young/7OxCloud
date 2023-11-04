import { UserLogin, UserInfo, UserUpdate, UserRegister, UserPasswordReset } from "@/api/user"
import { UserFollowList, UserFanList } from "@/api/user"
import { setToken, getToken, removeToken } from '@/utils/token'
import Vue from 'vue'

const state = {
    userInfo: {},
    followList: [],
    fanList: [],
}

const mutations = {
    SET_USER_INFO(state, userInfo) {
        state.userInfo = userInfo
    },
    SET_TOKEN(state, token) {
        setToken(token)
    },
    REMOVE_TOKEN(state) {
        removeToken()
    },
    REMOVE_USER_INFO(state) {
        state.userInfo = {}
    },
    SET_USER_FOLLOW(state, followList) {
        state.followList = followList
    },
    SET_USER_FAN(state, fanList) {
        state.fanList = fanList
    }
}

const actions = {
    UserLogin({ dispatch, commit }, data) {
        return new Promise((resolve, reject) => {
            UserLogin(data).then(res => {
                if (res.status === 200) {
                    commit('SET_TOKEN', res.data.token)
                    Vue.prototype.$message.success(res.data.msg)
                    dispatch('UserInfo')
                    resolve(res)
                } else {
                    Vue.prototype.$message.error(res.data.msg)
                    reject(res)
                }
            }).catch(error => {
                reject(error)
                throw new Error(error)
            })
        })
    },
    UserLogout({ commit }) {
        commit('REMOVE_TOKEN')
        commit('REMOVE_USER_INFO')
    },
    UserInfo({ dispatch, commit }) {
        return getToken() && new Promise((resolve, reject) => {
            UserInfo().then(res => {
                if (res.status === 200) {
                    commit('SET_USER_INFO', res.data.data)
                    dispatch('UserFollowList')
                    dispatch('UserFanList')
                    resolve(res)
                } else {
                    Vue.prototype.$message.error(res.data.msg)
                    reject(res)
                }
            }).catch(error => {
                reject(error)
                throw new Error(error)
            })
        })
    },
    UserUpdate({ dispatch }, data) {
        return new Promise((resolve, reject) => {
            UserUpdate(data).then(res => {
                if (res.status === 200) {
                    Vue.prototype.$message.success(res.data.msg)
                    dispatch('UserInfo')
                    resolve(res)
                } else {
                    Vue.prototype.$message.error(res.data.msg)
                    reject(res)
                }
            }).catch(error => {
                reject(error)
                throw new Error(error)
            })
        })
    },
    UserRegister({ commit, dispatch }, data) {
        return new Promise((resolve, reject) => {
            UserRegister(data).then(res => {
                if (res.status === 200) {
                    commit('SET_TOKEN', res.data.token)
                    Vue.prototype.$message.success(res.data.msg)
                    dispatch('UserInfo')
                    resolve(res)
                } else {
                    Vue.prototype.$message.error(res.data.msg)
                    reject(res)
                }
            }).catch(error => {
                reject(error)
                throw new Error(error)
            })
        })
    },
    UserPasswordReset({ commit }, data) {
        return new Promise((resolve, reject) => {
            UserPasswordReset(data).then(res => {
                if (res.status === 200) {
                    Vue.prototype.$message.success(res.data.msg)
                    resolve(res)
                } else {
                    Vue.prototype.$message.error(res.data.msg)
                    reject(res)
                }
            }).catch(error => {
                reject(error)
                throw new Error(error)
            })
        })
    },
    UserFollowList({ commit }) {
        return new Promise((resolve, reject) => {
            UserFollowList().then(res => {
                if (res.status === 200) {
                    commit('SET_USER_FOLLOW', res.data.data.followings || [])
                    resolve(res)
                } else {
                    Vue.prototype.$message.error(res.data.msg)
                    reject(res)
                }
            }).catch(error => {
                reject(error)
                throw new Error(error)
            })
        })
    },
    UserFanList({ commit }) {
        return new Promise((resolve, reject) => {
            UserFanList().then(res => {
                if (res.status === 200) {
                    commit('SET_USER_FAN', res.data.data.fans || [])
                    resolve(res)
                } else {
                    Vue.prototype.$message.error(res.data.msg)
                    reject(res)
                }
            }).catch(error => {
                reject(error)
                throw new Error(error)
            })
        })
    }
}

const getters = {
    followCount: state => state.followList.length,
    fanCount: state => state.fanList.length
}

export default {
    state,
    mutations,
    actions,
    getters
}