<template>
    <div class="view_component mt-3">
        <!-- <selectMember v-if="showSelectMember" :member_list="dummyMember" @close="showSelectMember = false" /> -->
        <div class="ml-1 d-flex align-center justify-start">
            <div class="action" @click="addMember">
                <span class="icon icon-user-plus hover big"></span>
            </div>
        </div>
        <div class="d-flex flex-wrap">
            <div v-for="dummy in latestMemberList" :key="dummy.uuid" class="user_card" :class="(loggedIn().sub === dummy.uuid)? 'is_owner': ''">
                <div class="d-flex align-center justify-evenly">
                    <span class="member_name dots">Konstantin Gasser</span>
                    <div v-if="loggedIn().sub !== dummy.uuid" @click="removeMember(dummy.uuid)">
                        <span class="icon icon-user-minus hover big"></span>
                    </div>
                </div>
                <div class="d-flex align-center justify-center">
                    <img class="circle-img medium" :src="dummy.profile_img_url" alt="">
                </div>
                <div class="member_info">

                    <span class="member_name dots">@{{dummy.username}}</span>
                    <!-- <span class="member_name dots">{{dummy.first_name}} {{dummy.last_name}}</span> -->
                    <span class="member_pos dots">{{dummy.orgn_position}}</span>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import selectMember from '@/components/apps/ModalSelectMember.vue';

    export default {
        name: 'TabMember',
        components: {
            selectMember,
        },
        data() {
            return {
                showSelectMember: false,
                isEdit: false,
                latest_member_list: [],
            };
        },
        props: ['member_list'],
        mounted() {
            this.latest_member_list = this.$props.member_list? this.$props.member_list : [];
            console.log("member list: ", this.latest_member_list);
        },
        computed: {
            latestMemberList() {
                return this.latest_member_list.reverse();
            }
        },
        methods: {
            setMode() {
                this.isEdit = !this.isEdit;
                // emit panel is in edit mode
                // disable tabs until saved
                this.$emit('inEdit', this.isEdit);
            },
            removeMember(id) {
                this.dummyMember = this.dummyMember.filter(item => item.uuid !== id);
            },
            addMember() {
                this.showSelectMember = !this.showSelectMember;
                // const id = this.dummyMember.length > 0 ? this.dummyMember[this.dummyMember.length -1].id+1 : 1;
                // this.dummyMember.push({id: id});
            },
            loggedIn () {
                const token = localStorage.getItem("token");
                var base64Url = token.split('.')[1];
                var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
                var jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
                    return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
                }).join(''));

                return JSON.parse(jsonPayload);
            },
        },
    };
</script>

<style>
.user_card {
    width: 150px;
    height: 150px;
    margin: 5px;
    border-radius: 8px;
    /* box-shadow: 0px 0px 5px 2px rgb(0,0,0, 0.1); */
    border: 1px solid #30363D;
    padding: 5px;
    background: #1E1E1E;
    display: grid;
    align-content: space-between; 
}
.user_card.is_owner {
    border: 1px solid #10d574;
}
.user_card .member_info {
    display: grid;
    justify-content: center;
    text-align: center;
}
.member_info .member_name {
    color: #fff;
    font-weight: 600;
}

.member_info .member_pos {
    opacity: 0.8;
    font-size: 12px;
}

</style>
