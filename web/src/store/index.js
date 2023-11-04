import user from "./user";
import category from "./category";
import Vuex from "vuex";
import Vue from "vue";
Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        user,
        category
    }
})