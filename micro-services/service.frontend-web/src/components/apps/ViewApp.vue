<template>
    <div class="view_app">
        <div class="app_list">
            <!-- <div class="add_new_app d-flex justify-center align-center">
                <button @click="modeCreateApp()" class="btn btn-standard">Create App <span class="">🙌</span></button>
            </div> -->
            <div class="app_name_list">
                <p class="info-text" v-if="apps == null || apps.length === 0">
                    Mhm looks like you do not have any apps yet - <a @click="modeCreateApp()">go create one!</a>
                </p>
                <div :class="{selected: selectedApp===app.uuid}" class="app-name d-flex justify-between align-center" v-for="app in apps" :key="app.uuid" @click="loadApp(app.uuid)">
                    <span class="dots medium-font" >{{ app.name }}</span>
                    <span v-if="app.private" class="icon icon-lock"></span> 
                    <span v-if="app.private === undefined" class="icon icon-unlock"></span>
                </div>
                <button @click="modeCreateApp()" class="btn btn-standard w-100">Create App <span class="">🙌</span></button>
            </div>
        </div>
        <div>
            <TabCreateApp v-if="isInCreateMode" @createdApp="updateState" :orgn_domain="activeApp.owner?.orgn_domain" />
            <div v-if="!isInCreateMode">
                <h1 class="super-lg d-flex align-center">
                     {{activeApp.owner?.orgn_domain}}/{{activeApp.app?.name}}
                    <div v-if="activeApp?.app?.locked" class="locked d-flex align-center justify-center">
                        <div>locked</div>
                    </div>
                    <div  v-if="!app_unsaved && !activeApp?.app?.locked" class="saved d-flex align-center justify-center">
                        <div>Saved</div>
                    </div>
                    <div  v-if="app_unsaved && !activeApp?.app?.locked" class="unsaved d-flex align-center justify-center">
                        <div>Unsaved Changes</div>
                    </div>
                    <div v-if="activeApp?.app?.uuid === sync_app?.uuid && sync_app?.sync && !activeApp?.app?.locked" class="sync d-flex align-center justify-center" @click="syncAppChanges(activeApp?.app?.uuid)">
                        <div>sync data</div>
                    </div>
                </h1>
                <div class="desc_test">
                    <span class="app-tag blue" v-for="tag in activeApp?.app?.tags" :key="tag">
                        #{{tag}}
                    </span>
                </div>
                <!-- <hr> -->
                <Tabs class="mt-3 mb-4" ref="Tabs" :update="activeTab" :initTab="activeTab" :tabs="getTabs" @tabChange="tabChange"/>
                <General v-if="activeTab === 'Overview'" 
                    :app_locked="activeApp?.app?.locked === true" 
                    :app_token="activeApp?.token" 
                    :app_uuid="activeApp?.app?.uuid" 
                    :app_owner="activeApp?.owner?.uuid"
                    @drop_app="drop_app" 
                    @loadApp="loadApp"/>

                <Config v-if="activeTab == 'Configuration'" 
                    :app_locked="activeApp?.app?.locked === true" 
                    :app_config="activeApp?.config" 
                    :app_uuid="activeApp?.app?.uuid"
                    @setdoc="setdoc" 
                    @appchange="markUnsaved"/>

                <InviteMember v-if="activeTab == 'Invite'" 
                    :app_uuid="activeApp?.app?.uuid" 
                    :member="activeApp?.app?.member" 
                    :app_owner="activeApp?.owner" 
                    :app_name="activeApp?.app?.name"/>

            </div>
        </div>
    </div>
</template>

