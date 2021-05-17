<template>
    <div>
        <!-- <div class="view_component">
            already invited {{colleagues}}
        </div> -->
        <div class="view_component">
            <h2>Invite Colleagues</h2>
            <br>
            <div v-for="item in colleagues" :key="item.uuid" class="t-row d-flex align-center justify-between">
                <div class="username">@{{item.username}}</div>
                <div>{{item.orgn_position}}</div>
                <div>
                    <select class="custom-select">
                        <option selected>App Role</option>
                        <option value="1">Editor</option>
                        <option value="2">Viewer</option>
                    </select>
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
            colleagues: null,
        }
    },
    props: {
        app_uuid: {
            type: String,
            default: null,
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
    font-weight: bolder;
}
</style>