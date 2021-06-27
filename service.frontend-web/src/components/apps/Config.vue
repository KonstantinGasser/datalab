<template>
    <div class="">
        <br>
        <h1 class="d-flex justify-between">Funnel Configuration {{app_locked}}</h1>
        <small><span class="link" @click="showCfg('funnel')">how does it work?</span></small>
        <div class="view_component funnel_view mt-0 pl-0 dash">
            <div class="d-flex align-center justify-end">
                    <button v-if="app_locked === undefined || !app_locked" class="btn btn-standard" @click="updateStages">Save</button>
                </div>
            <div class="">
                <div class="form-col col d-flex flex-wrap">
                    <div v-for="f in app_config?.funnel" :key="f.id" class="d-flex align-center m-1">
                        <div class="funnel">
                            <div class="d-flex justify-end trash-span">
                                <span v-if="f.id >= app_config?.funnel?.length && (app_locked === undefined || !app_locked)" class="icon icon-trash-2 hover" @click="removeStage(f.id)"></span>
                            </div>
                            <div class="d-flex justify-center align-center flex-col">
                                <div class="stage-name dots">{{f.name}}</div>
                                 <div class="stage-trigger">
                                    <div v-if="f.trigger == 1" class="dots"><span class="icon icon-at-sign standard-font"></span>URL</div>
                                    <div v-if="f.trigger == 2" class="dots"><span class="icon icon-at-sign standard-font"></span>OnClick</div>
                                </div>
                                <div class="stage-transition tooltip1">
                                    <div class="dots">{{f.transition}}{{f.regex}}</div>
                                    <span class="tooltiptext1">{{f.transition}}</span>
                                </div>
                            </div>
                        </div>
                        <div>
                             <span v-if="f.id < app_config?.funnel?.length" class="icon icon-chevron-right super"></span>
                         </div>
                    </div>
                     <div v-if="app_locked === undefined || !app_locked" class="funnel add-box d-flex align-center justify-even">
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
                                    <select name="" id="" class="custom-select" v-model="stage_trigger" :class="{'border-danger': stage_invalid}" >
                                        <option v-for="item in stage_triggers" :key="item.id" v-bind:value="{id: item.id, value: item.value}">
                                            {{ item.value }}
                                        </option>
                                    </select>
                                </div>
                            </div>
                            <div class="mt-1">
                                <div class=" col">
                                    <input v-model="stage_transition" class="form-control border" type="text" name="stage_tansition" id="stage_tansition" placeholder="Transition" :class="{'border-danger': stage_invalid}" >
                                </div>
                            </div>
                         </div>
                         <div>
                             <span v-if="app_locked === undefined || !app_locked" class="icon icon-plus hover super" @click="addStage"></span>
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
                <button v-if="app_locked === undefined || !app_locked" class="btn btn-standard" @click="updateCampaigns">Save</button>
            </div>
            <table class="table table-borderless">
            <thead>
                <tr>
                    <th>#</th>
                    <th>Campaign Name</th>
                    <th>URL Suffix</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in app_config?.campaign" :key="item.id">
                    <th>{{item.id}}</th>
                    <th>{{item.name}}</th>
                    <th>{{item.suffix}}</th>
                    <th><span v-if="app_locked === undefined || !app_locked"  class="icon icon-trash-2 hover" @click="removeCampaign(item.id)"></span></th>
                </tr>
                <tr v-if="app_locked === undefined || !app_locked">
                    <th class="v-center"></th>
                    <td><input v-model="campaign_name" type="text" placeholder="Name (E-Mail Campaign)" class="form-control border" :class="{'border-danger': campaign_invalid}" ></td>
                    <td>
                        <div class="input-group">
                            <div class="input-group-prepend">
                                <button class="btn  btn-standard btn-disabled" disabled type="button"><span class="icon icon-hash"></span></button>
                            </div>
                            <input v-model="campaign_suffix" type="text" placeholder="Suffix (ex. summer-sales)" class="form-control border" :class="{'border-danger': campaign_invalid}" >
                        </div>
                        
                    </td>
                    <td class="v-center"><span v-if="app_locked === undefined || !app_locked" class="icon icon-plus hover" @click="addCampaign"></span></td>
                </tr> 
            </tbody>
            </table>
        </div>
        <br>
        <h1>Interesting Buttons</h1>
        <small><span class="link" @click="showCfg('btn_time')">how does it work?</span></small>
        <div class="view_component table-height">
            <div class="d-flex align-center justify-end">
                <button v-if="app_locked === undefined || !app_locked" class="btn btn-standard" @click="updateBtnTime">Save</button>
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
                <tr v-for="item in app_config?.btn_time" :key="item.id">
                    <th>{{item.id}}</th>
                    <th>{{item.name}}</th>
                    <th>{{item.btn_name}}</th>
                    <th><span v-if="app_locked === undefined || !app_locked" class="icon icon-trash-2 hover" @click="removeBtnTime(item.id)"></span></th>
                </tr>
                <tr v-if="app_locked === undefined || !app_locked">
                    <th class="v-center"></th>
                    <td><input v-model="button_name" type="text" placeholder="Name (ex. Btn-Order)" class="form-control border" :class="{'border-danger': button_invalid}" ></td>
                    <td><input v-model="button_btn" type="text" placeholder="Btn (ex. btn_order)" class="form-control border" :class="{'border-danger': button_invalid}" ></td>
                    <td class="v-center"><span v-if="app_locked === undefined || !app_locked" class="icon icon-plus hover" @click="addBtnTime"></span></td>
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
            stage_invalid: false,
            stage_name: null,
            stage_transition: null,
            stage_triggers: [{id: 1, value: "URL"}, {id: 2, value: "OnClick"}],
            stage_trigger: {id: 2, value: "OnClick"},

            campaign_invalid: false,
            campaign_name: null,
            campaign_suffix: null,

            button_invalid: false,
            button_name: null,
            button_btn: null,
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
        app_locked: {
            type: Boolean,
            default: true,
        }
    },
    mounted() {

    },
    computed: {

    },
    methods: {
        isRegex(pattern) {
            var parts = pattern.split('/'),
                regex = pattern,
                options = "";
            if (parts.length > 1) {
                regex = parts[1];
                options = parts[2];
            }
            try {
                new RegExp(regex, options);
                return true;
            }
            catch(e) {
                return false;
            }
        },
        addStage() {
            
            if (this.stage_invalid) this.stage_invalid = false;
            // check for naming conflict (values must be unique)
            const tmp = this.$props.app_config?.funnel?.filter(item => item.name === this.stage_name || item.transition === this.stage_transition)
            if (tmp?.length > 0) {
                this.$moshaToast("Funnel Name and Transition must be unique", {type: 'danger',position: 'top-center', timeout: 3000})
                this.stage_invalid = true;
                return
            }

            // selecting and checking for state trigger 0,1,2
            const tmp2 = this.stage_triggers.filter(item => item.id === this.stage_trigger.id)
            if (tmp2?.length === 0) {
                this.$moshaToast("Please select a Stage-Trigger", {type: 'danger',position: 'top-center', timeout: 3000})
            }
            // extract regex pattern from transition
            const regex_pattern = this.stage_transition.substring(
                this.stage_transition.search("{")+1,this.stage_transition.search("}")
            )
            console.log("Regex: ", regex_pattern)
            console.log("stage transition: ", this.stage_transition)
            if (regex_pattern?.length > 0 && !this.isRegex(regex_pattern)) {
                this.$moshaToast("Provided Regex might be wrong", {type: 'warning',position: 'top-center', timeout: 3000})
            }
            let count = this.$props.app_config?.funnel?.length + 1
            if (Number.isNaN(count)) {
                count = 1
            }
            this.$emit("appchange", {unsaved: true, type: "funnel-add", item: {
                    id: count,
                    name: this.stage_name,
                    transition: this.stage_transition.replace("{"+regex_pattern+"}", ""), // cut-off regex
                    regex: regex_pattern,
                    trigger: this.stage_trigger.id,
                }
            })
            this.stage_transition = null;
            this.stage_name = null;
            this.stage_trigger = {id: 2, value: "OnClick"};
        },
        removeStage(id) {
            this.$emit("appchange", {unsaved: true, type: "funnel-remove", item: id}) 
        },
        updateStages() {
            let options = {
                headers: {
                    'Authorization': localStorage.getItem("token"),
                }
            }; 
            const payload = {
                flag: "funnel",
                app_uuid: this.$props.app_uuid,
                stages: this.$props.app_config?.funnel,
            }
            axios.post("http://localhost:8080/api/v1/app/config/update", payload, options).then(res => {
                // this.$toast.success("Updated Funnel information");
                this.$moshaToast(res.data.msg, {type: 'success',position: 'top-center', timeout: 3000})
                this.$emit("appchange", {unsaved: false, type: "funnel-saved"});

            }).catch(err => this.$moshaToast(err.response?.data?.msg, {type: 'danger',position: 'top-center', timeout: 3000}));
        },
        addCampaign() {
            if (this.campaign_invalid) this.campaign_invalid = false;
            const tmp = this.$props.app_config?.campaign?.filter(item => item.name === this.campaign_name || item.suffix === this.campaign_suffix)
            if (tmp?.length > 0) {
                this.$toast.warning("Campaign Name and Prefix must be unique");
                this.campaign_invalid = true;
                return
            }
            let count = this.$props.app_config?.campaign?.length + 1
            if (Number.isNaN(count)) {
                count = 1
            }
            this.$emit("appchange", {unsaved: true, type: "campaign-add", item: {
                    id: count,
                    name: this.campaign_name,
                    suffix: this.campaign_suffix,
                }
            }) 
        },
        removeCampaign(id) {
            this.$emit("appchange", {unsaved: true, type: "campaign-remove", item: id})
        },
        updateCampaigns() {
            let options = {
                headers: {
                    'Authorization': localStorage.getItem("token"),
                }
            }; 

            const payload = {
                flag:"campaign",
                app_uuid: this.$props.app_uuid,
                records: this.$props.app_config?.campaign,
            }
            axios.post("http://localhost:8080/api/v1/app/config/update", payload, options).then(res => {
                this.$moshaToast(res.data.msg, {type: 'success',position: 'top-center', timeout: 3000})
                this.$emit("appchange", {unsaved: false, type: "campaign-saved"});

            }).catch(err => this.$moshaToast(err.response?.data?.msg, {type: 'danger',position: 'top-center', timeout: 3000}));
        },

        addBtnTime() {
            if (this.button_invalid) this.button_invalid = false;
            const tmp = this.$props.app_config?.btn_time?.filter(item => item.name === this.button_name || item.btn_name === this.button_btn);
            if (tmp?.length > 0) {
                this.$toast.warning("Button Name and Identifier must be unique");
                this.button_invalid = true;
                return
            }

            let count = this.$props.app_config?.btn_time?.length + 1
            if (Number.isNaN(count)) {
                count = 1
            }
            this.$emit("appchange", {unsaved: true, type: "btn-add", item: {
                    id: count,
                    name: this.button_name,
                    btn_name: this.button_btn,
                }
            }) 
        },
        removeBtnTime(id) {
            this.$emit("appchange", {unsaved: true, type: "btn-remove", item: id})
        },
        updateBtnTime() {
            let options = {
                headers: {
                    'Authorization': localStorage.getItem("token"),
                }
            }; 
            console.log("DATA: ",this.$props.app_config?.btn_time)
            const payload = {
                flag:"btntime",
                app_uuid: this.$props.app_uuid,
                btn_defs: this.$props.app_config?.btn_time,
            }
            axios.post("http://localhost:8080/api/v1/app/config/update", payload, options).then(res => {
                this.$moshaToast(res.data.msg, {type: 'success',position: 'top-center', timeout: 3000})
                this.$emit("appchange", {unsaved: false, type: "btn-saved"})
            }).catch(err => this.$moshaToast(err.response?.data?.msg, {type: 'danger',position: 'top-center', timeout: 3000}));
        },
        showCfg(type) {
            this.$emit("setdoc", type);
        },
    },
}
</script>

<style scoped>
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

.dash {
    border: 1px solid #ccc !important;
    border-style: dashed !important;
}
.funnel_view {
    background: transparent;
    border: none;
}

.table-height {
    max-height: 300px;
    overflow-y: auto;
}
.btn-disabled {
    border-radius: 4px;
    background: var(--menu-bg);
    color: var(--btn-font-hover);
    opacity: 1;
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
    max-width: 220px;
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
    text-align: center;
    width: 100%;
    font-size: 14px;
    color: var(--txt-small);
}

.add-box {
    opacity: 0.5;
    border: none;
    width: 230px;
    max-width: 230px;
    height: 150px;
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