<template>
    <h1 class="super-lg">{{orgn_domain}}/{{appName}}</h1>
    <div class="view_component">
        <div class="d-flex justify-end align-center mt-2">
            <button class="btn btn-standard" @click="createNewApp()">Create App</button>
        </div>
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
        <div class="form-row">
            <div class="form-group col">
                <div class="info-txt">Further configurations can be set in the app itself after its creation. There you will find
                    the options to configure a <strong>conversion rate funnel</strong>, <strong>campaign tracking</strong> and 
                    <strong>interesting buttons</strong>...👍
                </div>
            </div>
        </div>
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
        props: ["orgn_domain"],
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
                
                if (formValid) {
                    let options = {
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': localStorage.getItem("token"),
                        }
                    };
                    await axios.post("http://localhost:8080/api/v1/app/create", {
                            app_name: this.appName,
                            app_desc: this.appDesc,
                            app_url: this.appURL,
                        }, options
                    ).then((resp) => {
                        console.log("Resp ",resp);
                        this.$toast.success("App " + this.appName + " has been created");
                        this.$emit("createdApp", {"type": "show_app", "app_uuid": resp.data.uuid});
                    }).catch(err => {
                        console.log(err)
                        this.$toast.error(err.response.data);
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
    border-radius: 8px;
    height: max-content;
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

.checkboxes {
    columns: 3 8em;
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
