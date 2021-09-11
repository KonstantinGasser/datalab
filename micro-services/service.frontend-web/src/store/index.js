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
    online: [],
    sync_app: {
      uuid: null,
      sync: false,
    },
  },
  mutations: {
    INIT_LOAD(state, event) {
      event.notifications?.forEach(item => {
        state.notifications.push(item);
      })
    },
    APP_INVITE(state, event) {
      state.notifications.push(event);
    },
    APP_INVITE_REMINDER(state, event) {
      state.notifications.push(event)
    },
    IS_ONLINE(state, event) {
      switch (event?.event) {
        case 3: // new user came online
          console.log("Event: ",event)
          state.online.push(event?.receiver_uuid)
          console.log("State: ",state.online)
          break;
        case 4: // user goes offline
          state.online = state.online.filter(uuid => uuid != event?.receiver_uuid)
          break;
      }
      console.log("State: ",state.online)
    },
    POP_NOTIFICATION(state, event) {
      state.notifications = state.notifications.filter(item => !(item.timestamp === event.timestamp))
    },
    SYNC_APP(state, event) {
      state.sync_app.uuid = event?.value?.app_uuid
      state.sync_app.sync = event?.value?.sync
    },
    UNSYNC_APP(state) {
      state.sync_app.uuid = null
      state.sync_app.sync = false
    },
    PURGE_CONN(state) {
      state.notifications = []
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
