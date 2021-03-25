<template>
  <div class="side_menu">
      <div class="menu">
        <!-- <h4>Analysis Functions</h4> -->
        <MenuItem @click="setActive('view_dashboard')" :tabName="'view_dashboard'" :item="'Dashboard'" />
        <MenuItem @click="setActive('view_queries')" :tabName="'view_queries'" :item="'Queries'" />
        <MenuItem @click="setActive('view_app')" :tabName="'view_app'" :item="'Apps'" />
        <MenuItem @click="setActive('view_account')" :tabName="'view_account'" :item="'Account'" />
        <MenuItem @click="setActive('view_settings')" :tabName="'view_settings'" :item="'Settings'" />
      </div>
      <div class="menu bottom-set">
        <!-- <h4>App Settings</h4> -->
        <MenuItem @click="setActive('view_logout')" :item="'Logout 👋'" /> 
      </div>
  </div>
</template>

<script>
import MenuItem from '@/components/side_menu/MenuItem.vue';

export default {
  name: 'SideMenu',
  data() {
    return {

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
    /* padding: 65px 10px 10px 20px; */
    /* background: linear-gradient(0deg, #50e3c2 0%,#10d574 100%); */
    display: grid;
    grid-template-rows: min-content min-content;
    color: #fff;
    
    display: grid;
    justify-content: flex-start;
    height: 100%;
    align-content: space-around;
    /* background: rgba(255,255,255,1); */
    border-radius: 8px 0 0 8px;
}
/* .bottom-set {
  margin-top: 100%;
} */
</style>
