<template>
    <div class="view_component mt-3">
        <!-- <selectMember v-if="showSelectMember" :member_list="dummyMember" @close="showSelectMember = false" /> -->
        <div class="ml-1 d-flex align-center justify-start">
            <div class="action" @click="addMember">
                <span class="icon icon-user-plus hover big"></span>
            </div>
        </div>
        <div class="d-flex flex-wrap">
            <div v-for="dummy in latestMemberList" :key="dummy.uuid" class="user_card">
                <div class="d-flex align-center justify-end">
                    <div @click="removeMember(dummy.id)">
                        <span class="icon icon-user-minus hover big"></span>
                    </div>
                </div>
                <div class="d-flex align-center justify-center">
                    <img class="circle-img medium" :src="dummy.profile_img_url" alt="">
                </div>
                <div class="member_info">
                    <span class="member_name dots">{{dummy.first_name}} {{dummy.last_name}}</span>
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
                return this.latest_member_list;
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
                this.dummyMember = this.dummyMember.filter(item => item.id !== id);
            },
            addMember() {
                this.showSelectMember = !this.showSelectMember;
                // const id = this.dummyMember.length > 0 ? this.dummyMember[this.dummyMember.length -1].id+1 : 1;
                // this.dummyMember.push({id: id});
            }
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
