<template>
      <div class="menu_item" :class="{'active': isActive}">
        <div class="menu_text">
            {{ item }}
        </div>
        <span v-if="bubble && notifications > 0" class="bubble"><i class="icon icon-bell"></i></span>
    </div>
</template>

<script>

export default {
  name: 'MenuItem',
  data() {
    return {
    };
  },
  props: {
      item: String,
      tabName: String,
      bubble: Boolean,
  },
  methods: {
  },
  computed: {
      isActive() {
          return this.$props.tabName === this.$store.state.activeTab;
      },
      notifications() {
          return this.$store.state.notifications.length;
      },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.menu_item {
    padding: 5px;
    margin: 15px 5px;
    display: flex;
    align-items: center;
    justify-content: flex-start;
    border-radius: 8px;
    width: 135px;
    height: 35px;
    transition: all .2s ease-in-out;
}

.menu_item:hover {
    cursor: pointer;
    width: 100%;
    background: var(--sub-border);
}
.menu_text {
    color: var(--h-color);
    font-size: 20px;
    text-align: left;
    position: relative;
}
.menu_item:hover .menu_text {
    background: var(--gradient-green);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}

.menu_item.active {
    background: var(--sub-border);
    /* border: 1px solid #30363D; */
}
.active .menu_text {
    background: var(--gradient-green);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}

.bubble {
    width: 25px;
    height: 25px;
    position: absolute;
    margin-left: 125px;
    margin-top: -25px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: red;
    border-radius: 50%;
    color: white
}
</style>
