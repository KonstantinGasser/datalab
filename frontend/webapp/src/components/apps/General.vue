<template>
    <div>
        <h1>App Token</h1>
        <div class="view_component">
            <div class="form-row">
                <div class="form-group col">
                    <label for="">Verify Execution <br>
                        <small>(Your-Orgnaization-Domain/App-Name) 
                            <br>
                            This is required to ensure you actively want to create an App-Token
                        </small>
                    </label>
                    <div class="input-group">
                        <input v-model="token_placeholder" type="text" class="form-control" placeholder="Organisation-Domain/App-Name" aria-label="" aria-describedby="basic-addon1">
                        <div class="input-group-append">
                            <button class="btn btn-standard" @click="generateToken()" type="button">Authorize and Generate</button>
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
                        <textarea type="text" class="form-control" rows="2" readonly id="app_token_value" :value="token" aria-label="" aria-describedby="basic-addon1"></textarea>
                        <div class="input-group-append">
                            <button class="btn btn-standard" @click="copyTokenToClipboard()" type="button">Copy</button>
                        </div>
                    </div>
                    <div v-if="app_token || app.app_token" class=""><small>Token expires in {{get_valid_till.days}} days {{get_valid_till.hours}} hours</small></div>
                    <div class="mt-3">
                        Checkout the <a href="http://localhost:3000/docs/lib" target="_blank">documentation</a> 
                        on how to implement the client side
                    </div>
                </div>
            </div>
        </div>
        <br>
        <h1>~~Dangerous Water~~</h1>
        <div class="view_component">
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
                            <button class="btn btn-standard" @click="deleteApp(app.uuid)" type="button">delete 😮</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        name: 'General',
        components: {
        },
        data() {
            return {
                isEdit: false,
                token_placeholder: null,
                delete_app_verify: null,
                verified: false,
                app_token: null,
                new_img_url: null,
                header_name: "",
                valid_till: new Date().setDate(new Date().getDate() + 7),
                valid_days: null,
                valid_hours: null,
            };
        },
        props: ['app', 'token_placeholder'],
        computed: {
            token() {
                if (this.$props.app.app_token) {
                    return this.$props.app.app_token;
                } 
                if (this.app_token) {
                    return this.app_token;
                }
                return "";
                },
            get_valid_till() {
                const total = Math.abs(this.valid_hours - new Date()) / 1000;
                const hours = Math.floor( (total/(1000*60*60)) % 24 );
                const days = Math.floor( total/(1000*60*60*24) );
                return {days: days, hours: hours};
            },
        },
        methods: {
            generateToken() {
                if (this.$props.token_placeholder === null) {
                    this.$toast.warning("Please provide the correct Organization/AppName");
                    return
                }
                const appOrgn = this.$props.token_placeholder.split("/");
                if (appOrgn.length < 2) {
                    this.$toast.warning("Please provide the correct Organization/AppName");
                    return
                }
                let options = {
                    headers: {
                        'Authorization': localStorage.getItem("token"),
                    }
                };
                axios.post("http://localhost:8080/api/v2/view/app/generate/token", {
                    app_uuid: this.$props.app?.uuid,
                    app_name: appOrgn[1],
                    orgn_name: appOrgn[0],
                    orgn_domain: this.$props.app?.owner?.orgn_domain,
                    app_url: this.$props.app?.URL,
                }, options).then(res => {
                    console.log(res);
                    this.app_token = res.data.app_token;
                    this.$toast.success(res.data.msg);
                }).catch(err => {
                    this.$toast.warning(err.response.data.msg);
                });
            },
            deleteApp(id) {
                if (this.delete_app_verify == null) {
                    this.$toast.warning("Please provide the correct Organization/AppName");
                    return
                }
                const appOrgn = this.delete_app_verify.split("/");
                console.log("I: ", appOrgn);
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
                axios.post("http://localhost:8080/api/v2/view/app/delete", {
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
                    console.log(err);
                    this.$toast.warning("Sorry app could not be removed");
                    return;
                });
            },
            copyTokenToClipboard() {
                navigator.clipboard.writeText(this.token).then(res => {
                    
                    this.$toast.success("Token copied to clipboard");
                }).catch(err => {
                    this.$toast.error("Failed to copy token...");
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
