<template>
    <div class="">
        <br>
        <h1 class="d-flex justify-between">Funnel Configuration</h1>
        <small><span class="link" @click="showCfg('funnel')">how does it work?</span></small>
        <div class="view_component funnel_view mt-0 pl-0">
            <div class="d-flex align-center justify-end">
                    <button class="btn btn-standard" @click="updateStages">Update</button>
                </div>
            <div class="">
                <div class="form-col col d-flex flex-wrap">
                    <div v-for="f in funnel" :key="f.id" class="d-flex align-center m-1">
                        <div class="funnel">
                            <div class="d-flex justify-end trash-span">
                                <span v-if="f.id >= funnel.length" class="icon icon-trash-2 hover" @click="removeStage(f.id)"></span>
                            </div>
                            <div class="d-flex justify-center align-center flex-col">
                                <div class="stage-name">{{f.name}}</div>
                                <div class="stage-transition">{{f.transition}}</div>
                            </div>
                        </div>
                        <div>
                             <span v-if="f.id < funnel.length" class="icon icon-chevron-right super"></span>
                         </div>
                    </div>
                     <div class="funnel add-box d-flex align-center justify-even">
                         <div class="d-flex align-center">
                             <span class="icon icon-chevrons-right super"></span>
                         </div>
                         <div>
                             <div class="">
                                <div class=" col">
                                    <input v-model="stage_name" class="form-control border" type="text" name="stage_name" id="stage_name" placeholder="Stage Name" :class="{'border-danger': stage_invalid}" >
                                </div>
                            </div>
                            <div class="mt-1">
                                <div class=" col">
                                    <input v-model="stage_transition" class="form-control border" type="text" name="stage_tansition" id="stage_tansition" placeholder="Transition" :class="{'border-danger': stage_invalid}" >
                                </div>
                            </div>
                         </div>
                         <div>
                             <span class="icon icon-plus hover super" @click="addStage"></span>
                         </div>
                    </div>
                </div>
            </div>
        </div>
        <br>
        <h1>Campaign Tracking</h1>
        <small> <span class="link" @click="showCfg('campaign')">how does it work?</span></small>
        <div class="view_component table-height">
            <div class="d-flex align-center justify-end">
                <button class="btn btn-standard" @click="updateCampaigns">Update</button>
            </div>
            <table class="table table-borderless">
            <thead>
                <tr>
                    <th>#</th>
                    <th>Campaign Name</th>
                    <th>URL Prefix</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in campaign" :key="item.id">
                    <th>{{item.id}}</th>
                    <th>{{item.name}}</th>
                    <th>{{item.prefix}}</th>
                    <th><span class="icon icon-trash-2 hover" @click="removeCampaign(item.id)"></span></th>
                </tr>
                <tr>
                    <th class="v-center">{{campaign_count}}</th>
                    <td><input v-model="campaign_name" type="text" placeholder="Name (E-Mail Campaign)" class="form-control border" :class="{'border-danger': campaign_invalid}" ></td>
                    <td><input v-model="campaign_prefix" type="text" placeholder="Prefix (ex. summer-sales)" class="form-control border" :class="{'border-danger': campaign_invalid}" ></td>
                    <td class="v-center"><span class="icon icon-plus hover" @click="addCampaign"></span></td>
                </tr> 
            </tbody>
            </table>
        </div>
        <br>
        <h1>Interesting Buttons</h1>
        <small><span class="link" @click="showCfg('btn_time')">how does it work?</span></small>
        <div class="view_component table-height">
            <div class="d-flex align-center justify-end">
                <button class="btn btn-standard" @click="updateBtnTime">Update</button>
            </div>
            <table class="table table-borderless">
            <thead>
                <tr>
                    <th>#</th>
                    <th>Name</th>
                    <th>Button Name</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in buttons" :key="item.id">
                    <th>{{item.id}}</th>
                    <th>{{item.name}}</th>
                    <th>{{item.btn_name}}</th>
                    <th><span class="icon icon-trash-2 hover" @click="removeBtnTime(item.id)"></span></th>
                </tr>
                <tr>
                    <th class="v-center">{{button_count}}</th>
                    <td><input v-model="button_name" type="text" placeholder="Name (ex. Btn-Order)" class="form-control border" :class="{'border-danger': button_invalid}" ></td>
                    <td><input v-model="button_btn" type="text" placeholder="Btn (ex. btn_order)" class="form-control border" :class="{'border-danger': button_invalid}" ></td>
                    <td class="v-center"><span class="icon icon-plus hover" @click="addBtnTime"></span></td>
                </tr> 
            </tbody>
            </table>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    name: "Configuration",
    data() {
        return {
            funnel: [],
            stage_name: null,
            stage_transition: null,
            funnel_count: 0,
            stage_invalid: false,

            campaign: [],
            campaign_name: null,
            campaign_prefix: null,
            campaign_count: 0,
            campaign_invalid: false,

            buttons: [],
            button_name: null,
            button_btn: null,
            buttons_count: 0,
            button_invalid: false,
            unsaved_changes: false,
        };
    },
    props: {
        config_uuid: String,
        app_config: {
            type: Object,
            default: null,
        },
        app_uuid: {
            type: String,
            default: null,
        },

    },
    mounted() {
        console.log("Configs", this.$props.app_config);
        if (this.$props.app_config === null || Object.keys(this.$props.app_config).length === 0){
            this.funnel = [];
            this.funnel_count = 1;
            this.campaign = [];
            this.campaign_count = 1;
            this.buttons = [];
            this.buttons_count = 1;
            return
        }
        this.funnel = this.$props.app_config.funnel ? this.$props.app_config?.funnel: [];
        this.funnel_count = this.funnel.length + 1;

        this.campaign = this.$props.app_config.campaign ? this.$props.app_config?.campaign: [];
        this.campaign_count = this.campaign.length + 1;

        this.buttons = this.$props.app_config.btn_time ? this.$props.app_config?.btn_time: [];
        this.buttons_count = this.buttons.length + 1;
    },
    computed: {
        unsavedChanges() {
            console.log(this.unsaved_changes);
        }
    },
    methods: {
        addStage() {
            if (this.stage_invalid) this.stage_invalid = false;
            const tmp = this.funnel.filter(item => item.name === this.stage_name || item.transition === this.stage_transition)
            if (tmp.length > 0) {
                this.$toast.warning("Funnel Name and Transition must be unique");
                this.stage_invalid = true;
                return
            }

            this.funnel.push({
                id: this.funnel_count,
                name: this.stage_name,
                transition: this.stage_transition,
            });
            this.stage_name = null;
            this.stage_transition = null;
            this.funnel_count++;
            this.$emit("appchange", true);
        },
        removeStage(id) {
            console.log(id);
            this.funnel = this.funnel.filter(item => item.id != id);
            this.funnel.forEach((item,i) => {
                item.id = i+1;
            });
            this.funnel_count--;
            this.$emit("appchange", true);
        },
        updateStages() {
            let options = {
                headers: {
                    'Authorization': localStorage.getItem("token"),
                }
            }; 
            const payload = {
                app_uuid: this.$props.app_uuid,
                stages: this.funnel,
            }
            axios.post("http://localhost:8080/api/v1/app/config/upsert?flag=funnel", payload, options).then(res => {
                console.log(res);
                this.$toast.success("Updated Funnel information");
                this.$emit("appchange", false);
            }).catch(err => this.$toast.error(err.response.data));
        },
        addCampaign() {
            if (this.campaign_invalid) this.campaign_invalid = false;
            const tmp = this.campaign.filter(item => item.name === this.campaign_name || item.prefix === this.campaign_prefix)
            if (tmp.length > 0) {
                this.$toast.warning("Campaign Name and Prefix must be unique");
                this.campaign_invalid = true;
                return
            }
            this.campaign.push(
                {
                    id: this.campaign_count,
                    name: this.campaign_name,
                    prefix: this.campaign_prefix,
                }
            );
            this.campaign_name = null;
            this.campaign_prefix = null;
            this.campaign_count++;
            this.$emit("appchange", true);
        },
        removeCampaign(id) {
            this.campaign = this.campaign.filter(item => item.id != id);
            this.campaign.forEach((item,i) => {
                item.id = i+1;
            });
            this.campaign_count--;
            this.$emit("appchange", true);
        },
        updateCampaigns() {
            let options = {
                headers: {
                    'Authorization': localStorage.getItem("token"),
                }
            }; 
            const payload = {
                app_uuid: this.$props.app_uuid,
                records: this.campaign,
            }
            axios.post("http://localhost:8080/api/v1/app/config/upsert?flag=campaign", payload, options).then(res => {
                console.log(res);
                this.$toast.success("Updated Campaign information");
                this.$emit("appchange", false);
            }).catch(err => this.$toast.error(err.response.data));
        },

        addBtnTime() {
            if (this.button_invalid) this.button_invalid = false;
            const tmp = this.buttons.filter(item => item.name === this.button_name || item.btn_name === this.button_name);
            if (tmp.length > 0) {
                this.$toast.warning("Button Name and Identifier must be unique");
                this.button_invalid = true;
                return
            }

            this.buttons.push({
                id: this.buttons_count,
                name: this.button_name,
                btn_name: this.button_btn,
            })

            this.button_name = null;
            this.button_btn = null;
            this.buttons_count++;
            this.$emit("appchange", true);
        },
        removeBtnTime(id) {
            this.buttons = this.buttons.filter(item => item.id != id);
            this.buttons.forEach((item,i) => {
                item.id = i+1;
            });
            this.buttons_count--;
            this.$emit("appchange", true);
        },
        updateBtnTime() {
            let options = {
                headers: {
                    'Authorization': localStorage.getItem("token"),
                }
            }; 
            const payload = {
                app_uuid: this.$props.app_uuid,
                btn_defs: this.buttons,
            }
            axios.post("http://localhost:8080/api/v1/app/config/upsert?flag=btn_time", payload, options).then(res => {
                console.log(res);
                this.$toast.success("Updated Interesting-Buttons information");
                this.$emit("appchange", false);
            }).catch(err => this.$toast.error(err.response.data));
        },
        showCfg(type) {
            console.log(type);
            this.$emit("setdoc", type);
        },
    },
}
</script>

<style sceped>
.main-cfg {
    height: 70vh;
    overflow-y: auto;
}
.view_component {
    margin-top: 15px;
    padding: 15px;
    border-radius: 8px;
    height: max-width;
}

.funnel_view {
    background: transparent;
    border: none;
}

.table-height {
    max-height: 300px;
    overflow-y: auto;
}

.v-center {
    vertical-align: middle !important;
}

td .icon {
    font-size: 22px;
}
.funnel {
    padding: 10px;
    width: auto;
    height: 100px;
    min-width: 150px;
    max-width: 200px;
    background: var(--sub-bg);
    border: 1px solid var(--sub-border);
    border-radius: 8px;
    margin: 0 5px;
}

.funnel .stage-name {
    font-size: 18px;
    color: var(--h-color);
}
.funnel .stage-transition {
    font-size: 14px;
    color: var(--txt-small);
}

.add-box {
    opacity: 0.5;
    border: none;
    width: 210px;
    max-width: 210px;
    background: none;
}
.add-box:focus,.add-box:hover {
    opacity: 1;
}

.trash-span {
    height: 14px;
}

.unsaved-changes {
    border: 1px solid orange;
}
</style>