<template>
  <div class="side_menu">
      <div class="menu">
        <!-- <h4>Analysis Functions</h4> -->
        <!-- <a href="http://192.168.178.103:8080">Checkout referrer link</a> -->
        <MenuItem @click="setActive('view_dashboard')" :tabName="'view_dashboard'" :item="'Dashboard'" />
        <MenuItem @click="setActive('view_queries')" :tabName="'view_queries'" :item="'Queries'" />
        <MenuItem @click="setActive('view_app')" :tabName="'view_app'" :item="'Apps'" />
        <MenuItem @click="setActive('view_docs')" :tabName="'view_docs'" :item="'Docs'" />
        <MenuItem @click="setActive('view_account')" :tabName="'view_account'" :item="'Account'" />
        <MenuItem @click="setActive('view_notify')" :tabName="'view_notify'" :item="'Notification'" :bubble="true"/>
        <!-- <MenuItem @click="setActive('view_settings')" :tabName="'view_settings'" :item="'Settings'" /> -->
      </div>
      <div class="divider"></div>
      <div class="menu">
          <MenuItem @click="setActive('view_logout')" :item="'Logout 👋'" />
          <div class="d-flex justify-center">
            <input id="toggle" class="toggle" type="checkbox" @change="setMode($event)">
          </div>
        </div>
  </div>
</template>

<script>
import MenuItem from '@/components/side_menu/MenuItem.vue';

export default {
  name: 'SideMenu',
  data() {
    return {
      mode: false,
    };
  },
  components: {
      MenuItem,
  },
  methods: {
      setActive(value) {
          this.$store.commit('setActiveTab', value);
          this.$emit('setActive', value);
      },
      setMode() {
        var root = document.documentElement
        if (this.mode) {
          root.style.setProperty("--main-bg", "#cccccc1B");
          root.style.setProperty("--sub-bg", "#fff");
          root.style.setProperty("--tab-bg", "#fff");
          root.style.setProperty("--btn-font-hover", "#fff");
          root.style.setProperty("--tab-font-selected", "#fff");
          root.style.setProperty("--h-color", "#666666");
          root.style.setProperty("--txt-small", "#666666AA");
          root.style.setProperty("--main-color", "#666666AA");
        } else {
          root.style.setProperty("--main-bg", "#1E1E1E");
          root.style.setProperty("--sub-bg", "#2F2F2F");
          root.style.setProperty("--btn-font-hover", "#121212");
          root.style.setProperty("--tab-font-selected", "#121212");
          root.style.setProperty("--tab-bg", "#666666");
          root.style.setProperty("--h-color", "#ccc");
          root.style.setProperty("--txt-small", "#FFFFFFAA");
          root.style.setProperty("--main-color", "#FFFFFFAA");
        }
        this.mode = !this.mode;
      },
  },
  computed: {
      checkActive(value) {
          return this.items[value];
      }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

h4 {
    font-size: 20px;
    margin: 5px 0px;
    color: #000;
}
.side_menu {
    display: grid;
    grid-template-rows: min-content min-content;
    display: grid;
    justify-content: flex-start;
    height: 100%;
    align-content: flex-start;
    border-radius: 8px 0 0 8px;
}


.divider {
  height: 100px;
}
.toggle {
  --size: 1.5rem;
  -webkit-appearance: none;
     -moz-appearance: none;
          appearance: none;
  outline: none;
  cursor: pointer;
  width: var(--size);
  height: var(--size);
  box-shadow: inset calc(var(--size) * 0.33) calc(var(--size) * -0.25) 0;
  border-radius: 999px;
  color:  #03045e;
  transition: all 500ms;
}
.toggle:checked {
  --ray-size: calc(var(--size) * -0.4);
  --offset-orthogonal: calc(var(--size) * 0.65);
  --offset-diagonal: calc(var(--size) * 0.45);
  transform: scale(0.75);
  color: #ffaa00;
  box-shadow: inset 0 0 0 var(--size), calc(var(--offset-orthogonal) * -1) 0 0 var(--ray-size), var(--offset-orthogonal) 0 0 var(--ray-size), 0 calc(var(--offset-orthogonal) * -1) 0 var(--ray-size), 0 var(--offset-orthogonal) 0 var(--ray-size), calc(var(--offset-diagonal) * -1) calc(var(--offset-diagonal) * -1) 0 var(--ray-size), var(--offset-diagonal) var(--offset-diagonal) 0 var(--ray-size), calc(var(--offset-diagonal) * -1) var(--offset-diagonal) 0 var(--ray-size), var(--offset-diagonal) calc(var(--offset-diagonal) * -1) 0 var(--ray-size);
}

.toggle {
  z-index: 1;
}
.toggle:checked ~ .background {
  --bg: white;
}
/* .toggle:checked ~ .title {
  --color: hsl(40, 100%, 50%);
}

.title {
  --color: hsl(240, 100%, 95%);
  color: var(--color);
  z-index: 1;
  cursor: pointer;
  display: block;
  padding: 0.5rem 0 0;
  transition: color 500ms;
} */


</style>
