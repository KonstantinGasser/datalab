<template>
    <div class="view_component d-flex">
        <div class="form-row w-100">
           <div class="form-group col">
                <div class="form-group col">
                    <label for="">App Name</label>
                    <input v-model="appName" type="text" class="form-control" name="" id="app_name" placeholder="try something meaningful">
                </div>
                <div class="form-group col">
                    <label for="">App URL</label>
                    <input v-model="appURL" type="text" class="form-control" name="" id="app_url" placeholder="https://awesome.app.dev">
                    <small>We need to know the URL in order to verify that the client is indeed allowed to provide user data</small>
                </div>
            </div>
            <div class="form-group col">
                <label for="">App Description</label>
               <textarea v-model="appDesc" class="form-control" name="" id="app_desc" rows="2" placeholder="what is the app about?"></textarea>
            </div>
        </div>
    </div>
    <div class="view_component">
        
        <!-- <div class="form-row">
            <div class="form-group col">
                <label for="">What do you want to monitor</label>
                <div class="d-flex justify-between align-center">
                    <div class="custom-control custom-switch">
                    <input v-model="appCfgs" :value="'mouse-move-map'" @change="setConfig($event)" type="checkbox" class="custom-control-input" id="mouse-move-map">
                    <label class="custom-control-label" for="mouse-move-map">Mouse Movements</label>
                </div>
                <div class="custom-control custom-switch">
                    <input v-model="appCfgs" :value="'customer_journey'" @change="setConfig($event)" type="checkbox" class="custom-control-input" id="customer_journey">
                    <label class="custom-control-label" for="customer_journey">Customer Journey</label>
                </div>
                </div>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col">
                <div class="d-flex justify-evenly align-center">
                    <div class="custom-control custom-switch">
                    <input v-model="appCfgs" :value="'heat_map'" @change="setConfig($event)" type="checkbox" class="custom-control-input" id="heat_map">
                    <label class="custom-control-label" for="heat_map">Heat-Map of mouse movements</label>
                </div>
                <div class="custom-control custom-switch">
                    <input v-model="appCfgs" :value="'customer_journey'" @change="setConfig($event)" type="checkbox" class="custom-control-input" id="customer_journey">
                    <label class="custom-control-label" for="customer_journey">Customer Journey</label>
                </div>
                </div>
            </div>
        </div> -->
    </div>
    <div class="view_component">
        <div class="form-row">
            <div class="from-group col-6">
                <label for="">Add Colleagues to the App</label>
                <input type="text" name="" id="search_member_field" class="form-control" placeholder="Search">
            </div>
            <div class="from-group col d-flex flex-wrap justify-center align-center">
                <div class="user_card">
                    <div class="d-flex align-center justify-end">
                        <div @click="addMemberToApp(dummy)">
                            <span class="icon icon-user-plus hover big"></span>
                        </div>
                    </div>
                    <div class="d-flex align-center justify-center">
                        <img class="circle-img medium" src="http://www.expertyou.de:8080/member/expert/266/profile/photo_266_1604926599.jpeg" alt="">
                    </div>
                    <div class="member_info">
                        <span class="member_name dots">Konstantin Gasser</span>
                        <span class="member_pos dots">Software Engineer</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="d-flex justify-end align-center mt-2">
        <button class="btn btn-standard" @click="createNewApp()">Create App {{ appName }}</button>
    </div>
</template>

<script>

    import axios from 'axios';

    export default {
        name: 'TabCreateApp',
        components: {},
        data() {
            return {
                isEdit: false,
                appName: null,
                appURL: null,
                appDesc: null,
                appMember: [],
                appCfgs: [],
            };
        },
        methods: {
            setConfig(event) {
                if(!event.checked) {
                    this.appCfgs.filter(item => {item !== event.defaultValue});
                    return
                }
                this.appCfgs.push(event.defaultValue);
            },
            setMember(member) {
                this.appMember.push(member);
            },
            setMode() {
                this.isEdit = !this.isEdit;
                // emit panel is in edit mode
                // diable tabs until saved
                this.$emit('inEdit', this.isEdit);
            },
            async createNewApp (){
                
                let formValid = true;
                if (this.appURL === null || this.appURL === '') {
                    this.$toast.error('App URL is required');
                    formValid = false;
                }
                if (!this.validURL(this.appURL)) {
                    this.$toast.error('Mhm does not look like an URL does it ?');
                    formValid = false;
                }
                if (this.appName === null || this.appName === '') {
                    this.$toast.error('App Name is required');
                    formValid = false;
                }
                if (this.appDesc === null || this.appDesc === '') {
                    this.$toast.error('App Description is required');
                    formValid = false;
                }
                if (this.appCfgs.length === 0) {
                    this.$toast.warning("If no App Configs are set the app will not generate any data...");
                }
                
                if (formValid) {
                    let options = {
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': localStorage.getItem("token"),
                        }
                    };
                    await axios.post("http://localhost:8080/api/v2/view/app/create", {
                            app_name: this.appName,
                            app_description: this.appDesc,
                            app_member: this.appMember,
                            app_settings: this.appCfgs,
                            app_url: this.appURL,
                        }, options
                    ).then((resp) => {
                        console.log("Resp ",resp);
                        this.$toast.success("App " + this.appName + " has been created");
                        this.$emit("createdApp", {"type": "show_app", "app_uuid": resp.data.app_uuid});
                    }).catch(err => {
                        this.$toast.error("Sorry app could not be created");
                        return;
                    });
                }
            },
            validURL(str) {
                var pattern = new RegExp('^(https?:\\/\\/)?'+ // protocol
                    '((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|'+ // domain name
                    '((\\d{1,3}\\.){3}\\d{1,3}))'+ // OR ip (v4) address
                    '(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*'+ // port and path
                    '(\\?[;&a-z\\d%_.~+=-]*)?'+ // query string
                    '(\\#[-a-z\\d_]*)?$','i'); // fragment locator
                return !!pattern.test(str);
            },
        },
    };
</script>

<style scoped>
.view_component {
    margin-bottom: 15px;
    padding: 15px;
    background: #1E1E1E;
    border-radius: 8px;
    height: max-content;
    border: 1px solid #30363D;
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

.test {
    padding: 5px;
    background: lightgray;
    height: 100%;
    width: auto;
}

.selected_member_list {

}
.selected_member {
    padding: 2px 5px;
    border-radius: 8px;
    background: rgba(0,0,0,0.1);
}
</style>
