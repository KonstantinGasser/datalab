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
      <!-- <h2>App Details</h2>
            <div class="view_component">
                <div class="form-row">
                    <div class="form-group col">
                        <div class="form-row mb-2">
                            <label for="">App Name</label>
                            <input type="text" class="form-control" readonly name="" id="" :placeholder="app.name">
                        </div>
                    </div>
                    <div class="form-group col">
                        <label for="">App URL</label>
                        <input v-model="appURL" type="text" class="form-control" name="" readonly id="app_url" :placeholder="app.URL">
                    </div>
                </div>
                <div class="form-row">
                    <div class="form-group col">
                        <label for="">App Description</label>
                        <textarea readonly class="form-control" name="" id="app_desc" rows="2" :placeholder="app.description"></textarea>
                    </div>
                    <div class="form-group col">
                        <label for="">App Settings</label>
                        <ul class="tag-list">
                            <li class="tag-standard" v-for="item in app.settings" :key="item">{{item}}</li>
                        </ul>
                    </div>
                </div>
            </div> -->
            <!-- <hr> -->
            <!-- <br>
            <h1>Images for Mouse-Movements</h1>
            <div class="view_component">
                    <div class="img-upload d-flex justify-between">
                        <div class="d-flex justift-between w-100">
                            <div class="input-group mr-3 w-50">
                                <input v-model="new_img_url" type="text" class="form-control" placeholder="URL for Image">
                            </div>
                            <div class="d-flex align-center">
                                <label for="file-upload" class="custom-file-upload">Select File  📄</label>
                                <input id="file-upload" type="file" accept="image/png, image/jpeg" enctype="multipart/form-data"/>
                            </div>
                        </div>
                        <div class="">
                            <button type="submit" class="btn btn-standard" @click="onUpload">Update images</button>
                        </div>
                    </div>
                <hr>
                <div class="img-grid">
                    <div class="img-card">
                        <img src="https://lh3.googleusercontent.com/proxy/l7r-8QDwkB07BJwYZ7tivKsKDqGwvhTcOriEyi3Cxh8gaW18_vU8OYpgBsKlIbSAqzZr5Ji2ZXtFopbUnDZ-zg6C4kDFOLe58fzQqMaAKhflp7r2-8v9NjHYjbeUA7n0To8Zo1E_YShTSux2wJCAkPfgfC7Mjw8CGzpuWpU" alt="">
                        <input type="text" class="form-control" placeholder="app/products/iphones"/>
                    </div>
                </div>
            </div> -->
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
            };
        },
        props: ['app', 'token_placeholder'],
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
                    app_uuid: this.$props.app.uuid,
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
    height: max-width;
}
h2 {
    margin: 5px 0px;
}

</style>
