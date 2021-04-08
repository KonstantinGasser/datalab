<template>
    <div class="scroll_view">
        <h2>App Details</h2>
        <div class="view_component">
            <div class="form-row">
                <div class="form-group col">
                    <div class="form-row mb-2">
                        <label for="">App Name</label>
                        <input type="text" class="form-control" readonly name="" id="" :placeholder="app.app_name">
                    </div>
                    <div class="form-row">
                        <label for="">App Description</label>
                        <textarea type="text" class="form-control" rows="2" readonly id="app_token_value" :placeholder="app.app_description"></textarea>
                    </div>
                    <!-- <label for="">App Name</label>
                    <input type="text" class="form-control" readonly name="" id="" :placeholder="app.app_name"> -->
                </div>
                <!-- <div class="form-group col">
                    <label for="">App Description</label>
                    <input type="text" class="form-control" readonly name="" id="" :placeholder="app.app_description">
                </div> -->
                <div class="form-group col">
                    <label for="">App Settings</label>
                    <ul class="tag-list">
                        <li class="tag-standard" v-for="item in app.app_setting" :key="item">{{item}}</li>
                    </ul>
                </div>
            </div>
        </div>
        <hr>
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
                        <input v-model="verify_app_name" type="text" class="form-control" placeholder="Domain/AppName" aria-label="" aria-describedby="basic-addon1">
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
                    <div class="mt-3">
                        Checkout the <a href="http://localhost:3000/docs/lib" target="_blank">documentation</a> 
                        on how to implement the client side
                    </div>
                </div>
            </div>
        </div>
        <!-- <hr>
        <h1>Implementing the Client-Library</h1>
        <div class="view_component">
            <div class="form-row">
                <div class="form-group">
                    <label for="">Implementing the Client-Library</label>
                    <vue3-markdown-it :source='source' />
                </div>
            </div>
        </div> -->
        <hr>
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
                        <input v-model="verify_app_name" type="text" class="form-control" placeholder="Domain/AppName" aria-label="" aria-describedby="basic-addon1">
                        <div class="input-group-append">
                            <button class="btn btn-standard" @click="verifyStep()" type="button">delete 😮</button>
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
        name: 'TabAppDetails',
        components: {},
        data() {
            return {
                isEdit: false,
                verify_app_name: null,
                verified: false,
                app_token: "This is a cool application token",
                source: "```func Read(b []byte) error```",
            };
        },
        props: ['app'],
        computed: {
            token() {
                console.log(this.$props.app.app_token);
                if (this.$props.app.app_token) {
                    return this.$props.app.app_token;
                } 
                if (this.app_token) {
                    return this.app_token;
                }
                return "";
                // return this.$props.app.app_token ? this.$props.app.app_token != null : this.app_token;
                },
        },
        methods: {
            setMode() {
                this.isEdit = !this.isEdit;
                // emit panel is in edit mode
                // diable tabs until saved
                this.$emit('inEdit', this.isEdit);
            },
            generateToken() {
                if (this.verify_app_name == null) {
                    this.$toast.warning("Please provide the correct Organization/AppName");
                    return
                }
                const appOrgn = this.verify_app_name.split("/");
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
                    app_uuid: this.$props.app.app_uuid,
                    app_name: appOrgn[1],
                    orgn_name: appOrgn[0],
                }, options).then(res => {
                    console.log(res);
                    this.app_token = res.data.app_token;
                    this.$toast.success(res.data.msg);
                }).catch(err => {
                    if (err.response.status === 403) {
                        this.$toast.error("Organization/AppName do not match")
                    }
                    console.log(err);
                });
            },
            copyTokenToClipboard() {
                var token = document.getElementById("app_token_value");
                token.select();
                token.setSelectionRange(0, 99999)
                document.execCommand("copy");
                console.log(token);
                // alert("Copied the text: " + token.value);
            },
        },
    };
</script>

<style scoped>
.view_component {
    margin-top: 15px;
    padding: 15px;
    background: #1E1E1E;
    border-radius: 8px;
    height: max-width;
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
</style>
