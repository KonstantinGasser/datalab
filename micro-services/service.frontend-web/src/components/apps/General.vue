<template>
    <div>
        <br>
        <h1>App Token</h1>
        <div class="view_component">
            <div class="form-row">
                <div class="form-group col">
                    <label for="">Generate <br>
                        <small>
                            generating an App-Token will lead to this App being in a locked state not allowing to make change
                            to any configurations.
                        </small>
                    </label>
                    <div class="input-group">
                        <input v-model="verification_step" type="text" class="form-control" placeholder="Domain/App-Name" aria-label="" aria-describedby="basic-addon1">
                        <div class="input-group-append">
                            <button class="btn btn-standard" style="width:65px;" @click="generateToken()" type="button"><span class="icon icon-tag"></span></button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="view_component">
            <div class="form-row">
                <div class="form-group col">
                    <label for="">Your App Token</label>
                    <div class="input-group">
                        <textarea type="text" class="form-control" rows="4" readonly id="app_token_value" :value="jwt" aria-label="" aria-describedby="basic-addon1"></textarea>
                        <div class="input-group-append">
                            <button class="btn btn-standard" style="width:65px;" @click="copyTokenToClipboard()" type="button"><span class="icon icon-clipboard"></span></button>
                        </div>
                    </div>
                    <div class=""><small v-if="expTimeSet">Token expires in {{expTimeSet?.days}} days {{expTimeSet?.hours}} hours</small></div>
                    <div class="mt-3">
                        Checkout the <a href="http://192.168.0.177:3000/docs/lib" target="_blank">documentation</a> 
                        on how to implement the client side
                    </div>
                </div>
            </div>
        </div>
        <br>
        <h1 v-if="loggedInUser?.sub === app_owner">Dangerous Water 🙀</h1>
        <div v-if="app_locked && loggedInUser?.sub === app_owner" class="view_component">
            <div class="form-row">
                <div class="form-group col">
                    <label for="" class="danger-text">Unlock App<br>
                        <small>unlocking the application will invalidate the current App-Token, purging the data generated to fare
                        </small>
                    </label>
                    <div class="input-group">
                        <input  type="text" class="form-control" placeholder="Domain/AppName" aria-label="" aria-describedby="basic-addon1">
                        <div class="input-group-append">
                            <button class="btn btn-standard" style="width:65px;" @click="unlockApp(app_uuid)" type="button">
                                <span class="icon icon-unlock"></span>
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div v-if="loggedInUser?.sub === app_owner" class="view_component">
            <div class="form-row">
                <div class="form-group col">
                    <label for="" class="danger-text">Delete App<br>
                        <small>(Your-Orgnaization-Domain/App-Name) 
                            <br>
                            We just want to make sure you really want to delete the app
                        </small>
                    </label>
                    <div class="input-group">
                        <input v-model="delete_app_verify" type="text" class="form-control" placeholder="Domain/AppName" aria-label="" aria-describedby="basic-addon1">
                        <div class="input-group-append">
                            <button class="btn btn-standard" style="width:65px;" @click="deleteApp(app.uuid)" type="button">😮</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';
    import jwt_decode from "jwt-decode";

    export default {
        name: 'General',
        components: {
        },
        data() {
            return {
                loggedInUser: null,
                verification_step: null,
                token_string: null,
                token_exp: null,
            };
        },
        props: {
            app_token:{
                type: Object,
                default: null,
            },
            app_uuid: {
                type: String,
                default: "",
            },
            app_locked: {
                type: Boolean,
                default: true,
            },
            app_owner: {
                type: String,
                default: "",
            },
        },
        created(){
            this.loggedInUser = jwt_decode(localStorage.getItem("token"))
            if (this.$props.app_token !== undefined || this.$props.app_token !== null) {
                this.token_string = this.$props.app_token?.jwt
                this.token_exp = this.$props.app_token?.expiration
            }
        },
        computed: {
            jwt(){
                return this.$props.app_token?.jwt
            },
            exp(){
                if (this.$props.app_token?.expiration === undefined || this.$props.app_token?.expiration === null) {
                    return this.token_exp
                }
                return this.$props.app_token?.expiration
            },
            expTimeSet() {
                const total = Math.abs(this.exp*1000 - new Date().getTime());
                const hours = Math.floor( (total/(1000*60*60)) % 24 );
                const days = Math.floor( total/(1000*60*60*24) );
                if (Number.isNaN(days)) {
                    return null
                }
                return {days: days, hours: hours};
            },
        },
        methods: {
            generateToken() {
                // this.$emit("loadApp", "world");
                if (this.verification_step === null) {
                    this.$moshaToast("Please provide the correct Organization/AppName", {type: 'warning',position: 'top-center', timeout: 3000})
                    return
                }
                const appOrgn = this.verification_step.split("/");
                if (appOrgn.length < 2) {
                    this.$moshaToast("Please provide the correct Organization/AppName", {type: 'warning',position: 'top-center', timeout: 3000})
                    return
                }
                let options = {
                    headers: {
                        'Authorization': localStorage.getItem("token"),
                    }
                };

                axios.post("http://192.168.0.177:8080/api/v1/app/token/issue", {
                    app_uuid: this.$props.app_uuid,
                    app_name: appOrgn[1],
                    orgn_domain: appOrgn[0],
                    // app_origin: this.$props.app?.app?.URL,
                }, options).then(res => {
                    this.token_string = res.data.app_token?.jwt;
                    this.token_exp = res.data.app_token?.expiration;
                    this.$emit("loadApp", this.$props.app_uuid);
                    this.$moshaToast(res.data.msg, {type: 'success',position: 'top-center', timeout: 3000})
                    this.$store.commit("UNSYNC_APP");
                  
                }).catch(err => {
                    this.$moshaToast(err.response?.data?.msg, {type: 'danger',position: 'top-center', timeout: 3000})
                });
                
            },
            unlockApp(id) {
                let options = {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': localStorage.getItem("token"),
                    }
                };
                axios.post("http://192.168.0.177:8080/api/v1/app/unlock", {
                        app_uuid: id,
                    }, options
                ).then(resp => {
                    // if (resp.status == 200) {
                    this.$emit("loadApp", id);
                    this.$moshaToast(resp?.data?.msg, {type: 'success',position: 'top-center', timeout: 3000})
                    // }
                }).catch(err => {
                    this.$moshaToast(err.response?.data?.msg, {type: 'danger',position: 'top-center', timeout: 3000})
                    return;
                });
            },
            deleteApp(id) {
                if (this.delete_app_verify == null) {
                    this.$toast.warning("Please provide the correct Organization/AppName");
                    return
                }
                const appOrgn = this.delete_app_verify.split("/");
                if (appOrgn.length < 2) {
                    this.$toast.warning("Please provide the correct Organization/AppName");
                    return
                }
                let options = {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': localStorage.getItem("token"),
                    }
                };
                axios.post("http://192.168.0.177:8080/api/v2/view/app/delete", {
                        app_uuid: id,
                        orgn_name: appOrgn[0],
                        app_name: appOrgn[1],
                    }, options
                ).then((resp) => {
                    if (resp.status == 200) {
                        this.$toast.success("App has been deleted");
                        this.$emit("drop_app", {"type": "drop_app", "app_uuid": id});
                    }
                }).catch(err => {
                    this.$toast.warning("Sorry app could not be removed");
                    return;
                });
            },
            copyTokenToClipboard() {
                navigator.clipboard.writeText(this.token).then(() => {
                    this.$moshaToast("App Token copied", {type: 'success',position: 'top-center', timeout: 3000})
                }).catch(() => {
                    this.$moshaToast("Failed to copy App Token", {type: 'danger',position: 'top-center', timeout: 3000})
                });
            },
        },
    };
</script>

<style scoped>
.view_component {
    margin-top: 15px;
    padding: 15px;
    border-radius: 8px;
}
h2 {
    margin: 5px 0px;
}

</style>
