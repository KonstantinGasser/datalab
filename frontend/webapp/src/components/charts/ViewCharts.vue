<template>
    <div class="view_app">
        <div class="d-flex">
            <SinglePoint :name="'Page Views'" :value="'24.000'" :emoji="'icon-users'" :theme="'green'" />
            <!-- <SinglePoint :name="'New Session'" :value="'+42%'" :emoji="'icon-target'" :theme="'blue'" /> -->
            <SinglePoint :name="'Avg. Session Time'" :value="'~17 min'" :emoji="'icon-clock'" :theme="'orange'" />
            <SinglePoint :name="'Most used Browser'" :value="'Safari'" :emoji="'icon-globe'" :theme="'blue'" />
        </div>
    </div>
</template>

<script>
    import SinglePoint from '@/components/charts/SinglePoint';

    export default {
        name: 'ViewCharts',
        components: {
            SinglePoint,
        }
        
    };
</script>

<style scoped>
.view_app {
    display: grid;
    height: 100%;
}

</style>