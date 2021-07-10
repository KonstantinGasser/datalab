<template>
<SideMenu class="app-menu" @setActive="showView"/>
  <div class="app-layout"> 
    <!-- <SideMenu class="app-menu" @setActive="showView"/> -->
    <ViewApp v-if="active_view === 'view_app'" class="app-view" :use_uuid="selected_uuid"/>
    <ViewAccount v-if="active_view === 'view_account'" class="app-view"/>
    <ViewOverview v-if="active_view === 'view_dashboard'" class="app-view" @openApp="openApp"/>
    <!-- <ViewCharts v-if="active_view === 'view_dashboard'" class="app-view" /> -->
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
import ViewOverview from '@/components/overview/ViewOverview';
import ViewCharts from '@/components/charts/ViewCharts';
import DocClient from '@/components/docs/DocClient';
import NotificationCenter from '@/components/notifications/NotificationCenter';

export default {
  name: 'Dashboard',
  data() {
    return {
      selected_uuid: null,
      active_view: 'view_app',
    };
  },
  components: {
    CompanyThumb,
    ViewApp,
    ViewOverview,
    ViewCharts,
    ViewAccount,
    NotificationCenter,
    DocClient,
    SideMenu,
  },
  created() {
    const url = "ws://192.168.0.177:8008/api/v1/datalab/live?token="+ localStorage.getItem("token");
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
    openApp(uuid) {
      
      this.selected_uuid = uuid;
      console.log("hello world ", this.selected_uuid);
      this.active_view = "view_app";
    }
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.app-layout {
  align-items: center;
  justify-content: center;
  margin: 0 auto;

  width: 100%;
  min-width: 750px;
  max-width: 1250px;
  height: 100%;
  padding: 65px 25px 0px 25px;

}

.app-view {
  grid-column: 2;
  grid-row: 2;
 
  border-radius: 8px;
}

</style>
