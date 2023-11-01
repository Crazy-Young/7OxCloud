import user from "./user";
import Vuex from "vuex";
import Vue from "vue";
Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        user
    }
})