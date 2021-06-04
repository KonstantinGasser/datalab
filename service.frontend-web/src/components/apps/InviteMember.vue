<template>
    <div>
        <br>
        <div class="">
            <h1>Invite Colleagues</h1>
            <br>
            <div class="notify-table">
                <div v-for="item in colleagues" :key="item.uuid" class="notify-row">
                    <div class="d-flex align-center">
                        <div>
                            <div class="avatar medium">
                                <img :src="item.Avatar" alt="">
                            </div>
                        </div>
                        <div class="ml-2">
                            <h2><strong>{{item.first_name}} {{item.last_name}}</strong></h2>
                            <h4>@{{item.username}}</h4>
                        </div>
                    </div>
                    <div class="actions">            
                        <div class="px-1"> 
                            <button v-if="item.status === 0 && loggedInUser.sub === app_owner?.uuid" class="btn accept" @click="invite(item)">invite</button>
                            <button v-if="item.status === 0 && loggedInUser.sub !== app_owner?.uuid" class="btn accept" disabled>invite</button>
                            <div  v-if="item.status === 1" class="invited pending">Pending 
                                <span v-if="loggedInUser.sub === app_owner?.uuid" @click="reminder(item)" class="ml-1 icon icon-external-link hover"></span>
                            </div>
                            <div  v-if="item.status === 2 && item.uuid !== app_owner?.uuid" class="invited accepted">Accepted</div>
                            <div  v-if="item.uuid === app_owner?.uuid" class="invited owner">Owner</div>
                            <div  v-if="item.status === 3" class="invited rejected">Rejected</div>
                        </div>
                    </div> 
                </div>
            </div>
        </div>
    </div>
</template>

<script>


import axios from "axios";
import jwt_decode from "jwt-decode";

export default {
    name: "InviteMember",
    data() {
        return {
            loggedInUser: null,
            colleagues: [],
            inTeam: [],
        }
    },
    props: {
        app_name: {
            type: String,
            default: null,
        },
        app_owner: {
            type: Object,
            default: "",
        },
        app_uuid: {
            type: String,
            default: null,
        },
        member: {
            type: Array,
            default: null,
        }
    },
    computed: {
    },
    async created() {
        this.loggedInUser = jwt_decode(localStorage.getItem("token"))
        let options = {
            headers: {
                'Authorization': localStorage.getItem("token"),
            }
        };
        const payload = {
            app_uuid: this.$props.app_uuid
        }
        const resp = await axios.post("http://192.168.0.177:8080/api/v1/app/member/invitable", payload, options);
        if (resp.status != 200) {
            this.$moshaToast(resp.data, {type: 'danger',position: 'top-center', timeout: 3000})
            return;
        }
        this.colleagues = resp.data.user;
    },
    methods: {
        updateRole(item, role) {
            item["role"] = role
        },
        async invite(user) {
            let options = {
                headers: {
                    'Authorization': localStorage.getItem("token"),
                }
            };
            const payload = {
                app_uuid: this.$props.app_uuid,
                invited_uuid: user.uuid,
            }
            const resp = await axios.post("http://192.168.0.177:8080/api/v1/app/invite", payload, options);
            if (resp.status === 200) {
                this.$moshaToast(resp.data.msg, {type: 'success',position: 'top-center', timeout: 3000})
                this.inTeam.push(user.uuid)
                user["status"] = 1
                return
            }
            this.$moshaToast(resp.response.data.msg, {type: 'danger',position: 'top-center', timeout: 3000})   
        },
        async reminder(user) {
            let options = {
                headers: {
                    'Authorization': localStorage.getItem("token"),
                }
            };
            const payload = {
                app_name: this.$props.app_name,
                app_uuid: this.$props.app_uuid,
                user_uuid: user.uuid,
            }
            try {
                const resp = await axios.post("http://192.168.0.177:8080/api/v1/app/invite/reminder", payload, options);
            if (resp.status === 200) {
                this.$moshaToast(resp.data.msg, {type: 'success',position: 'top-center', timeout: 3000})
                return
            }
            } catch(err) {
                this.$moshaToast(err.response?.data?.msg, {type: 'danger',position: 'top-center', timeout: 3000})   
            }
            
            
        }
    },
}
</script>

<style scoped>
h4 {
    color: var(--btn-bg-hover);
}
.t-row {
    padding: 10px 15px;
    width: 400px;
    background: var(--sub-border);
    border-radius: 4px;
    color: #000;
}
.username {
    font-weight: 400;
    font-size: 16px;
}

.notify-table {
    display: grid;
    row-gap: 15px;
    width: 100%;
    padding: 15px 25px;
}

.notify-row {
    width: 100%;
    padding: 10px 20px;
    border-radius: 8px;
    color: var(--h-color);
    background: var(--sub-border);
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.emoji-line {
    font-size: 18px;
    margin-bottom: 10px;
}

.actions {
    display: flex;
    align-items: center;
}

.actions .btn {
    color: var(--h-color);
    border-radius: 8px;
    width: 100px;
}

.btn.accept {
    background: #10d57425;
    border: 1px solid #10d574;
}
.btn.accept:hover {
    background: #10d57475;
    border: 1px solid #10d574;
}
.btn.accept:disabled:hover {
    background: #10d57425;
    border: 1px solid #10d574;
}

.invited {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100px;
    height: 35px;
    border-radius: 50px;
    /* color: #00000075; */
}
.pending {
    background: #f7fd0450;
    border: 1px solid #f7fd04;
    color: var(--font-yellow);
}
.accepted {
    background: #10d57450;
    border: 1px solid #10d574;
    color: var(--font-green);
}

.rejected {
    background: #d9042975;
    border: 1px solid #d90429;
}

.owner {
    background: #5465ff54;
    border: 1px solid #5465ff; 
    color: var(--font-blue);
}

</style>