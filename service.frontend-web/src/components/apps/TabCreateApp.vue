<template>
    <h1 class="super-lg">{{orgn_domain}}/{{appName}}</h1>
    <div class="view_component">
        <h2>General</h2>
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
                <label for="">App Tags</label>
                <div class="tag-input">
                    <input class="form-control" v-model="newTag" type="text" :disabled="tags?.length >= 5"
                        placeholder="#awesome"
                        @keydown.enter="addTag(newTag)" 
                        @keydown.prevent.tab="addTag(newTag)"/>
                    <ul class="tags">
                        <li v-for="(tag, index) in tags" :key="tag" class="tag">
                        {{ tag }}
                        <button class="delete" @click="removeTag(index)">x</button>
                        </li>
                    </ul>
                </div>
               <!-- <textarea v-model="appDesc" class="form-control" name="" id="app_desc" rows="2" placeholder="what is the app about?"></textarea> -->
            </div>
        </div>
    </div>
    <div class="view_component">
        <h2>Visibility <small>(default is public for Organization)</small></h2>
        <div class="form-row">
            <div class="form-group col">
                <div class="d-flex justify-center align-center">
                    <div class="custom-control custom-switch">
                        <input v-model="private_app" type="checkbox" class="custom-control-input" id="customSwitch1">
                        <label class="custom-control-label" for="customSwitch1"></label>
                    </div>
                    <span class="ml-1 icon icon-lock" style="font-size:20px"></span>
                </div>
                <hr>
                <div>
                    <p>
                        If you set the <strong>Visibility</strong> to public - any one from your organization can see and contribute to the <strong>App</strong>.
                        The obligation to <strong>create an App-Token</strong>, <strong>delete the App</strong> and to <strong>unlock the App</strong> remains only for the
                        <strong>App-Owner</strong>
                    </p>
                </div>
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
    import {ref} from 'vue';

    export default {
        name: 'TabCreateApp',
        components: {},
        data() {
            return {
                tags: [],
                isEdit: false,
                appName: null,
                appURL: null,
                appDesc: null,
                private_app: false,
                appMember: [],
                appCfgs: [],
            };
        },
         setup(){
            const addTag = (tag) => {
                tags.value.push(tag); // add the new tag to the tags array
                newTag.value = ""; // reset newTag
            };
            
            const tags = ref([]);
            const newTag = ref('') //keep up with new tag
            return { tags, newTag, addTag }
        },
        props: ["orgn_domain"],
        methods: {
            removeTag(index) {
                console.log(this.tags)
                this.tags.splice(index, 1);
            },
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
                    this.$moshaToast('App URL is required', {type: 'danger',position: 'top-center', timeout: 3000})
                    formValid = false;
                }
                if (!this.validURL(this.appURL)) {
                    this.$moshaToast('Mhm does not look like an URL does it ?', {type: 'danger',position: 'top-center', timeout: 3000})
                    formValid = false;
                }
                if (this.appName === null || this.appName === '') {
                    this.$moshaToast('App Name is required', {type: 'danger',position: 'top-center', timeout: 3000})
                    formValid = false;
                }
                if (this.tags.length <= 0) {
                    this.$moshaToast('App Tags are required', {type: 'danger',position: 'top-center', timeout: 3000})
                    formValid = false;
                }
                
                if (formValid) {
                    let options = {
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': localStorage.getItem("token"),
                        }
                    };
                    await axios.post("http://192.168.0.177:8080/api/v1/app/create", {
                            app_name: this.appName,
                            tags: this.tags,
                            app_url: this.appURL,
                            is_private: this.private_app,
                        }, options
                    ).then((resp) => {
                        this.$moshaToast("App " + this.appName + " has been created", {type: 'success',position: 'top-center', timeout: 3000})
                        this.$emit("createdApp", {"type": "show_app", "app_uuid": resp.data.uuid});
                    }).catch(err => {
                        this.$moshaToast(err.response?.data?.msg, {type: 'danger',position: 'top-center', timeout: 3000})
                        return;
                    });
                }
            },
            validURL(str) {
                const re = new RegExp("http://(localhost|127.0.0.1):[0-9]+")
                const matches = re.exec(str)
                if (matches && matches?.length > 0) {
                    return true
                }
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


ul {
  list-style: none;
  display: flex;
  align-items: center;
  gap: 7px;
  margin: 0;
  padding: 0;
}
/* .tag-input > input {
  width: 100%;
  padding: 10px;
} */
.tag {
    margin-top: 5px;
    border: 1px solid #5465ff;
    background: #5465ff54;
    padding: 2px 10px;
    border-radius: 15px;
    color: #5465ff;
    white-space: nowrap;
    transition: 0.1s ease background;
}
.delete {
  color: white;
  background: none;
  outline: none;
  border: none;
  cursor: pointer;
}
</style>
