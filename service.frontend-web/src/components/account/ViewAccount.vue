<template>
    <div class="view_app">
        <div></div>
        <div>
            <h2>Personal Information</h2>
            <div class="view_component">
                <div class="d-flex justify-end align-center mt-2">
                    <button class="btn btn-standard" @click="updateAccount()">Update</button>
                </div>
                <div class="form-row justfy-start">
                    <div class="form-group col d-grid justify-center align-center">
                        <div class="d-flex align-center justify-center">
                            <div class="avatar">
                                <img class="circle-img big" :src="user.avatar" alt="">
                            </div>
                        </div>
                        <div class="flex-wrap d-flex align-center justify-center">
                            <span class="text-capture big">@{{user.username}}</span>
                        </div>
                    </div>
                    <div class="form-group col">
                        <div class="form-row" style="margin-bottom: 32px;">
                            <label for="">First Name</label>
                            <input v-model="user.first_name" class="form-control" type="text" name="" id="first_name" :placeholder="user.first_name">
                        </div>
                        <div class="form-row">
                            <label for="">Last Name</label>
                            <input v-model="user.last_name" class="form-control" type="text" name="" id="last_name" :placeholder="user.last_name">
                        </div>
                        
                    </div>
                    <!-- <div class="form-group col">
                        <label for="">Last Name</label>
                        <input v-model="user.last_name" class="form-control" type="text" name="" id="last_name" :placeholder="user.last_name">
                    </div> -->
                </div>
            </div>
            <br>
            <h2>Company Information</h2>
            <div class="view_component">
                <div class="form-row">
                    <div class="form-group col">
                        <label for="">Organization Domain</label>
                        <input class="form-control" type="text" name="" id="orgn_domain" readonly :placeholder="user.orgn_domain">
                    </div>
                    <div class="form-group col">
                        <label for="">Organization Position</label>
                        <input v-model="user.orgn_position" class="form-control" type="text" name="" id="last_name" :placeholder="user.orgn_position">
                    </div>
                </div>
            </div>
            <br>
            <h2>Colleagues</h2>
            <br>
            <div class="d-flex flex-wrap">
                <div v-for="item in colleagues" :key="item.uuid">
                    <div class="colleague d-grid align-center py-2 px-2">
                        <div class="d-flex align-center justify-center">
                            <div class="avatar medium" :class="{'online': isOnline(item.uuid)}">
                                <img :src="item.avatar" alt="">
                            </div>
                        </div>
                        <div class="d-grid align-center justify-center">
                            <div class="d-flex align-center justify-center"><h2><strong>{{item.first_name}} {{item.last_name}}</strong></h2></div>
                            <div class="d-flex align-center justify-center"><h4>@{{item.username}}</h4></div>
                        </div>
                        <div class="d-flex align-center justify-center">
                            {{item.orgn_position}}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';
    import jwt_decode from 'jwt-decode';

    export default {
        name: 'ViewAccount',
        components: {
        },
        data() {
            return {
                loggedInUser: null,
                colleagues: [],
                user: {
                    first_name: null,
                    last_name: null,
                    orgn_domain: null,
                    orgn_position: null,
                    username: null,
                },
            };
        },
        created() {
            this.loggedInUser = jwt_decode(localStorage.getItem("token"));

            let options = {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': localStorage.getItem("token"),
                }
            };
            
            axios.get("http://192.168.0.177:8080/api/v1/user/colleagues", options).then(resp => {
                const tmp = resp?.data?.user;
                this.colleagues = tmp.filter(item => item.uuid != this.loggedInUser.sub);
            })
            
        },
        mounted() {
            this.user = this.fetchUpdate();
        },
        methods: {
            isOnline(uuid) {
                if (this.$store.state.online === undefined) {
                    return false
                }
                for (let i = 0; i < this.$store.state.online?.length; i++) {
                    if (this.$store.state.online[i] === uuid) {
                        return true
                    }
                }
                return false
            },
            async updateAccount() {
                let options = {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': localStorage.getItem("token"),
                    }
                };
                
                const resp = await axios.post("http://192.168.0.177:8080/api/v1/user/profile/update", {
                    firstname: this.user.first_name,
                    lastname: this.user.last_name,
                    position: this.user.orgn_position,
                }, options)
                if (resp.status !== 200) {
                    this.$moshaToast(resp?.data?.msg, {type: 'danger',position: 'top-center', timeout: 3000})
                    return
                }
                this.$moshaToast(resp?.data?.msg, {type: 'success',position: 'top-center', timeout: 3000})
                // fetch updated user
                this.fetchUpdate();
            },
            fetchUpdate() {
                let options = {
                    headers: {
                        'Authorization': localStorage.getItem("token"),
                    }
                };

                axios.get("http://192.168.0.177:8080/api/v1/user/profile", options).then(resp => {
                    this.user = resp.data.user;
                }).catch(err => {
                    if (err.response.status === 401) {
                        localStorage.removeItem('token');
                        this.$router.replace({ name: 'login' });
                    }
                });
            },
        },
    };
</script>

<style scoped>
h4 {
    color: var(--btn-bg-hover);
}
.online {
    border: 2px solid var(--btn-bg-hover) !important;
}
.colleague {
    width: 200px;
}

.tab-line {
    grid-column: 1;
    grid-row: 1;
}

.view_component {
    margin-bottom: 15px;
    padding: 15px;
    border-radius: 8px;
    height: max-content;
}
h2 {
    margin: 5px 0px;
}

.view_app {
    height: 100%;
}

.app_list {
    background: #0D1116;
    border-radius: 8px;
    padding: 15px;
    height: max-content;
    min-height: 225px;
    max-height: 100%;
    overflow-y: scroll;
    border: 1px solid #30363D;
}

.app_name_list {
    margin-top: 25px;
}
.app-name {
    background:linear-gradient(135deg, #50e3c2 0%,#10d574 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent; 
    font-size: 16px;
    font-weight: bold;
    padding: 5px;
    margin: 5px 0;
    border: 1px solid #10d574;
    border-radius: 8px;
    border-style: dashed;
}
.app-name:hover {
    cursor: pointer;
    text-decoration: underline;
}

</style>
