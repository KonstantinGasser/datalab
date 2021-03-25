<template>
      <div class="menu_item" :class="{'active': isActive}">
        <div class="menu_text">
            {{ item }}
        </div>
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
  },
  methods: {
  },
  computed: {
      isActive() {
          return this.$props.tabName === this.$store.state.activeTab;
      }
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
    background:#0000001c;
    /* border: 1px solid #30363D; */
}
.menu_text {
    color: #fff;
    font-size: 20px;
    text-align: left;
}
.menu_item:hover .menu_text {
    background:linear-gradient(135deg, #50e3c2 0%,#10d574 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}

.menu_item.active {
    background: #0000001c;
    /* border: 1px solid #30363D; */
}
.active .menu_text {
    background:linear-gradient(135deg, #50e3c2 0%,#10d574 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}
</style>
