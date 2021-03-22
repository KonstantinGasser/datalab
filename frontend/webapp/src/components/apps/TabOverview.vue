<template>
<div></div>
    <!-- <div class="view_component">
        <div class="d-flex justify-between align-center">
            <h1>General: {{ isEdit }}</h1>
            <span @click="setMode" class="icon icon-edit big hover"></span>
        </div>
        <div class="d-flex align-center justify-evenly flex-wrap">
            <div class="test d-flex align-center">testing something</div>
            <div class="test d-flex align-center">testing another thing</div>
        </div>
    </div> -->
</template>

<script>

    export default {
        name: 'TabOverview',
        components: {},
        data() {
            return {
                isEdit: false,
            };
        },
        methods: {
            setMode() {
                this.isEdit = !this.isEdit;
                // emit panel is in edit mode
                // diable tabs until saved
                this.$emit('inEdit', this.isEdit);
            }
        },
    };
</script>

<style scoped>
.view_component {
    margin-top: 15px;
    padding: 15px;
    background: #fff;
    border-radius: 8px;
    height: 94%;
}
h2 {
    margin: 5px 0px;
}
.pos_1_1 {
    grid-row: 1;
}
.pos_1_2 {
    grid-row: 2;
}

.test {
    padding: 5px;
    background: lightgray;
    height: 100%;
    width: auto;
}
</style>
