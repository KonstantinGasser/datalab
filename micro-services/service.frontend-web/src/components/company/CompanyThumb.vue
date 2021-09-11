<template>
    <div class="profile-row" @click="redirect">
        <div class="company_info">
            <div class="name">{{company}}</div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'CompanyThumb',
        data() {
            return { company: "com.datalab"};
        },
        props: ['status'],
        methods: {
            redirect() {
                this.$emit('setActive', 'view_account');
            },
        },
    };
</script>

<style scoped>
.profile-row {
    display: flex;
    align-items: center;
}
.profile-row:hover {
    cursor: pointer;
}
.company_info {
    /* color: #fff; */
    font-size: 28px;
    font-weight: bold;
}
</style>
