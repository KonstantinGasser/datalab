<template>
    <div>
        <div class="view_component">
            already invited {{app_uuid}}
        </div>
        <div class="view_component">
            request invite
        </div>
    </div>
</template>

<script>
export default {
    name: "InviteMember",
    props: {
        app_uuid: {
            type: String,
            default: null,
        },
    },
}
</script>