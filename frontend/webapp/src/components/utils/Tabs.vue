<template>
    <div class="tab-line">
        <div :class="{active: activeTab===tab.name}" v-for="tab in tabs" :key="tab.name" class="tab" @click="setActive(tab.name)">
            <span class="text">{{ tab.name }}</span>
        </div>
    </div>
</template>

<script>

export default {
  name: 'Tabs',
  data() {
      return {
          activeTab: this.initTab,
          block: false,
      };
  },
  props: {
      initTab: String,
      tabs: {
          type: Array,
          default() { return []; },
      },
  },
  methods: {
      setActive(clicked) {
        // allow tab change only if !blocked
        if (this.block) return;
        // emit change to parent
        this.activeTab = clicked;
        this.$emit('tabChange', clicked);
      }
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
hr {
    margin: 3px 0;
}
.tab-line {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  grid-column-gap: 5px;
}
.tab {
    background: #30363D;
    font-size: 14px;
    font-weight: bolder;
    border-radius: 8px;
    padding: 5px 10px;
    border: 1px solid transparent;
}
.tab:hover {
    cursor: pointer;
}
.tab-line.block {
    opacity: 0.5;
}
.tab-line.block .tab:hover {
    cursor: default;
}

.active {
    background: linear-gradient(135deg, #50e3c2 0%,#10d574 100%);
    border: 1px solid #30363D;
}
.active .text {
    -webkit-text-fill-color: #0D1116; 
}

.text {
    background: linear-gradient(135deg, #50e3c2 0%,#10d574 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}
</style>
