<template>
   <div>
        <h1>Any News? 🧐</h1>
        <div class="">
            <div class="no-notifications" v-if="notifications === null || notifications.length === 0">
                Looks like you are up-to date. We let you know when you have any news 😄🙌
            </div>
            <div class="notify-table">
                <div v-for="item in notifications" :key="item" class="notify-row">
                    <!-- <div>
                        <span @click="hideNotify(item)">hide</span>
                    </div> -->
                    <div class="notify-item" v-if="item.event === 0">
                        <div>
                            <div class="emoji-line d-flex justify-start"><strong>You got an App Invite</strong>&nbsp;- go check it out 🚀</div>
                            <div  class="notify-title">
                                👉 &nbsp;
                                <strong>{{item.value?.app_owner}}</strong> invited you to join and contribute to&nbsp;
                                <strong>"{{item.value?.app_name}}"</strong> <br>
                            </div>
                        </div>
                        <div class="actions">
                            <div class="d-flex justify-between col-gap-15">
                                <div class="">
                                    <button class="btn accept" @click="acceptInvite(item, item.value?.app_uuid)">Accept</button>
                                </div>

                                <div class="">
                                    <button class="btn reject" @click="rejectInvite(item, item.value?.app_uuid)">Reject</button>
                                </div>
                            </div>
                            <div class="d-flex align-center">
                                <span @click="hideNotify(item)" class="hover">❌</span>
                            </div>
                        </div>
                    </div>
                    <div class="notify-item" v-if="item.event === 1">
                        <div>
                            <div class="emoji-line d-flex justify-start"><strong>Everyone should get a second change</strong>&nbsp;🤗</div>
                            <div  class="notify-title">
                                ❗️&nbsp;
                                <strong>{{item.value?.app_owner}}</strong> reminds you to join&nbsp;
                                <strong>"{{item.value?.app_name}}"</strong> <br>
                            </div>
                        </div>
                        <div class="actions">
                            <div class="d-flex justify-between col-gap-15">
                                <div class="mr-1">
                                    <button class="btn accept" @click="acceptInvite(item, item.value?.app_uuid)">Accept</button>
                                </div>

                                <div class="ml-1">
                                    <button class="btn reject" @click="rejectInvite(item, item.value?.app_uuid)">Reject</button>
                                </div>
                            </div>
                            <div class="d-flex align-center">
                                <span @click="hideNotify(item)" class="hover">❌</span>
                            </div>
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
  name: 'NotificationCenter',
  data() {
    return {
        loggedInUser: null,
        notifies: [],
    };
  },
  components: {
  },
  computed: {
      notifications() {
          return this.$store.state.notifications;
      },
  },
  created() {
      this.loggedInUser = jwt_decode(localStorage.getItem("token"));
  },
  methods: {
      async acceptInvite(item, app_uuid) {
          let options = {
                headers: {
                    'Authorization': localStorage.getItem("token"),
                }
            };
            const payload = {
                app_uuid: app_uuid,
                event_timestamp: item.timestamp,
            }
            const resp = await axios.post("http://localhost:8080/api/v1/app/invite/accept", payload, options);
            if (resp.status != 200) {
                this.$moshaToast(resp.data.msg, {type: 'danger',position: 'top-center', timeout: 3000})
                return
            }
        //     localStorage.setItem("token", resp.data.token)
            this.$moshaToast(resp.data?.msg, {type: 'success',position: 'top-center', timeout: 3000})
            this.popNotification(item)
      },
      async hideNotify(item) {
          let options = {
                headers: {
                    'Authorization': localStorage.getItem("token"),
                }
            };
            const payload = {
                user_uuid: this.loggedInUser.sub,
                timestamp: item.timestamp,
            }
            const resp = await axios.post("http://localhost:8008/api/v1/datalab/hide/event", payload, options);
            if (resp.status != 200) {
                this.$moshaToast(resp.data?.msg, {type: 'danger',position: 'top-center', timeout: 3000})
                return
            }
            this.$moshaToast(resp.data?.msg, {type: 'success',position: 'top-center', timeout: 3000})
            this.popNotification(item)
      },
      rejectInvite(app_uuid) {
      },
      popNotification(item) {
          this.$store.commit("POP_NOTIFICATION", item)
      }
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.hover {
    cursor: pointer;
}
.col-gap-15 {
    column-gap: 15px;
}
.app-view {
    height: 100%;
}

.no-notifications {
    margin-top: 75px;
    background: #5465ff54;
    border: 1px solid #5465ff;
    color: #fff;
    padding: 25px;
    text-align: center;
    font-size: 20px;
    border-radius: 8px;
}

.notify-table {
    display: grid;
    justify-items: flex-start;
    row-gap: 15px;
    width: 100%;
    padding: 15px 25px;
}

.notify-row {
    width: 100%;
}
.notify-item {
    width: 100%;
    height: max-content;
    max-width: 100%;
    overflow-wrap: break-word;
    column-gap: 25px;
    padding: 5px 25px;
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
    column-gap: 15px;
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
.btn.reject {
    background: #d9042925;
    border: 1px solid #d90429;
}

.btn.accept:hover {
    background: #10d57475;
    border: 1px solid #10d574;
}
.btn.reject:hover {
    background: #d9042975;
    border: 1px solid #d90429;
}
</style>
