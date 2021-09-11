<template>
    <div class="single_box">
        <div class="d-flex align-center justify-between">
            <div class="emoji" :class="theme"><span class="icon" :class="emoji"></span></div>
            <div class="single_data">
                <div class="value">{{value}}</div>
                <div class="sub">{{name}}</div>
            </div>
        </div>
        <div class="time_for"><span class="icon icon-calendar"></span> Last Week</div>
    </div>
</template>

<script>
export default {
    name: "SinglePoint",
    props: ['name', 'emoji', 'value', 'theme'],
}
</script>

<style scoped>

.single_box {
    background: var(--sub-bg);
    width: 250px;
    height: min-content;
    border-radius: 14px;
    box-shadow: 0 0 10px 1px var(--sub-border);
    padding: 10px 15px;
    margin: 0 10px;
}

.value {
    font-size: 30px;
    font-weight: bolder;
    text-align: right;
    color: var(--h-color);
}
.sub {
    text-align: right;
}
.emoji {
    font-size: 30px;
    opacity: 0.6;
    height: 50px;
    width: 50px;
    background: var(--sub-bg);
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 14px;
    text-align: center;
    margin: 10px 0;
}
.green {
    color: #10d574;
    box-shadow: 0 0 14px 1px #10d574aa;
}

.blue {
    color: #3FB0FE;
    box-shadow: 0 0 14px 1px #3FB0FE;
}

.orange {
    color: #FEB43C;
    box-shadow: 0 0 14px 1px #FEB43C;
}

</style>