<script>
    // import TabAppDetails from '@/components/apps/TabAppDetails.vue';
    import TabCreateApp from '@/components/apps/TabCreateApp.vue';
    import Tabs from '@/components/utils/Tabs.vue';
    import General from '@/components/apps/General.vue';
    import Config from '@/components/apps/Config';
    import InviteMember from '@/components/apps/InviteMember';
    import axios from 'axios';
    import jwt_decode from "jwt-decode";

    export default {
        name: 'ViewApp',
        components: {
            Tabs,
            TabCreateApp,
            General,
            Config,
            InviteMember,
        },
        computed: {
            getTabs() {
                console.log(this.activeApp?.app?.is_private === undefined)
                if ((this.activeApp?.app?.is_private === undefined) === true) {
                    console.log("no invite")
                    return [
                        {name:'Overview', emoji: "icon-package"},
                        {name:'Configuration', emoji: "icon-sliders"},
                    ]
                }
                return [
                        {name:'Overview', emoji: "icon-package"},
                        {name:'Configuration', emoji: "icon-sliders"},
                        {name:'Invite', emoji: "icon-users"},
                    ]
            },
            sync_app() {
                return this.$store.state.sync_app
            },
            blockTabs() {
                return this.tabsBlocked;
            },
            showCreateApp() {
                return this.isInCreateMode;
            },
        },
        props: {
            use_uuid: {
                type: String,
                default: null,
            },
        },
        data() {
            return {
                loggedInUser: null,
                selectedApp: null,
                isInCreateMode: false,
                activeTab: 'App Details',
                tabsBlocked: false,
                newApp: { name: '' },
                apps: [],
                activeApp: {},
                app_unsaved: false,
                token_placeholder: "Organization-Domain/App-Name",
                activeTab: "Overview",
                tabs: [
                        {name:'Overview', emoji: "icon-package"},
                        {name:'Configuration', emoji: "icon-sliders"},
                        {name:'Invite', emoji: "icon-users"},
                    ]
            };
        },
        async created() {
            this.loggedInUser = jwt_decode(localStorage.getItem("token"));
            // fetch initial data of app list
            
            const init_data = await this.getAppList();
            if (init_data.data.apps === undefined || init_data.data.apps === null || init_data.data.apps.length <= 0 || init_data.status != 200) {
                this.apps = [];
                this.isInCreateMode = true;
            }
            else {
                this.apps = init_data.data.apps;
                const init_app = await this.getApp(this.apps[0].uuid);

                this.activeApp = init_app;
                this.selectedApp = this.activeApp?.app?.uuid;
                this.isInCreateMode = false;
            }
        },
        methods: {
            async syncAppChanges(uuid){
                const data = await this.getApp(uuid);
                if (data.app === undefined || data.app == null) {
                    return
                }
                this.activeApp = data; 
                this.$store.commit("UNSYNC_APP")
            },
            async loadApp(uuid) {
                this.isInCreateMode = false;
                const data = await this.getApp(uuid);
                if (data.app === undefined || data.app == null) {
                    return
                }
                this.activeApp = data;
                this.selectedApp = data.app?.uuid;
                this.activeTab = "Overview";
                this.app_unsaved = false;
            },
            async getAppList() {
                let options = {
                    headers: {
                        'Authorization': localStorage.getItem("token"),
                    }
                };
                const resp = await axios.get("http://192.168.0.177:8080/api/v1/app/all", options)
                if (resp.status != 200) {
                    this.$toast.error(resp.data);
                }
                return resp
            },
            async getApp(uuid) {
                let options = {
                    headers: {
                        'Authorization': localStorage.getItem("token"),
                    }
                };
                let resp = {}
                try {
                    resp = await axios.get("http://192.168.0.177:8080/api/v1/app?app="+uuid, options)
                    if (resp.status != 200) {
                        this.$toast.error(resp.data);
                    }
                    return resp.data
                }catch(err) {
                    this.isInCreateMode = true
                    return null;
                }
                
            },
            tabChange(tab) {
                this.activeTab = tab;
            },
            async updateState(event) {
                const init_data = await this.getAppList();
                this.apps = init_data.data.apps.reverse();
                const init_app = await this.getApp(this.apps[0].uuid);
                this.activeApp = init_app;
                this.selectedApp = this.activeApp?.app?.uuid;
                this.isInCreateMode = false;
            },
            markUnsaved(value) {
                this.$store.commit("UNSYNC_APP")
                switch (value.type) {
                    case "funnel-add":
                        if (this.activeApp?.config?.funnel === undefined) {
                            this.activeApp.config.funnel = []
                        }
                        this.activeApp?.config?.funnel?.push(value.item)
                        this.app_unsaved = value.unsaved
                        break;
                    case "funnel-remove":
                        this.activeApp.config.funnel = this.activeApp?.config?.funnel?.filter(item => item.id !== value.item)
                        this.app_unsaved = value.unsaved
                        break;
                    case "funnel-saved":
                        this.app_unsaved = false;
                        break;
                    case "campaign-add":
                        if (this.activeApp?.config?.campaign === undefined) {
                            this.activeApp.config.campaign = []
                        }
                        this.activeApp?.config?.campaign?.push(value.item)
                        this.app_unsaved = value.unsaved
                        break;
                    case "campaign-remove":
                        this.activeApp.config.campaign = this.activeApp?.config?.campaign?.filter(item => item.id !== value.item)
                        this.app_unsaved = value.unsaved
                        break;
                    case "campaign-saved":
                        this.app_unsaved = false;
                        break;
                    case "btn-add":
                        if (this.activeApp?.config?.btn_time === undefined) {
                            this.activeApp.config.btn_time = []
                        }
                        this.activeApp?.config?.btn_time?.push(value.item)
                        this.app_unsaved = value.unsaved
                        break;
                    case "btn-remove":
                        this.activeApp.config.btn_time = this.activeApp?.config?.btn_time?.filter(item => item.id !== value.item)
                        this.app_unsaved = value.unsaved
                        break;
                    case "btn-saved":
                        this.app_unsaved = false;
                        break;
                }
            },
            drop_app() {
            },
            modeCreateApp() {
                this.isInCreateMode = true;
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

.desc_test {
    display: flex;
    flex-wrap: wrap;
    max-width: 500px;
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
    color: var(--h-color);
    font-size: 16px;
    padding: 5px;
    margin: 5px 0;
    border-radius: 8px;
    border-style: dashed;
    border-color: var(--sub-border);
    border-width: 1px;
}
.app-name:hover {
    cursor: pointer;
    color: var(--menu-bg);
}
.selected {
    background: var(--sub-border) !important;
    color: var(--menu-bg);
}

.saved {
    margin-left: 15px;
    /* width: auto;
    height: 25px; */
    border-radius: 8px;
    background: var(--menu-bg);
    padding: 5px 15px;
    color: var(--sub-bg);
    font-size: 16px;
    font-weight: bold;
}

.unsaved {
    margin-left: 15px;
    width: auto;
    height: 25px;
    border-radius: 8px;
    background: #ffa500;
    padding: 5px 15px;
    color: var(--sub-bg);
    font-size: 16px;
    font-weight: bold;
}

.locked {
    margin-left: 15px;
    background: #ffa50050;
    border: 1px solid #ffa500;
    border-radius: 8px;
    padding: 5px 15px;
    font-size: 16px;
    font-weight: bold;
    color: var(--sub-bg);
}
.locked .icon {
    margin-right: 5px;
    font-size: 18px;
    font-weight: bolder;
}

.sync {
    cursor: pointer;
    margin-left: 15px;
    width: auto;
    height: 25px;
    border-radius: 8px;
    background: #ffa50050;
    border: 1px solid #ffa500;
    padding: 5px 15px;
    color: var(--sub-bg);
    font-size: 16px;
    font-weight: bold;
}
.sync span {
    font-size: 16px;
    font-weight: bolder;
}
</style>