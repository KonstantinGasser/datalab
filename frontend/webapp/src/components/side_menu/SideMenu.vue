<template>
  <div class="side_menu">
      <div class="menu">
        <!-- <h4>Analysis Functions</h4> -->
        <MenuItem @click="setActive('view_dashboard')" :tabName="'view_dashboard'" :item="'Dashboard'" />
        <MenuItem @click="setActive('view_queries')" :tabName="'view_queries'" :item="'Queries'" />
        <MenuItem @click="setActive('view_app')" :tabName="'view_app'" :item="'Apps'" />
        <MenuItem @click="setActive('view_account')" :tabName="'view_account'" :item="'Account'" />
        <!-- <MenuItem @click="setActive('view_settings')" :tabName="'view_settings'" :item="'Settings'" /> -->
        <MenuItem @click="setActive('view_logout')" :item="'Logout 👋'" /> 
        <div class="custom-control custom-switch d-flex justify-center">
            <input v-model="appCfgs" :value="'css-mode'" @change="setMode($event)" type="checkbox" class="custom-control-input" id="css-mode">
            <label class="custom-control-label" for="css-mode"></label>
            <!-- {{mode ? "Light Mode" : "Dark Mode"}} -->
        </div>
      </div>
      <!-- <div class="menu bottom-set"> -->
        <!-- <h4>App Settings</h4> -->
        <!-- <div class="custom-control custom-switch">
            <input v-model="appCfgs" :value="'mouse-move-map'" @change="setMode($event)" type="checkbox" class="custom-control-input" id="mouse-move-map">
            <label class="custom-control-label" for="mouse-move-map">{{mode ? "Light Mode" : "Dark Mode"}}</label>
        </div>
        <MenuItem @click="setActive('view_logout')" :item="'Logout 👋'" /> 
      </div>  -->
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
          root.style.setProperty("--btn-font-hover", "#fff");
          root.style.setProperty("--tab-font-selected", "#fff");
          root.style.setProperty("--h-color", "#666666");
          root.style.setProperty("--txt-small", "#666666AA");
          root.style.setProperty("--main-color", "#666666AA");
        } else {
          root.style.setProperty("--main-bg", "#1E1E1E");
          root.style.setProperty("--sub-bg", "#1E1E1E");
          root.style.setProperty("--btn-font-hover", "#121212");
          root.style.setProperty("--tab-font-selected", "#121212");
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
    align-content: space-around;
    border-radius: 8px 0 0 8px;
}
</style>
