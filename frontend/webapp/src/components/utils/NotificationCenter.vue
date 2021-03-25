<template>
    <div class="notification-center">
        <div><span class="icon icon-bell green big hover" @click="show_notification = !show_notification"></span></div>
        <div v-if="showNotifications" class="notification-popover">
          <div class="notification-handle"></div>
            some notification
        </div>
      </div>
</template>

<script>

    export default {
        name: 'NotificationCenter',
        components: {},
        data() {
            return {
                show_notification: true,
            };
        },
          computed: {
            showNotifications() {
            return !this.show_notification;
            },
        },
        methods: {
        },
    };
</script>

<style scoped>
.notification-center {
  width: min-content;
  display: flex;
  align-content: center;
  position: relative;
}
.notification-popover {
  position: absolute;
  width: 250px;
  height: min-content;
  min-height: 150px;
  background: #1E1E1E;
  right: 0;
  top: 30px;
  z-index: 99;
  border-radius: 8px 8px 8px 8px;
  box-shadow: 0px 0px 12px 16px rgba(0, 0, 0, 0.2);
  border: 1px solid #30363D;
}
.notification-handle {
  width: 12px;
  height: 12px;
  background: #1E1E1E;
  transform: rotate(45deg);
  border-radius: 3px;
  margin: -5px 0 0 233px;
  border-top: 1px solid #30363D;
  border-left: 1px solid #30363D;
}
</style>
