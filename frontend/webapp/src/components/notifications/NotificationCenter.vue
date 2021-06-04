<template>
   <div>
        <h1>Any News? 🧐</h1>
        <div class="">
            <!-- Hello NotificationCenter
            {{notifications}} -->
            <div class="notify-table">
                <div v-for="item in notifications" :key="item" class="notify-row">
                    {{item}}
                    <div v-if="item.value?.event === 0">
                        <div class="emoji-line d-flex justify-start"><strong>You got an App Invite</strong>&nbsp;- go check it out 🚀</div>
                        <div  class="notify-title">
                            👉 &nbsp;
                            <strong>{{item.value?.app_owner}}</strong> invited you to join and contribute to&nbsp;
                            <strong>"{{item.value?.app_name}}"</strong> <br>
                        </div>
                    </div>
                    <div class="actions">
                        <div class="py-1">
                            <button class="btn accept" @click="acceptInvite(item, item.value?.app_uuid)">Accept</button>
                        </div>

                        <div class="py-1">
                            <button class="btn reject" @click="rejectInvite(item, item.value?.app_uuid)">Reject</button>
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
  name: 'NotificationCenter',
  data() {
    return {
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
  methods: {
      async acceptInvite(item, app_uuid) {
          console.log(item)
          console.log(app_uuid)
          let options = {
                headers: {
                    'Authorization': localStorage.getItem("token"),
                }
            };
            const payload = {
                app_uuid: app_uuid,
            }
            const resp = await axios.post("http://192.168.178.103:8080/api/v1/app/member/invite/accept", payload, options);
            if (resp.status != 200) {
                this.$toast.error("Could not send invite feedback");
                return
            }
        //     localStorage.setItem("token", resp.data.token)
            this.$toast.success("Cool! You can now see the App")
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
.app-view {
    height: 100%;
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
    display: grid;
    align-content: space-around;
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
