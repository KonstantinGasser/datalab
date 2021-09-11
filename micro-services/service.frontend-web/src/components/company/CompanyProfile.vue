<template>
    <div class="main-view">
        <Tabs ref="Tabs" :class="{ block: blockTabs }" :initTab="'Profile'" :tabs="[{name:'Profile'},{name:'Subscription'},{name:'User'}]" @tabChange="tabChange"/>
        <TabProfile v-if="activeTab === 'Profile'" @inEdit="enableDisableTabs"/>
    </div>
</template>

<script>
    import Tabs from '@/components/utils/Tabs.vue';
    import TabProfile from '@/components/company/tabs/TabProfile.vue';

    export default {
        name: 'CompanyProfile',
        components: {
            Tabs,
            TabProfile,
        },
        computed: {
            blockTabs() {
                return this.tabsBlocked;
            }
        },
        data() {
            return {
                activeTab: 'Profile',
                tabsBlocked: false,
            };
        },
        props: ['status'],
        methods: {
            tabChange(tab) {
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


</style>
