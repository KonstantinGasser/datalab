<template>
    <div class="tab-line">
        <div :class="{active: update===tab.name}" v-for="tab in tabs" :key="tab.name" class="tab" @click="setActive(tab.name)">
            <span>{{tab.emoji}}</span>&nbsp<span class="text">{{ tab.name }}</span>
        </div>
    </div>
</template>

<script>

export default {
  name: 'Tabs',
  data() {
      return {
          activeTab: this.$props.initTab,
          block: false,
      };
  },
  props: {
      initTab: String,
      update: String,
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
      },
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
    background: var(--tab-bg);
    font-size: 14px;
    font-weight: bolder;
    border-radius: 8px;
    padding: 5px 10px;
    border: 1px solid var(--tab-bg);
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
    background: var(--btn-bg-hover);
    border: 1px solid var(--sub-bg);
}
.active .text {
    -webkit-text-fill-color: var(--tab-font-selected); 
}

.text {
    background: var(--gradient-green);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}
</style>
