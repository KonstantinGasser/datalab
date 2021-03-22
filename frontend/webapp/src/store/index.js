import { createStore } from 'vuex';

export default createStore({
  state: {
    activeTab: 'view_app',
  },
  mutations: {
    setActiveTab(state, status) {
      state.activeTab = status;
    },
  },
  actions: {
  },
  modules: {
  },
});
