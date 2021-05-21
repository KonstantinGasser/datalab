import { createStore } from 'vuex';
import main from "../main";

export default createStore({
  state: {
    socket: {
      activeTab: 'view_app',
      // Connection Status
      isConnected: false,
      // Message content
      message: "",
    },
    notifications: [],
  },
  mutations: {
    APP_INVITE(state, event) {
      console.log("This is a test: ", event);
      console.log("State: ", state.notifications);
      event["timestamp"] = new Date().getTime();
      state.notifications.push(event);
    },
    POP_NOTIFICATION(state, event) {
      state.notifications = state.notifications.filter(item => !(item.timestamp === event.timestamp))
    },
    // Connection open
    SOCKET_ONOPEN (state, event) {
      Vue.prototype.$socket = event.currentTarget
      state.socket.isConnected = true
    },
    SOCKET_ONCLOSE (state) {
      state.socket.isConnected = false
    },
    SOCKET_ONERROR (state, event) {
      console.error(state, event)
    },
    // default handler called for all methods
    SOCKET_ONMESSAGE (state, message) {
      state.socket.message = message
    },
    // mutations for reconnect methods
    SOCKET_RECONNECT(state, count) {
      console.info(state, count)
    },
    SOCKET_RECONNECT_ERROR(state) {
      state.socket.reconnectError = true;
    },
    setActiveTab(state, status) {
      state.activeTab = status;
    },
  },
  actions: {
  },
  modules: {
  },
});
