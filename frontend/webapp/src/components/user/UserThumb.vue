<template>
    <div class="profile-row">
        <div class="user-image">
            <img v-bind:src="user.user_img_url" alt=""
            v-bind:class="[status ? 'connected' : 'not-connected']">
        </div>
        <div class="user-info">
            <!-- <div class="name">{{user.name}}</div>
            <div class="username">@{{user.username}}</div> -->
        </div>
    </div>
</template>

<script>
    export default {
        name: 'UserThumb',
        data() {
            return { user: "Konstantin" };
        },
        props: ['status'],
    };
</script>

<style scoped>
.profile-row {
    display: flex;
    align-items: center;
    width: max-content;
    max-width: 250px;
}
.profile-row:hover {
    cursor: pointer;
}
.user-image img {
    object-fit: cover;
    width: 45px;
    height: 45px;
    border-radius: 50%;
    margin-right: 10px;
}
.connected {
    border: 2px solid #fff;
}
.not-connected {
    border: 2px solid #30363D;
}
.user-info {
    line-height: 1.2;
    display: grid;
    align-items: center;
}
.name, .username {
    color: #fff;
    max-width: 200px;
    overflow:hidden;
    white-space:nowrap;
    text-overflow: ellipsis;
}

.name {
    font-size: 18px;
    font-weight: bolder;
}
.username {
    opacity: 0.7;
}
</style>
