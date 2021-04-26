<template>
    <div class="main-cfg">
        <h1>Funnel Configuration</h1>
        <div class="view_component">
            <div class="form-row">
                <div class="form-col col">
                    <!-- <small class="info_txt"> -->
                        In order to track which customer acts in which stage of the funnel,
                        you must provide information about the stage and their transitions to the next stage 😬
                    <!-- </small> -->
                </div>
            </div>
            <div class="form-row mt-2">
                <div class="form-col col">
                    <strong>"Stage Name"</strong>: is the name you want the stage to have, it has no effect on the logic 
                    <br>
                    <strong>"Transition"</strong>: this is <strong>important</strong>. Only with the transition we are able to track when
                    a customer jumps into the next stage <small>(rn "Transition" must be the name of the HTML-Element)</small> 
                </div>
            </div>
            <div class="mt-3">
                <div class="d-flex align-center justify-end mb-2">
                    <button class="btn btn-standard" @click="updateStages">Update</button>
                </div>
                <div class="form-col col d-flex flex-wrap">
                    <div v-for="f in funnel" :key="f.id" class="d-flex align-center m-1">
                        <div class="funnel">
                            <div class="d-flex justify-end trash-span">
                                <span v-if="f.id >= funnel.length - 1" class="icon icon-trash-2 hover" @click="removeStage(f.id)"></span>
                            </div>
                            <div class="d-flex justify-center align-center flex-col">
                                <div class="stage-name">{{f.name}}</div>
                                <div class="stage-transition">{{f.transition}}</div>
                            </div>
                        </div>
                        <div>
                             <span v-if="f.id < funnel.length - 1" class="icon icon-chevron-right super"></span>
                         </div>
                    </div>
                     <div class="funnel add-box d-flex align-center justify-even">
                         <div class="d-flex align-center">
                             <span class="icon icon-chevrons-right super"></span>
                         </div>
                         <div>
                             <div class="">
                                <div class=" col">
                                    <input v-model="stage_name" class="form-control" type="text" name="stage_name" id="stage_name" placeholder="Stage Name">
                                </div>
                            </div>
                            <div class="mt-1">
                                <div class=" col">
                                    <input v-model="stage_transition" class="form-control" type="text" name="stage_tansition" id="stage_tansition" placeholder="Transition">
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
        <div class="view_component table-height">
            <div class="d-flex align-center justify-end mb-2">
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
                    <td><input v-model="campaign_name" type="text" placeholder="Name (E-Mail Campaign)" class="form-control"></td>
                    <td><input v-model="campaign_prefix" type="text" placeholder="Prefix (#prop-email)" class="form-control"></td>
                    <td class="v-center"><span class="icon icon-plus hover" @click="addCampaign"></span></td>
                </tr> 
            </tbody>
            </table>
        </div>
        <br>
        <h1>Interesting Buttons</h1>
        <div class="view_component table-height">
            <div class="d-flex align-center justify-end mb-2">
                <button class="btn btn-standard" @click="updateCampaigns">Update</button>
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
                    <th>{{item.btn}}</th>
                    <th><span class="icon icon-trash-2 hover"></span></th>
                </tr>
                <tr>
                    <th class="v-center">{{button_count}}</th>
                    <td><input v-model="button_name" type="text" placeholder="Name (Btn-Order)" class="form-control"></td>
                    <td><input v-model="button_btn" type="text" placeholder="Btn (btn_order)" class="form-control"></td>
                    <td class="v-center"><span class="icon icon-plus hover"></span></td>
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
            stage_name: null,
            stage_transition: null,
            stage_count: 0,
            funnel: [],
            campaign_name: null,
            campaign_prefix: null,
            campaign_count: 0,
            campaign: [],
            buttons: [],
            button_count: 0,
            button_name: null,
            button_btn: null,
        };
    },
    props: ["app_uuid"],
    computed: {
    },
    methods: {
        addStage() {
            this.funnel.push({
                id: this.stage_count,
                name: this.stage_name,
                transition: this.stage_transition,
            });
            this.stage_name = null;
            this.stage_transition = null;
            this.stage_count++;
        },
        removeStage(id) {
            this.funnel = this.funnel.filter(item => item.id != id);
            this.funnel.forEach((item,i) => {
                item.id = i;
            });
            this.funnel_count--;
        },
        updateStages() {
            let options = {
                headers: {
                    'Authorization': localStorage.getItem("token"),
                }
            }; 
            const payload = {
                app_uuid: this.$props.app_uuid,
                funnel: this.funnel,
            }
            axios.post("http://localhost:8080/api/v2/view/app/update/config?resource=funnel", payload, options).then(res => {
                console.log(res);
            }).catch(err => console.log(err));
        },
        addCampaign() {
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
        },
        removeCampaign(id) {
            this.campaign = this.campaign.filter(item => item.id != id);
            this.campaign.forEach((item,i) => {
                item.id = i;
            });
            this.campaign_count--;
        },
        updateCampaigns() {
            let options = {
                headers: {
                    'Authorization': localStorage.getItem("token"),
                }
            }; 
            const payload = {
                app_uuid: this.$props.app_uuid,
                campaign: this.campaign,
            }
            axios.post("http://localhost:8080/api/v2/view/app/update/config?resource=campaign", payload, options).then(res => {
                console.log(res);
            }).catch(err => console.log(err));
        },
    },
}
</script>

<style sceped>
.main-cfg {
    height: 700px;
    overflow-y: scroll;
}
.view_component {
    margin-top: 15px;
    padding: 15px;
    border-radius: 8px;
    height: max-width;
}

.table-height {
    max-height: 300px;
    overflow-y: scroll;
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
    width: 220px;
    max-width: 220px;
}
.add-box:focus,.add-box:hover {
    opacity: 1;
}


.trash-span {
    height: 14px;
}
</style>