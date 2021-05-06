<template>
    <div class="view_app">
        <div class="app_list">
            <div class="add_new_app d-flex justify-center align-center">
                <button @click="modeCreateApp()" class="btn btn-standard">Create App <span class="">🙌</span></button>
            </div>
            <div class="app_name_list">
                <p class="info-text" v-if="app_list == null || app_list.length === 0">
                    Mhm looks like you do not have any apps yet - <a @click="modeCreateApp()">go create one!</a>
                </p>
                <div class="app-name d-flex justify-start align-center" v-for="app in app_list" :key="app.uuid" @click="getAppDetails(app.uuid)">
                    <span class="dots medium-font" >{{ app.name }}</span>
                    <!-- <span class="icon icon-delete hover big" @click="removeApp(app.uuid)"></span> -->
                </div>
            </div>
        </div>
        <div>
            <TabCreateApp v-if="isInCreateMode" @createdApp="updateState" :orgn_domain="activeApp.owner?.orgn_domain" />
            <div v-if="!isInCreateMode">
                <h1 class="super-lg">{{activeApp.owner?.orgn_domain}}/{{activeApp.name}}</h1>
                <div class="desc_test">{{activeApp.description}}</div>
                <hr>
                <Tabs class="my-2" ref="Tabs" :update="activeTab" :initTab="activeTab" :tabs="tabs" @tabChange="tabChange"/>
                <General v-if="activeTab === 'Overview'" :app="activeApp" @drop_app="drop_app" :token_placeholder="token_placeholder"/>
                <Config v-if="activeTab == 'Configuration'" :app_config="activeApp.app_config" :config_uuid="activeApp.config_ref" @setdoc="setdoc"/>
                <DocClient v-if="activeTab === 'Documentation'" :hasToken="activeApp?.app_token !== ''" @goCreateToken="goCreateToken"/>
            </div>
            <!-- <TabAppDetails ref="tab_app_token" @drop_app="updateState" v-cloak v-if="activeTab === 'App Details' && !isInCreateMode" :app="activeApp"/> -->
        </div>
    </div>
</template>

