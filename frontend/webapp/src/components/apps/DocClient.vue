<template>
    <div>
        <div v-if="!hasToken" class="view_component">
            <div class="form-row">
                <div class="form-group col">
                    <div class="info-txt">
                        Mhm looks like you have not yet created and application token. You need one of those
                        in order to get the client side working... <a @click="navCreateToken">go - create one</a> 😉
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "DocClientLib",
    data() {
        return {};
    },
    props: ['hasToken'],
    methods: {
        navCreateToken() {
            this.$emit("goCreateToken");
        },
    },
}
</script>

<style scoped>
.view_component {
    margin-top: 15px;
    padding: 15px;
    border-radius: 8px;
    height: max-width;
}
</style>