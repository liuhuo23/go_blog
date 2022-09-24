import { createStore, createLogger } from "vuex";

export default createStore({
    state() {
        return {
            count: 1
        };
    },
    mutations: {
        increment(state) {
            state.count++;
        }
    },
    actions: {
        increment(context) {
            context.commit('increment');
        }
    },
    getters: {
        count(state) {
            return this.state.count;
        }
    }
})