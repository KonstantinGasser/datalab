// <template>
//     <div>
//         <h1 class="super-lg">{{app.owner?.orgn_domain}}/{{app.name}}</h1>
//         <div class="desc_test">{{app.description}}</div>
//         <hr>
//         <Tabs class="my-2" ref="Tabs" :update="activeTab" :initTab="activeTab" :tabs="tabs" @tabChange="tabChange"/>
//         <General v-if="activeTab === 'Overview'" :app="app" @drop_app="drop_app" :token_placeholder="token_placeholder"/>
//         <Config v-if="activeTab == 'Configuration'" :app_uuid="app.uuid" :app_config="app.app_config" />
//         <DocClient v-if="activeTab === 'Client Library'" :hasToken="false" @goCreateToken="goCreateToken"/>
//     </div>
// </template>

// <script>
//     import Tabs from '@/components/utils/Tabs.vue';
//     import General from '@/components/apps/General.vue';
//     import Config from '@/components/apps/Config';
//     import DocClient from '@/components/apps/DocClient.vue';

//     export default {
//         name: 'TabAppDetails',
//         components: {
//             Tabs,
//             General,
//             Config,
//             DocClient,
//         },
//         data() {
//             return {
//                 token_placeholder: "Organization-Domain/App-Name",
//                 activeTab: "Overview",
//                 tabs: [{name:'Overview', emoji: "🎛"},{name:'Configuration', emoji: "⚙️"},{name:'Client Library', emoji: "📓"}]
//             };
//         },
//         props: ['app'],
//         created() {
//             // this.header_name = this.$props.app.owner.orgn_domain + "/" + this.$props.app.name;
//         },
//         computed: {
//             token() {
//                 console.log(this.$props.app.app_token);
//                 if (this.$props.app.app_token) {
//                     return this.$props.app.app_token;
//                 } 
//                 if (this.app_token) {
//                     return this.app_token;
//                 }
//                 return "";
//                 },
//         },
//         methods: {
//              tabChange(tab) {
//                 this.activeTab = tab;
//             },
//             drop_app(event) {
//                 this.$emit("drop_app", event);
//             },
//             goCreateToken() {
//                 this.token_placeholder = this.$props.app?.owner?.orgn_domain + "/" + this.$props.app?.name;
//                 this.activeTab = "Overview";
//             },
//         },
//     };
// </script>

// <style scoped>
// .view_component {
//     margin-top: 15px;
//     padding: 15px;
//     border-radius: 8px;
//     height: max-width;
// }
// h2 {
//     margin: 5px 0px;
// }
// .pos_1_1 {
//     grid-row: 1;
// }
// .pos_1_2 {
//     grid-row: 2;
// }

// input[type="file"] {
//     display: none;
// }
// .custom-file-upload {
//     border: 1px solid #50e3c2;
//     color: #10d574;
//     border-radius: 8px;
//     display: inline-block;
//     padding: 6px 12px;
//     cursor: pointer;
// }

// .desc_test {
//     height: auto;
//     max-height: 85px;
//     overflow-y: scroll;
// }

// .custom-file-upload:hover {
//   -webkit-font-smoothing: antialiased;
//   -moz-osx-font-smoothing: grayscale;
//   background: linear-gradient(0deg, #50e3c2 0%,#10d574 100%);
//   color: #121212;
// }

// .img-grid {
//     display: flex;
//     width: 100%;
//     max-width: 800px;
//     overflow-x: scroll;
// }

// .img-card {
//     margin: 0 5px;
//     width: 300px;
//     min-width: 300px;
//     height: auto;
//     padding: 5px;
// }
// .img-card img {
//     height: 150px;
//     width: 100%;
//     border-radius: 6px;
//     margin-bottom: 5px;
// }
// </style>
