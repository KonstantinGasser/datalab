<template>
    <div class="view_app">
        <div class="app_list">
            <div class="add_new_app d-flex justify-center align-center">
                <button @click="modeCreateApp()" class="btn btn-standard">Create new App <span class="icon icon-plus"></span></button>
            </div>
            <div class="app_name_list">
                <p class="info-text" v-if="app_list == null || app_list.length === 0">
                    Mhm looks like you do not have any apps yet - <a @click="modeCreateApp()">go create one!</a>
                </p>
                <div class="app-name d-flex justify-between align-center" v-for="app in app_list" :key="app.uuid">
                    <span class="dots standard-font" @click="getAppDetails(app.uuid)">{{ app.name }}</span>
                    <span class="icon icon-delete hover big" @click="removeApp(app.uuid)"></span>
                </div>
            </div>
        </div>
        <div>
            <Tabs v-if="!showCreateApp" ref="Tabs" :class="{ block: blockTabs }" :initTab="'App Details'" :tabs="[{name:'App Details'},{name:'Member'}]" @tabChange="tabChange"/>
            <TabCreateApp v-if="showCreateApp" @createdApp="updateState" />
            <TabAppDetails ref="tab_app_token" v-cloak v-if="activeTab === 'App Details' && !showCreateApp" :app="activeApp"/>
            <TabMember ref="tab_member" v-if="activeTab === 'Member' && !showCreateApp" :member_list="activeApp.app_member"/>
        </div>
    </div>
</template>

<script>
    import Tabs from '@/components/utils/Tabs.vue';
    import TabAppDetails from '@/components/apps/TabAppDetails.vue';
    import TabCreateApp from '@/components/apps/TabCreateApp.vue';
    import TabMember from '@/components/apps/TabMember.vue';
    import axios from 'axios';

    export default {
        name: 'ViewApp',
        components: {
            Tabs,
            TabCreateApp,
            TabAppDetails,
            TabMember,
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
                isInCreateMode: false, // change back to false once api works again. true just to check on the other views
                activeTab: 'App Details',
                tabsBlocked: false,
                newApp: { name: '' },
                apps: [],
                activeApp: {},
            };
        },
        async created() {
            // fetch initial data
            this.getViewApp().then(data => {
                this.apps = data.app_list;
                this.activeApp = data.app_details;
                if (this.apps == null || this.apps.length === 0) {
                    this.isInCreateMode = true;
                } else {
                    this.isInCreateMode = false;
                    // if apps call full app 
                }
                // if (this.app !== null && this.apps.length === 0) this.isInCreateMode = true;
            }).catch(error => {
               if (err.response.status === 401) {
                        localStorage.removeItem('token');
                        this.$router.replace({ name: 'login' });
                }
                this.isInCreateMode = true;
            });
            
        },
        props: ['status'],
        methods: {
            updateState(event) {
                switch (event.type) {
                    case "show_app":
                        this.isInCreateMode = false;
                        // update app list on the left to show newly created app
                        this.getViewApp().then(data => {
                            this.apps = data.app_list;
                        }).catch(error => {
                            if (error.response.status === 401) {
                                localStorage.removeItem('token');
                                this.$router.replace({ name: 'login' });
                            }
                            console.log(error);
                            this.$toast.error("could not refresh app list");
                        });
                        
                        break;
                    default:
                        break;
                }
                this.isInCreateMode = false;
                
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
                return res.data;
                
            },
            async getAppDetails(uuid) {
                let options = {
                    headers: {
                        'Authorization': localStorage.getItem("token"),
                    }
                };
                axios.get("http://localhost:8080/api/v2/view/app/get?uuid="+uuid, options).then(resp => {
                    console.log(this.$refs);
                    this.activeApp = resp.data.app;
                    this.$refs.tab_member.updateComponent(resp.data.member);
                    this.$refs.tab_app_token.updateComponent(resp.data);
                }).catch(error => {
                    console.log(error);
                });
            },
            async removeApp(id) {
                let options = {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': localStorage.getItem("token"),
                    }
                };
                await axios.post("http://localhost:8080/api/v2/view/app/delete", {
                        app_uuid: id,
                    }, options
                ).then((resp) => {
                    if (resp.status == 200) {
                        this.$toast.success("App has been deleted");
                        this.apps = this.apps.filter(item => item.uuid != id);
                    }
                }).catch(err => {
                    this.$toast.warning("Sorry app could not be removed");
                    return;
                });
                const al = await this.getAppList();
                this.apps = al;
            },
            tabChange(tab) {
                this.isInCreateMode = false;
                this.activeTab = tab;
            },
            enableDisableTabs(toggle) {
                this.tabsBlocked = toggle;
                // block tab change by sending toggle to @Tabs
                this.$refs.Tabs.block = this.tabsBlocked;
            }
        },
    };
</script>

<style scoped>
.tab-line {
    grid-column: 1;
    grid-row: 1;
}

.view_app {
    display: grid;
    grid-template-columns: 185px 1fr;
    grid-column-gap: 15px;
    height: 100%;
}

.app_list {
    background: #1E1E1E;
    border-radius: 8px;
    padding: 15px;
    height: max-content;
    min-height: 225px;
    max-height: 100%;
    overflow-y: scroll;
    border: 1px solid #30363D;
}

.app_name_list {
    margin-top: 25px;
    padding: 0 7px;
}
.app-name {
    background:linear-gradient(135deg, #50e3c2 0%,#10d574 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent; 
    font-size: 16px;
    font-weight: bold;
    padding: 5px;
    margin: 5px 0;
    border: 1px solid #10d574;
    border-radius: 8px;
    border-style: dashed;
}
.app-name:hover {
    cursor: pointer;
    text-decoration: underline;
}

</style>
