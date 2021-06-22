<template>
    <div class="view_component">
        <div class="d-flex flex-wrap">
            <div class="app-card">
                <div class="d-flex justify-end">
                    <span class="icon icon-external-link hover" @click="open('aa6fa68c-c938-4910-9247-4543ac6d5acf')"></span>
                </div>
                <div class="d-grid h-100">
                    <div>
                        <h1>AwesomeApp</h1>
                        <h4>testing avatars in app</h4>
                        <hr>
                    </div>
                    <div class="app-info d-flex align-center justify-between">
                        <div class="app-owner d-flex align-end col-gap-5">
                            <div class="avatar small">
                                    <img src="" alt="">
                            </div>
                            <div class="d-grid">
                                <small>Konstantin Gasser</small>
                                <small>@kgasser</small>
                            </div>
                        </div>
                        <div class="app-stats">
                            <div><span class="icon icon-users"></span>5</div>
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
    name: "ViewOverview",
    data() {
        return {
            loggedInUser: null,
        }
    },
    props: {
    },
    computed: {
    },
    // async created() {
    //     this.loggedInUser = jwt_decode(localStorage.getItem("token"))
    //     let options = {
    //         headers: {
    //             'Authorization': localStorage.getItem("token"),
    //         }
    //     };
    //     const payload = {
    //         app_uuid: this.$props.app_uuid
    //     }
    //     const resp = await axios.post("http://localhost:8080/api/v1/app/member/invitable", payload, options);
    //     if (resp.status != 200) {
    //         this.$moshaToast(resp.data, {type: 'danger',position: 'top-center', timeout: 3000})
    //         return;
    //     }
    //     this.colleagues = resp.data.user;
    // },
    methods: {
       async open(uuid) {
        this.$emit("openApp", uuid)
       }
    },
}
</script>

<style scoped>
hr {
    border-top: 1px solid rgba(0,0,0,.1);
}
h4 {
    color: var(--h-color);
}
.view_component {
    background: none;
    margin-top: 15px;
    padding: 15px;
    border: none;
    border-radius: 8px;
    height: 100%;
}

.app-card {
    background: var(--sub-bg);
    border-radius: 8px;
    width: 300px;
    height: 150px;
    padding: 15px 20px;
    border: 1px solid var(--sub-border);
}

.app-stats {
    font-size: 18px;
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