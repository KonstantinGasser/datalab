<template>
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
                        <button class="btn btn-standard" @click="verifyStep()" type="button">Authorize and Generate</button>
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
                    <textarea type="text" class="form-control" rows="2" readonly id="app_token_value" :value="app_token" aria-label="" aria-describedby="basic-addon1"></textarea>
                    <div class="input-group-append">
                        <button class="btn btn-standard" @click="copyTokenToClipboard()" type="button">Copy</button>
                    </div>
                </div>
                <!-- <input type="text" disabled name="" class="form-control" id="app_token_value" :value="app_token">  -->
            </div>
        </div>
    </div>
    <div class="view_component">
        <div class="form-row">
            <div class="form-group">
                <label for="">Implementing the Client-Library</label>
                <div class="read-only-text">Lorem ipsum, dolor sit amet consectetur adipisicing elit. Debitis accusantium iure laborum fuga exercitationem nostrum iste, ducimus molestias! Iure corporis cum cumque, pariatur reprehenderit incidunt aliquid aliquam officia asperiores consequuntur?</div>
            </div>
        </div>
    </div>

</template>

<script>

    export default {
        name: 'TabAppToken',
        components: {},
        data() {
            return {
                isEdit: false,
                verify_app_name: null,
                verified: false,
                app_token: "This is a cool application token"
            };
        },
        props: ['app'],
        methods: {
            setMode() {
                this.isEdit = !this.isEdit;
                // emit panel is in edit mode
                // diable tabs until saved
                this.$emit('inEdit', this.isEdit);
            },
            verifyStep() {
                if (this.verify_app_name == null || this.verify_app_name !== "datalabs.dev/test") {
                    this.$toast.warning("Please provide the correct domain/name to proceed")
                    return
                }
                this.verified = true;
            },
            copyTokenToClipboard() {
                  var token = document.getElementById("app_token_value");
                token.select();
                token.setSelectionRange(0, 99999)
                document.execCommand("copy");
                alert("Copied the text: " + token.value);
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
