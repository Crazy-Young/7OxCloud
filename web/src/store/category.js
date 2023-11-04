import { VideoCategoryList } from "@/api/video"

const state = {
    categoryList: []
}

const mutations = {
    GET_CATEGORYLIST(state, categoryList) {
        state.categoryList = categoryList
    }
}

const actions = {
    async getCategoryList({ commit }) {
        return new Promise((resolve, reject) => {
            VideoCategoryList().then(res => {
                if (res.status === 200) {
                    const categories = (res.data.data || []).sort((a, b) => b.id - a.id)
                    commit('GET_CATEGORYLIST', categories)
                    resolve(categories)
                } else {
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
    categoryList: state => state.categoryList
}

export default {
    state,
    mutations,
    actions,
    getters
}