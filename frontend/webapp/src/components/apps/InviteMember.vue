<template>
    <div>
        <!-- <div class="view_component">
            already invited {{colleagues}}
        </div> -->
        <div class="">
            <h2>Invite Colleagues</h2>
            <br>
            <div class="notify-table">
                <div v-for="item in canInvite" :key="item.uuid" class="notify-row">
                    <div>
                        <div class="emoji-line d-flex justify-start">👉<strong>&nbsp;{{item.first_name}} {{item.last_name}} (@{{item.username}})</strong></div>
                        <div  class="notify-title">
                            Position <strong>{{item.orgn_position}}</strong>
                        </div>
                    </div>
                    <div class="actions">            
                        <div class="px-1">
                            <button v-if="item.status === undefined" class="btn accept" @click="invite(item)">invite</button>
                            <div  v-if="item.status === 1" class="invited pending">Pending</div>
                            <div  v-if="item.status === 2" class="invited accepted">Accepted</div>
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

export default {
    name: "InviteMember",
    data() {
        return {
            colleagues: [],
            inTeam: [],
        }
    },
    props: {
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
        canInvite(){
           if (this.colleagues !== null && this.$props.member !== null) {
                this.colleagues.forEach(col => {
                    for (let i = 0; i < this.$props.member.length; i++) {                       
                        if (col.uuid == this.$props.member[i].uuid) {
                            col["status"] = this.$props.member[i].status
                            console.log("Found: ", col)
                        }
                    }
                })
                console.log(this.colleagues)
            }
            return this.colleagues
        },
    },
    async created() {
        let options = {
            headers: {
                'Authorization': localStorage.getItem("token"),
            }
        };
        const resp = await axios.get("http://192.168.0.177:8080/api/v1/user/profile/colleagues", options);
        if (resp.status != 200) {
            this.$toast.error("Could not fetch Your Colleagues");
            return;
        }
        this.colleagues = resp.data.colleagues;
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
                app_role: 0,
            }
            const resp = await axios.post("http://192.168.0.177:8080/api/v1/app/member/invite", payload, options);
            if (resp.status === 200) {
                this.inTeam.push(user.uuid)
                user["status"] = 1
            }
            console.log(resp)
            
        }
    },
}
</script>

<style scoped>
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

.invited {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100px;
    height: 35px;
    border-radius: 50px;
}
.pending {
    background: #f7fd0450;
    border: 1px solid #f7fd04;
}
.accepted {
    background: #10d57450;
    border: 1px solid #10d574;
}

.rejected {
    background: #d9042975;
    border: 1px solid #d90429;
}


</style>