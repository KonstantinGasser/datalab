<template>
   <div>
        <h1>Notifications</h1>
        <div class="view_component">
            <!-- Hello NotificationCenter
            {{notifications}} -->
            <div class="notify-table">
                <div v-for="item in notifications" :key="item.id" class="notify-row">
                    <div class="title">{{item.title}}</div>
                    <div class="content">

                    </div>
                </div>
            </div>
        </div>
   </div>
</template>

<script>

export default {
  name: 'NotificationCenter',
  data() {
    return {
    };
  },
  components: {
  },
  computed: {
      notifications() {
          return this.$store.state.notifications;
      },
  },
  methods: {
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.app-view {
    height: 100%;
}

.notify-table {
    display: grid;
    align-content: flex-start;
    justify-content: flex-start;
    width: 100%;
    padding: 15px 25px;
}

.notify-row {
    width: 400px;
    padding: 15px 25px;
    border-radius: 8px;
    color: var(--h-color);
    background: var(--sub-border);
    display: flex;
    align-items: center;
    justify-content: space-between;
}
</style>
