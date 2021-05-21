<template>
  <div class="app-layout">
    <SideMenu class="app-menu" @setActive="showView"/>
    <div class="app-header">
    </div>
    <ViewApp v-if="active_view === 'view_app'" class="app-view" />
    <ViewAccount v-if="active_view === 'view_account'" class="app-view" />
    <ViewCharts v-if="active_view === 'view_dashboard'" class="app-view" />
    <DocClient v-if="active_view === 'view_docs'" class="app-view" />
    <NotificationCenter v-if="active_view === 'view_notify'" class="app-view" />
  </div>
  <vue-confirm-dialog></vue-confirm-dialog>
</template>

<script>
import SideMenu from '@/components/side_menu/SideMenu.vue';
import CompanyThumb from '@/components/company/CompanyThumb.vue';
import ViewApp from '@/components/apps/ViewApp.vue';
import ViewAccount from '@/components/account/ViewAccount.vue';
import ViewCharts from '@/components/charts/ViewCharts';
import DocClient from '@/components/docs/DocClient';
import NotificationCenter from '@/components/notifications/NotificationCenter';

export default {
  name: 'Dashboard',
  data() {
    return {
      active_view: 'view_app',
    };
  },
  components: {
    CompanyThumb,
    ViewApp,
    ViewCharts,
    ViewAccount,
    NotificationCenter,
    DocClient,
    SideMenu,
  },
  created() {
    const url = "ws://localhost:8008/api/v1/datalab/live?token="+ localStorage.getItem("token");
    this.$connect(url);
  },
  methods: {
    showView(view) {
      // logout user
      if (view === 'view_logout') {
        localStorage.removeItem('token');
        this.$router.replace({ name: 'login' });
      }
      // serve view
      this.active_view = view;
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.app-layout {
  display: grid;
  grid-template-columns: 200px minmax(auto, 1100px);
  grid-template-rows: auto 1fr;
  align-items: center;
  justify-content: center;
  margin: 0 auto;

  width: 90%;
  min-width: 750px;
  height: 90%;
}

.app-menu {
  grid-column: 1;
  grid-row: 2 / 4;
}

.app-header {
  display: flex;
  justify-content: flex-end;
  align-content: center;
  grid-column: 2;
  grid-row: 1;
  height: 45px;
  /* background: linear-gradient(270deg, #50e3c2 0%,#10d574 100%); */
}

.app-view {
  grid-column: 2;
  grid-row: 2;
 
  border-radius: 8px;
}
</style>