<script>
    // import TabAppDetails from '@/components/apps/TabAppDetails.vue';
    import TabCreateApp from '@/components/apps/TabCreateApp.vue';
    import Tabs from '@/components/utils/Tabs.vue';
    import General from '@/components/apps/General.vue';
    import Config from '@/components/apps/Config';
    import DocClient from '@/components/apps/DocClient.vue';
    import axios from 'axios';

    export default {
        name: 'ViewApp',
        components: {
            Tabs,
            TabCreateApp,
            General,
            Config,
            DocClient,
        },
        computed: {
            blockTabs() {
                return this.tabsBlocked;
            },
            showCreateApp() {
                return this.isInCreateMode;
            },
            app_list() {
                return this.apps;
            }
        },
        data() {
            return {
                isInCreateMode: false,
                activeTab: 'App Details',
                tabsBlocked: false,
                newApp: { name: '' },
                apps: [],
                activeApp: {},
                token_placeholder: "Organization-Domain/App-Name",
                activeTab: "Overview",
                tabs: [{name:'Overview', emoji: "🎛"},{name:'Configuration', emoji: "⚙️"},{name:'Documentation', emoji: "📓"}]
            };
        },
        async created() {
            // fetch initial data
            this.getViewApp().then(data => {
                this.apps = data.app_list;
                this.activeApp = data.app_details;
                this.activeApp["app_token"] = data.app_token;
                this.activeApp["app_config"] = {
                        "funnel": data.config_funnel,
                        "campaign": data.config_campaign,
                        "btn_time": data.config_btn_time,
                    }
                if (this.apps == null || this.apps.length === 0) {
                    this.isInCreateMode = true;
                } else {
                    this.isInCreateMode = false;
                }
            }).catch(err => {
               if (err.response.status === 401) {
                        localStorage.removeItem('token');
                        this.$router.replace({ name: 'login' });
                }
                this.isInCreateMode = true;
            });
            
        },
        props: ['status'],
        methods: {
            tabChange(tab) {
                this.activeTab = tab;
            },
            updateState(event) {
                // update app list on the left to show newly created app
                this.getViewApp().then(data => {
                    this.apps = data.app_list;
                    this.activeApp = data.app_details;
                    this.activeApp["app_token"] = data.app_token;
                    this.activeApp["app_config"] = {
                        "funnel": data.config_funnel,
                        "campaign": data.config_campaign,
                        "btn_time": data.config_btn_time,
                    }
                    this.isInCreateMode = false;
                }).catch(error => {
                    if (error.response.status === 401) {
                        localStorage.removeItem('token');
                        this.$router.replace({ name: 'login' });
                    }
                    console.log(error);
                    this.$toast.error("could not refresh app list");
                });
                
                
            },
            drop_app() {
                this.isInCreateMode = true;
                this.apps = this.apps.filter(item => item.uuid != event.app_uuid);
                this.getViewApp().then(data => {
                    this.apps = data.app_list;
                    this.activeApp = data.app_details;
                    this.activeApp["app_token"] = data.app_token;
                    this.activeApp["app_config"] = {
                        "funnel": data.config_funnel,
                        "campaign": data.config_campaign,
                        "btn_time": data.config_btn_time,
                    }
                }).catch(error => {
                    if (error.response.status === 401) {
                        localStorage.removeItem('token');
                        this.$router.replace({ name: 'login' });
                    }
                    console.log(error);
                    this.$toast.error("could not refresh app list");
                });
            },
            modeCreateApp() {
                this.isInCreateMode = true;
            },
            async getViewApp() {
                let options = {
                    headers: {
                        'Authorization': localStorage.getItem("token"),
                    }
                };

                const res = await axios.get("http://localhost:8080/api/v2/view/app/details", options)
                if (res.data == null || res.status >= 400) {
                    this.isInCreateMode = true;
                    console.log(this.isInCreateMode);
                    return null;
                }
                this.activeTab = "Overview";
                return res.data;
                
            },
            async getAppDetails(uuid) {
                let options = {
                    headers: {
                        'Authorization': localStorage.getItem("token"),
                    }
                };
                axios.get("http://localhost:8080/api/v2/view/app/get?uuid="+uuid, options).then(resp => {
                    if (this.isInCreateMode)
                        this.isInCreateMode = !this.isInCreateMode;
                    this.activeApp = resp.data.app;
                    this.activeApp["app_token"] = resp.data.app_token;
                    this.activeApp["app_config"] = {
                        "funnel": resp.data.config_funnel,
                        "campaign": resp.data.config_campaign,
                        "btn_time": resp.data.config_btn_time,
                    }
                    this.activeTab = "Overview";
                }).catch(error => {
                    console.log(error);
                });
            },
            setdoc(event) {
                this.activeTab = "Documentation";
            },
        },
    };
</script>

<style scoped>
.view_component {
    margin-top: 15px;
    padding: 15px;
    border-radius: 8px;
    height: max-width;
}
h2 {
    margin: 5px 0px;
}
.pos_1_1 {
    grid-row: 1;
}
.pos_1_2 {
    grid-row: 2;
}
.tab-line {
    grid-column: 1;
    grid-row: 1;
}
.view_app {
    display: grid;
    grid-template-columns: 20% 1fr;
    grid-column-gap: 15px;
    height: 100%;
}
  
.app_list {
    background: var(--sub-bg);
    border-radius: 8px;
    padding: 15px;
    height: max-content;
    min-height: 225px;
    max-height: 100%;
    overflow-y: auto;
    border: 1px solid var(--sub-border);
}
.app_name_list {
    margin-top: 25px;
    padding: 0 7px;
}
.app-name {
    background: var(--gradient-green);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent; 
    font-size: 16px;
    font-weight: bold;
    padding: 5px;
    margin: 5px 0;
    border-radius: 8px;
    border-style: dashed;
    border-color: var(--sub-border);
    border-width: 1px;
}
.app-name:hover {
    cursor: pointer;
    text-decoration: underline;
}
</style>